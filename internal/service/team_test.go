package service_test

import (
	"context"
	"testing"

	"github.com/deimossy/pr-reviewer-assignment-service/internal/models"
	"github.com/deimossy/pr-reviewer-assignment-service/internal/service/mocks"
	"github.com/stretchr/testify/assert"
)

func TestTeamService_Create(t *testing.T) {
	mockTeam := mocks.NewMockTeamService(t)
	ctx := context.Background()
	team := &models.Team{TeamName: "dev"}

	mockTeam.EXPECT().Create(ctx, team).Return(nil)

	err := mockTeam.Create(ctx, team)
	assert.NoError(t, err)
}

func TestTeamService_GetByName(t *testing.T) {
	mockTeam := mocks.NewMockTeamService(t)
	ctx := context.Background()
	name := "dev"
	expected := &models.Team{TeamName: name}

	mockTeam.EXPECT().GetByName(ctx, name).Return(expected, nil)

	result, err := mockTeam.GetByName(ctx, name)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}
