package service

import (
	"context"
	"github.com/deimossy/pr-reviewer-assignment-service/internal/models"
	"github.com/deimossy/pr-reviewer-assignment-service/internal/repository"
	"go.uber.org/zap"
)

type TeamService interface {
	Create(ctx context.Context, team *models.Team) error
	GetByName(ctx context.Context, name string) (*models.Team, error)
}

type teamService struct {
	teamRepo repository.TeamRepository
	userSvc  UserService
	logger   *zap.Logger
}

func NewTeamService(
	teamRepo repository.TeamRepository,
	userSvc UserService,
	logger *zap.Logger,
) TeamService {
	return &teamService{
		teamRepo: teamRepo,
		userSvc:  userSvc,
		logger:   logger,
	}
}

func (s *teamService) Create(ctx context.Context, team *models.Team) error {
	if err := s.teamRepo.Create(ctx, *team); err != nil {
		return err
	}

	for _, m := range team.Members {
		user := &models.User{
			UserId:   m.UserId,
			Username: m.Username,
			TeamName: team.TeamName,
			IsActive: m.IsActive,
		}
		if err := s.userSvc.Create(ctx, user); err != nil {
			return err
		}
	}
	return nil
}

func (s *teamService) GetByName(ctx context.Context, name string) (*models.Team, error) {
	team, err := s.teamRepo.GetByName(ctx, name)
	if err != nil {
		return nil, err
	}

	users, err := s.userSvc.GetTeamMembers(ctx, name)
	if err != nil {
		return nil, err
	}

	for _, u := range users {
		team.Members = append(team.Members, models.TeamMember{
			UserId:   u.UserId,
			Username: u.Username,
			IsActive: u.IsActive,
		})
	}
	return team, nil
}
