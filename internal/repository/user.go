package repository

import (
	"context"
	"github.com/deimossy/pr-reviewer-assignment-service/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, user models.User) error
	SetActive(ctx context.Context, userID string, active bool) error
	GetByID(ctx context.Context, userID string) (*models.User, error)
	GetTeamMembers(ctx context.Context, teamName string) ([]models.User, error)
}
