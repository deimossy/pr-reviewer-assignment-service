package service_test

import (
	"context"
	"testing"

	"github.com/deimossy/pr-reviewer-assignment-service/internal/models"
	"github.com/deimossy/pr-reviewer-assignment-service/internal/service/mocks"
	"github.com/stretchr/testify/assert"
)

func TestUserService_Create(t *testing.T) {
	mockUser := mocks.NewMockUserService(t)
	ctx := context.Background()
	user := &models.User{UserId: "u1"}

	mockUser.EXPECT().Create(ctx, user).Return(nil)

	err := mockUser.Create(ctx, user)
	assert.NoError(t, err)
}

func TestUserService_GetByID(t *testing.T) {
	mockUser := mocks.NewMockUserService(t)
	ctx := context.Background()
	userID := "u1"
	expected := &models.User{UserId: userID}

	mockUser.EXPECT().GetByID(ctx, userID).Return(expected, nil)

	user, err := mockUser.GetByID(ctx, userID)
	assert.NoError(t, err)
	assert.Equal(t, expected, user)
}

func TestUserService_GetTeamMembers(t *testing.T) {
	mockUser := mocks.NewMockUserService(t)
	ctx := context.Background()
	teamName := "dev"
	expected := []models.User{{UserId: "u1"}}

	mockUser.EXPECT().GetTeamMembers(ctx, teamName).Return(expected, nil)

	result, err := mockUser.GetTeamMembers(ctx, teamName)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestUserService_SetActive(t *testing.T) {
	mockUser := mocks.NewMockUserService(t)
	ctx := context.Background()
	userID := "u1"

	mockUser.EXPECT().SetActive(ctx, userID, true).Return(nil)

	err := mockUser.SetActive(ctx, userID, true)
	assert.NoError(t, err)
}
