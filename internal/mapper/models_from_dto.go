package mapper

import (
	"errors"
	"github.com/deimossy/pr-reviewer-assignment-service/internal/dto"
	"github.com/deimossy/pr-reviewer-assignment-service/internal/models"
)

func PullRequestFromDTO(d *dto.PullRequestDTO) (*models.PullRequest, error) {
	if d == nil {
		return nil, errors.New("PullRequestDTO is nil")
	}
	if d.PullRequestID == "" {
		return nil, errors.New("PullRequestID is required")
	}
	if d.PullRequestName == "" {
		return nil, errors.New("PullRequestName is required")
	}
	if d.AuthorID == "" {
		return nil, errors.New("AuthorID is required")
	}
	if d.Status == "" {
		d.Status = "OPEN"
	}
	return &models.PullRequest{
		PullRequestId:     d.PullRequestID,
		PullRequestName:   d.PullRequestName,
		AuthorId:          d.AuthorID,
		Status:            d.Status,
		CreatedAt:         d.CreatedAt,
		MergedAt:          d.MergedAt,
		AssignedReviewers: d.AssignedReviewers,
	}, nil
}

func TeamFromDTO(d *dto.TeamDTO) (*models.Team, error) {
	if d == nil {
		return nil, errors.New("TeamDTO is nil")
	}
	if d.TeamName == "" {
		return nil, errors.New("TeamName is required")
	}

	members := make([]models.TeamMember, len(d.Members))
	for i, m := range d.Members {
		members[i] = TeamMemberFromDTO(&m)
	}
	return &models.Team{
		TeamName: d.TeamName,
		Members:  members,
	}, nil
}

func TeamMemberFromDTO(d *dto.TeamMemberDTO) models.TeamMember {
	return models.TeamMember{
		UserId:   d.UserID,
		Username: d.Username,
		IsActive: d.IsActive,
	}
}
