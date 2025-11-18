package repository

import (
	"context"
	"github.com/deimossy/pr-reviewer-assignment-service/internal/models"
)

type PullRequestRepository interface {
	Create(ctx context.Context, pr models.PullRequest) error
	Merge(ctx context.Context, prID string) (*models.PullRequest, error)
	GetByID(ctx context.Context, prID string) (*models.PullRequest, error)
	ListByUser(ctx context.Context, userID string) ([]models.PullRequestShort, error)

	AssignReview(ctx context.Context, prID string, reviewerIDs []string) error
	ReplaceReview(ctx context.Context, prID string, oldReviewerID string, newReviewerID string) error
	GetReviewers(ctx context.Context, prID string) ([]string, error)
}
