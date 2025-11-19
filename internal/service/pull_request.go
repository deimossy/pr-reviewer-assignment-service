package service

import (
	"context"
	"errors"
	"github.com/deimossy/pr-reviewer-assignment-service/internal/models"
	"github.com/deimossy/pr-reviewer-assignment-service/internal/repository"
	customErrors "github.com/deimossy/pr-reviewer-assignment-service/pkg/errors"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

type PullRequestService interface {
	Create(ctx context.Context, pr *models.PullRequest) error
	Merge(ctx context.Context, prID string) (*models.PullRequest, error)
	ListByUser(ctx context.Context, userID string) ([]models.PullRequestShort, error)
	ReplaceReview(ctx context.Context, prID, oldReviewerID string) (string, *models.PullRequest, error)
	Stats(ctx context.Context) (map[string]int, error)
}

type pullRequestService struct {
	repo        repository.PullRequestRepository
	userService UserService
	teamService TeamService
	logger      *zap.Logger
}

func NewPullRequestService(repo repository.PullRequestRepository, userSvc UserService, teamSvc TeamService, logger *zap.Logger) PullRequestService {
	return &pullRequestService{
		repo:        repo,
		userService: userSvc,
		teamService: teamSvc,
		logger:      logger,
	}
}

func (s *pullRequestService) Create(ctx context.Context, pr *models.PullRequest) error {
	if pr.CreatedAt == nil {
		now := time.Now().UTC()
		pr.CreatedAt = &now
	}

	if pr.Status == "" {
		pr.Status = "OPEN"
	}

	if len(pr.AssignedReviewers) == 0 {
		if err := s.assignReviewers(ctx, pr); err != nil && !errors.Is(err, customErrors.ErrNoCandidate) {
			return err
		}
	}

	err := s.repo.Create(ctx, *pr)
	if err != nil {
		s.logger.Error("failed to create PR",
			zap.String("pr_id", pr.PullRequestId),
			zap.Error(err),
		)
		return err
	}
	return nil
}

func (s *pullRequestService) Merge(ctx context.Context, prID string) (*models.PullRequest, error) {
	pr, err := s.repo.Merge(ctx, prID)
	if err != nil {
		s.logger.Error("failed to merge PR",
			zap.String("pr_id", prID),
			zap.Error(err),
		)

		if errors.Is(err, customErrors.ErrPRAlreadyMerged) {
			return pr, customErrors.ErrPRAlreadyMerged
		}
		if errors.Is(err, customErrors.ErrPRNotFound) {
			return nil, customErrors.ErrPRNotFound
		}
		return nil, err
	}
	return pr, nil
}

func (s *pullRequestService) ListByUser(ctx context.Context, userID string) ([]models.PullRequestShort, error) {
	prs, err := s.repo.ListByUser(ctx, userID)
	if err != nil {
		s.logger.Error("failed to list PRs by user",
			zap.String("user_id", userID),
			zap.Error(err),
		)
		return nil, err
	}
	return prs, nil
}

func (s *pullRequestService) ReplaceReview(ctx context.Context, prID, oldReviewerID string) (string, *models.PullRequest, error) {
	pr, err := s.getPRForReplacement(ctx, prID)
	if err != nil {
		return "", nil, err
	}

	newReviewerID, err := s.selectReplacementReviewer(ctx, pr, oldReviewerID)
	if err != nil {
		return "", nil, err
	}

	if err := s.updatePRReviewer(ctx, prID, oldReviewerID, newReviewerID); err != nil {
		return "", nil, err
	}

	pr, err = s.repo.GetByID(ctx, prID)
	if err != nil {
		s.logger.Error("failed to get PR after replacing reviewer",
			zap.String("pr_id", prID),
			zap.Error(err),
		)
		return "", nil, err
	}
	return newReviewerID, pr, nil
}

func (s *pullRequestService) Stats(ctx context.Context) (map[string]int, error) {
	stats, err := s.repo.GetAssignmentsCount(ctx)
	if err != nil {
		s.logger.Error("failed to get PR stats",
			zap.Error(err),
		)
		return nil, err
	}
	return stats, nil
}

func (s *pullRequestService) getPRForReplacement(ctx context.Context, prID string) (*models.PullRequest, error) {
	pr, err := s.repo.GetByID(ctx, prID)
	if err != nil {
		return nil, customErrors.ErrPRNotFound
	}
	if pr.Status == "MERGED" {
		return nil, customErrors.ErrPRAlreadyMerged
	}
	return pr, nil
}

func (s *pullRequestService) selectReplacementReviewer(ctx context.Context, pr *models.PullRequest, oldReviewerID string) (string, error) {
	author, err := s.userService.GetByID(ctx, pr.AuthorId)
	if err != nil {
		if errors.Is(err, customErrors.ErrUserNotFound) {
			return "", customErrors.ErrNoCandidate
		}
		return "", err
	}

	team, err := s.teamService.GetByName(ctx, author.TeamName)
	if err != nil {
		if errors.Is(err, customErrors.ErrTeamNotFound) {
			return "", customErrors.ErrNoCandidate
		}
		return "", err
	}

	candidates := s.getReplacementCandidates(pr.AssignedReviewers, oldReviewerID, team.Members)
	if len(candidates) == 0 {
		return "", customErrors.ErrNoCandidate
	}
	return pickRandom(candidates), nil
}

func (s *pullRequestService) updatePRReviewer(ctx context.Context, prID, oldReviewerID, newReviewerID string) error {
	if err := s.repo.ReplaceReview(ctx, prID, oldReviewerID, newReviewerID); err != nil {
		if errors.Is(err, customErrors.ErrNotAssigned) {
			return customErrors.ErrNotAssigned
		}
		return err
	}
	return nil
}

func (s *pullRequestService) assignReviewers(ctx context.Context, pr *models.PullRequest) error {
	author, err := s.userService.GetByID(ctx, pr.AuthorId)
	if err != nil {
		if errors.Is(err, customErrors.ErrUserNotFound) {
			return customErrors.ErrNoCandidate
		}
		return err
	}

	team, err := s.teamService.GetByName(ctx, author.TeamName)
	if err != nil {
		if errors.Is(err, customErrors.ErrTeamNotFound) {
			return customErrors.ErrNoCandidate
		}
		return err
	}

	var candidates []string
	for _, member := range team.Members {
		if member.UserId != author.UserId && member.IsActive {
			candidates = append(candidates, member.UserId)
		}
	}

	if len(candidates) == 0 {
		return customErrors.ErrNoCandidate
	}

	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	r.Shuffle(len(candidates), func(i, j int) {
		candidates[i], candidates[j] = candidates[j], candidates[i]
	})

	n := 2
	if len(candidates) < 2 {
		n = len(candidates)
	}
	pr.AssignedReviewers = candidates[:n]
	return nil
}

func (s *pullRequestService) getReplacementCandidates(current []string, old string, members []models.TeamMember) []string {
	var candidates []string
	for _, m := range members {
		if m.UserId != old && m.IsActive && !contains(current, m.UserId) {
			candidates = append(candidates, m.UserId)
		}
	}
	return candidates
}

func pickRandom(slice []string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return slice[r.Intn(len(slice))]
}

func contains(slice []string, userId string) bool {
	for _, s := range slice {
		if s == userId {
			return true
		}
	}
	return false
}
