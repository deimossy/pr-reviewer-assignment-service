package teamrepo

import (
	"context"
	"github.com/deimossy/pr-reviewer-assignment-service/internal/models"
	"github.com/deimossy/pr-reviewer-assignment-service/internal/repository"
	"github.com/jmoiron/sqlx"
)

type teamRepository struct {
	db *sqlx.DB
}

func NewTeamRepository(db *sqlx.DB) repository.TeamRepository {
	return &teamRepository{db: db}
}

func (r *teamRepository) Create(ctx context.Context, team models.Team) error {
	// TODO: implement
	return nil
}

func (r *teamRepository) GetByName(ctx context.Context, name string) (*models.Team, error) {
	// TODO: implement
	return nil, nil
}
