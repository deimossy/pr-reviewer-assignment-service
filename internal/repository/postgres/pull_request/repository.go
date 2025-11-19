package prrepo

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/deimossy/pr-reviewer-assignment-service/internal/models"
	"github.com/deimossy/pr-reviewer-assignment-service/internal/repository"
	customErrors "github.com/deimossy/pr-reviewer-assignment-service/pkg/errors"
	pgHelpers "github.com/deimossy/pr-reviewer-assignment-service/pkg/postgres"
	"github.com/jmoiron/sqlx"
)

type pullRequestRepository struct {
	db *sqlx.DB
}

func NewPullRequestRepository(db *sqlx.DB) repository.PullRequestRepository {
	return &pullRequestRepository{db: db}
}

func (r *pullRequestRepository) Create(ctx context.Context, pr models.PullRequest) error {
	_, err := r.db.ExecContext(ctx, createPRQuery, pr.PullRequestId, pr.PullRequestName, pr.AuthorId, pr.Status, pr.CreatedAt)
	if err != nil {
		if e := pgHelpers.CheckUnique(err, "pull_request"); !errors.Is(e, err) {
			return e
		}
		if e := pgHelpers.CheckForeignKey(err, "user"); !errors.Is(e, err) {
			return e
		}
		return err
	}

	for _, revID := range pr.AssignedReviewers {
		if _, err := r.db.ExecContext(ctx, insertReviewerQuery, pr.PullRequestId, revID); err != nil {
			if e := pgHelpers.CheckForeignKey(err, "user"); !errors.Is(e, err) {
				return e
			}
			return err
		}
	}
	return nil
}

func (r *pullRequestRepository) Merge(ctx context.Context, prID string) (*models.PullRequest, error) {
	pr, err := r.GetByID(ctx, prID)
	if err != nil {
		return nil, err
	}

	if pr.Status == "MERGED" {
		return pr, nil
	}

	now := time.Now().UTC()
	_, err = r.db.ExecContext(ctx, updatePRStatusQuery, prID, "MERGED", now)
	if err != nil {
		return nil, err
	}

	pr.Status = "MERGED"
	pr.MergedAt = &now
	return pr, nil
}

func (r *pullRequestRepository) GetByID(ctx context.Context, prID string) (*models.PullRequest, error) {
	var pr models.PullRequest
	if err := r.db.GetContext(ctx, &pr, getPRByIDQuery, prID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, customErrors.ErrPRNotFound
		}
		return nil, err
	}

	reviewers, err := r.GetReviewers(ctx, prID)
	if err != nil {
		return nil, err
	}
	pr.AssignedReviewers = reviewers
	return &pr, nil
}

func (r *pullRequestRepository) ListByUser(ctx context.Context, userID string) ([]models.PullRequestShort, error) {
	var prs []models.PullRequestShort
	if err := r.db.SelectContext(ctx, &prs, listPRsByUserQuery, userID); err != nil {
		return nil, err
	}
	return prs, nil
}

func (r *pullRequestRepository) ReplaceReview(ctx context.Context, prID string, oldReviewerID string, newReviewerID string) error {
	pr, err := r.GetByID(ctx, prID)
	if err != nil {
		return err
	}

	if pr.Status == "MERGED" {
		return customErrors.ErrPRAlreadyMerged
	}

	res, err := r.db.ExecContext(ctx, deleteReviewerQuery, prID, oldReviewerID)
	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return customErrors.ErrNotAssigned
	}

	if _, err := r.db.ExecContext(ctx, insertReviewerQuery, prID, newReviewerID); err != nil {
		return err
	}
	return nil
}

func (r *pullRequestRepository) GetReviewers(ctx context.Context, prID string) ([]string, error) {
	var reviewers []string
	if err := r.db.SelectContext(ctx, &reviewers, getReviewersQuery, prID); err != nil {
		return nil, err
	}
	return reviewers, nil
}
