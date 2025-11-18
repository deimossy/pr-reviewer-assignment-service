package service

import (
	"context"

	"github.com/deimossy/pr-reviewer-assignment-service/internal/models"
	"github.com/deimossy/pr-reviewer-assignment-service/internal/repository"
)

type TeamService interface {
	Create(ctx context.Context, team *models.Team) error
	GetByName(ctx context.Context, name string) (*models.Team, error)
}

type teamService struct {
	repo repository.TeamRepository
}

func NewTeamService(repo repository.TeamRepository) TeamService {
	return &teamService{repo: repo}
}

func (s *teamService) Create(ctx context.Context, team *models.Team) error {
	return s.repo.Create(ctx, *team)
}

func (s *teamService) GetByName(ctx context.Context, name string) (*models.Team, error) {
	return s.repo.GetByName(ctx, name)
}
