package prrepo

import (
	"context"
	"github.com/deimossy/pr-reviewer-assignment-service/internal/models"
	"github.com/deimossy/pr-reviewer-assignment-service/internal/repository"
	"github.com/jmoiron/sqlx"
)

type pullRequestRepository struct {
	db *sqlx.DB
}

func NewPullRequestRepository(db *sqlx.DB) repository.PullRequestRepository {
	return &pullRequestRepository{db: db}
}

func (r *pullRequestRepository) Create(ctx context.Context, pr models.PullRequest) error {
	// TODO: implement
	return nil
}

func (r *pullRequestRepository) Merge(ctx context.Context, prID string) (*models.PullRequest, error) {
	// TODO: implement
	return nil, nil
}

func (r *pullRequestRepository) GetByID(ctx context.Context, prID string) (*models.PullRequest, error) {
	// TODO: implement
	return nil, nil
}

func (r *pullRequestRepository) ListByUser(ctx context.Context, userID string) ([]models.PullRequestShort, error) {
	// TODO: implement
	return nil, nil
}

func (r *pullRequestRepository) AssignReview(ctx context.Context, prID string, reviewerIDs []string) error {
	// TODO: implement
	return nil
}

func (r *pullRequestRepository) ReplaceReview(ctx context.Context, prID string, oldReviewerID string, newReviewerID string) error {
	// TODO: implement
	return nil
}

func (r *pullRequestRepository) GetReviewers(ctx context.Context, prID string) ([]string, error) {
	// TODO: implement
	return nil, nil
}
