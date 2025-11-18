package mapper

import (
	"errors"
	"time"

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
		return nil, errors.New("Status is required")
	}

	return &models.PullRequest{
		PullRequestId:     d.PullRequestID,
		PullRequestName:   d.PullRequestName,
		AuthorId:          d.AuthorID,
		Status:            d.Status,
		CreatedAt:         timeOrZero(d.CreatedAt),
		MergedAt:          d.MergedAt,
		AssignedReviewers: d.AssignedReviewers,
	}, nil
}

func PullRequestShortFromDTO(d *dto.PullRequestShortDTO) (*models.PullRequestShort, error) {
	if d == nil {
		return nil, errors.New("PullRequestShortDTO is nil")
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
		return nil, errors.New("Status is required")
	}

	return &models.PullRequestShort{
		PullRequestId:   d.PullRequestID,
		PullRequestName: d.PullRequestName,
		AuthorId:        d.AuthorID,
		Status:          d.Status,
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

func UserFromDTO(d *dto.UserDTO) (*models.User, error) {
	if d == nil {
		return nil, errors.New("UserDTO is nil")
	}
	if d.UserID == "" {
		return nil, errors.New("UserID is required")
	}
	if d.Username == "" {
		return nil, errors.New("Username is required")
	}
	if d.TeamName == "" {
		return nil, errors.New("TeamName is required")
	}

	return &models.User{
		UserId:   d.UserID,
		Username: d.Username,
		TeamName: d.TeamName,
		IsActive: d.IsActive,
	}, nil
}

func timeOrZero(t *time.Time) time.Time {
	if t == nil {
		return time.Time{}
	}

	return *t
}
