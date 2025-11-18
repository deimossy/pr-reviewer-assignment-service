package service

import (
	"context"

	"github.com/deimossy/pr-reviewer-assignment-service/internal/models"
	"github.com/deimossy/pr-reviewer-assignment-service/internal/repository"
)

type UserService interface {
	Create(ctx context.Context, user *models.User) error
	SetActive(ctx context.Context, userID string, active bool) error
	GetByID(ctx context.Context, userID string) (*models.User, error)
	GetTeamMembers(ctx context.Context, teamName string) ([]models.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Create(ctx context.Context, user *models.User) error {
	return s.repo.Create(ctx, *user)
}

func (s *userService) SetActive(ctx context.Context, userID string, active bool) error {
	return s.repo.SetActive(ctx, userID, active)
}

func (s *userService) GetByID(ctx context.Context, userID string) (*models.User, error) {
	return s.repo.GetByID(ctx, userID)
}

func (s *userService) GetTeamMembers(ctx context.Context, teamName string) ([]models.User, error) {
	return s.repo.GetTeamMembers(ctx, teamName)
}
