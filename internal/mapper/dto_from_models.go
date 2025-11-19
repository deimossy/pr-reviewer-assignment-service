package mapper

import (
	"github.com/deimossy/pr-reviewer-assignment-service/internal/dto"
	"github.com/deimossy/pr-reviewer-assignment-service/internal/models"
)

func PullRequestToDTO(m *models.PullRequest) *dto.PullRequestDTO {
	if m == nil {
		return nil
	}
	return &dto.PullRequestDTO{
		PullRequestID:     m.PullRequestId,
		PullRequestName:   m.PullRequestName,
		AuthorID:          m.AuthorId,
		Status:            m.Status,
		CreatedAt:         m.CreatedAt,
		MergedAt:          m.MergedAt,
		AssignedReviewers: m.AssignedReviewers,
	}
}

func PullRequestShortToDTO(m *models.PullRequestShort) *dto.PullRequestShortDTO {
	return &dto.PullRequestShortDTO{
		PullRequestID:   m.PullRequestId,
		PullRequestName: m.PullRequestName,
		AuthorID:        m.AuthorId,
		Status:          m.Status,
	}
}

func TeamToDTO(m *models.Team) *dto.TeamDTO {
	members := make([]dto.TeamMemberDTO, len(m.Members))
	for i, mm := range m.Members {
		members[i] = TeamMemberToDTO(&mm)
	}
	return &dto.TeamDTO{
		TeamName: m.TeamName,
		Members:  members,
	}
}

func TeamMemberToDTO(m *models.TeamMember) dto.TeamMemberDTO {
	return dto.TeamMemberDTO{
		UserID:   m.UserId,
		Username: m.Username,
		IsActive: m.IsActive,
	}
}

func UserToDTO(m *models.User) *dto.UserDTO {
	return &dto.UserDTO{
		UserID:   m.UserId,
		Username: m.Username,
		TeamName: m.TeamName,
		IsActive: m.IsActive,
	}
}
