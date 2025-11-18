package repository

import (
	"context"
	"github.com/deimossy/pr-reviewer-assignment-service/internal/models"
)

type TeamRepository interface {
	Create(ctx context.Context, team models.Team) error
	GetByName(ctx context.Context, name string) (*models.Team, error)
}
