package service

import (
	"context"

	"github.com/deimossy/pr-reviewer-assignment-service/internal/models"
	"github.com/deimossy/pr-reviewer-assignment-service/internal/repository"
)

type PullRequestService interface {
	Create(ctx context.Context, pr *models.PullRequest) error
	Merge(ctx context.Context, prID string) (*models.PullRequest, error)
	GetByID(ctx context.Context, prID string) (*models.PullRequest, error)
	ListByUser(ctx context.Context, userID string) ([]models.PullRequestShort, error)
	AssignReview(ctx context.Context, prID string, reviewerIDs []string) error
	ReplaceReview(ctx context.Context, prID string, oldReviewerID, newReviewerID string) error
	GetReviewers(ctx context.Context, prID string) ([]string, error)
}

type pullRequestService struct {
	repo repository.PullRequestRepository
}

func NewPullRequestService(repo repository.PullRequestRepository) PullRequestService {
	return &pullRequestService{repo: repo}
}

func (s *pullRequestService) Create(ctx context.Context, pr *models.PullRequest) error {
	return s.repo.Create(ctx, *pr)
}

func (s *pullRequestService) Merge(ctx context.Context, prID string) (*models.PullRequest, error) {
	return s.repo.Merge(ctx, prID)
}

func (s *pullRequestService) GetByID(ctx context.Context, prID string) (*models.PullRequest, error) {
	return s.repo.GetByID(ctx, prID)
}

func (s *pullRequestService) ListByUser(ctx context.Context, userID string) ([]models.PullRequestShort, error) {
	return s.repo.ListByUser(ctx, userID)
}

func (s *pullRequestService) AssignReview(ctx context.Context, prID string, reviewerIDs []string) error {
	return s.repo.AssignReview(ctx, prID, reviewerIDs)
}

func (s *pullRequestService) ReplaceReview(ctx context.Context, prID string, oldReviewerID, newReviewerID string) error {
	return s.repo.ReplaceReview(ctx, prID, oldReviewerID, newReviewerID)
}

func (s *pullRequestService) GetReviewers(ctx context.Context, prID string) ([]string, error) {
	return s.repo.GetReviewers(ctx, prID)
}
