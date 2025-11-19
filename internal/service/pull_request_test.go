package service_test

import (
	"context"
	"testing"

	"github.com/deimossy/pr-reviewer-assignment-service/internal/models"
	"github.com/deimossy/pr-reviewer-assignment-service/internal/service/mocks"
	"github.com/stretchr/testify/assert"
)

func TestPullRequestService_Create(t *testing.T) {
	mockPR := mocks.NewMockPullRequestService(t)
	ctx := context.Background()
	pr := &models.PullRequest{PullRequestId: "123"}

	mockPR.EXPECT().Create(ctx, pr).Return(nil)

	err := mockPR.Create(ctx, pr)
	assert.NoError(t, err)
}

func TestPullRequestService_Merge(t *testing.T) {
	mockPR := mocks.NewMockPullRequestService(t)
	ctx := context.Background()
	prID := "123"
	expectedPR := &models.PullRequest{PullRequestId: prID}

	mockPR.EXPECT().Merge(ctx, prID).Return(expectedPR, nil)

	pr, err := mockPR.Merge(ctx, prID)
	assert.NoError(t, err)
	assert.Equal(t, expectedPR, pr)
}

func TestPullRequestService_ListByUser(t *testing.T) {
	mockPR := mocks.NewMockPullRequestService(t)
	ctx := context.Background()
	userID := "user1"
	expected := []models.PullRequestShort{{PullRequestId: "123"}}

	mockPR.EXPECT().ListByUser(ctx, userID).Return(expected, nil)

	result, err := mockPR.ListByUser(ctx, userID)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestPullRequestService_ReplaceReview(t *testing.T) {
	mockPR := mocks.NewMockPullRequestService(t)
	ctx := context.Background()
	prID := "123"
	oldReviewer := "user1"
	newReviewer := "user2"
	pr := &models.PullRequest{PullRequestId: prID}

	mockPR.EXPECT().ReplaceReview(ctx, prID, oldReviewer).Return(newReviewer, pr, nil)

	id, result, err := mockPR.ReplaceReview(ctx, prID, oldReviewer)
	assert.NoError(t, err)
	assert.Equal(t, newReviewer, id)
	assert.Equal(t, pr, result)
}
