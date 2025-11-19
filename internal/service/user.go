package service

import (
	"context"
	"errors"

	"github.com/deimossy/pr-reviewer-assignment-service/internal/models"
	"github.com/deimossy/pr-reviewer-assignment-service/internal/repository"
	customErrors "github.com/deimossy/pr-reviewer-assignment-service/pkg/errors"
	"go.uber.org/zap"
)

type UserService interface {
	Create(ctx context.Context, user *models.User) error
	SetActive(ctx context.Context, userID string, active bool) error
	GetByID(ctx context.Context, userID string) (*models.User, error)
	GetTeamMembers(ctx context.Context, teamName string) ([]models.User, error)
}

type userService struct {
	repo   repository.UserRepository
	logger *zap.Logger
}

func NewUserService(repo repository.UserRepository, logger *zap.Logger) UserService {
	return &userService{repo: repo, logger: logger}
}

func (s *userService) Create(ctx context.Context, user *models.User) error {
	if err := s.repo.Create(ctx, *user); err != nil {
		s.logger.Error("Create user failed",
			zap.String("userID", user.UserId),
			zap.String("username", user.Username),
			zap.String("teamName", user.TeamName),
			zap.Error(err),
		)
		return err
	}
	return nil
}

func (s *userService) SetActive(ctx context.Context, userID string, active bool) error {
	err := s.repo.SetActive(ctx, userID, active)
	if err != nil {
		s.logger.Error("SetActive failed",
			zap.String("userID", userID),
			zap.Bool("active", active),
			zap.Error(err),
		)
		if errors.Is(err, customErrors.ErrUserNotFound) {
			return customErrors.ErrUserNotFound
		}
		return err
	}
	return nil
}

func (s *userService) GetByID(ctx context.Context, userID string) (*models.User, error) {
	user, err := s.repo.GetByID(ctx, userID)
	if err != nil {
		s.logger.Error("GetByID failed",
			zap.String("userID", userID),
			zap.Error(err),
		)
		if errors.Is(err, customErrors.ErrUserNotFound) {
			return nil, customErrors.ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}

func (s *userService) GetTeamMembers(ctx context.Context, teamName string) ([]models.User, error) {
	members, err := s.repo.GetTeamMembers(ctx, teamName)
	if err != nil {
		s.logger.Error("GetTeamMembers failed",
			zap.String("teamName", teamName),
			zap.Error(err),
		)
		return nil, err
	}
	return members, nil
}
