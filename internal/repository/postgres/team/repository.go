package teamrepo

import (
	"context"
	"database/sql"
	"errors"
	pgHelpers "github.com/deimossy/pr-reviewer-assignment-service/pkg/postgres"

	"github.com/deimossy/pr-reviewer-assignment-service/internal/models"
	"github.com/deimossy/pr-reviewer-assignment-service/internal/repository"
	customErrors "github.com/deimossy/pr-reviewer-assignment-service/pkg/errors"
	"github.com/jmoiron/sqlx"
)

type teamRepository struct {
	db *sqlx.DB
}

func NewTeamRepository(db *sqlx.DB) repository.TeamRepository {
	return &teamRepository{db: db}
}

func (r *teamRepository) Create(ctx context.Context, team models.Team) error {
	_, err := r.db.ExecContext(ctx, createTeamQuery, team.TeamName)
	if err != nil {
		return pgHelpers.CheckUnique(err, "team")
	}
	return nil
}

func (r *teamRepository) GetByName(ctx context.Context, name string) (*models.Team, error) {
	var t models.Team
	if err := r.db.GetContext(ctx, &t, getByNameQuery, name); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, customErrors.ErrTeamNotFound
		}
		return nil, err
	}
	return &t, nil
}
