package userrepo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/deimossy/pr-reviewer-assignment-service/internal/models"
	"github.com/deimossy/pr-reviewer-assignment-service/internal/repository"
	customErrors "github.com/deimossy/pr-reviewer-assignment-service/pkg/errors"
	pgHelpers "github.com/deimossy/pr-reviewer-assignment-service/pkg/postgres"
	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user models.User) error {
	_, err := r.db.ExecContext(ctx, createUserQuery, user.UserId, user.Username, user.TeamName, user.IsActive)
	if err != nil {
		return pgHelpers.CheckUnique(err, "user")
	}
	return nil
}

func (r *userRepository) SetActive(ctx context.Context, userID string, active bool) error {
	res, err := r.db.ExecContext(ctx, setActiveQuery, userID, active)
	if err != nil {
		return pgHelpers.CheckForeignKey(err, "user")
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return customErrors.ErrUserNotFound
	}
	return nil
}

func (r *userRepository) GetByID(ctx context.Context, userID string) (*models.User, error) {
	var u models.User
	if err := r.db.GetContext(ctx, &u, getByIDQuery, userID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, customErrors.ErrUserNotFound
		}
		return nil, err
	}
	return &u, nil
}

func (r *userRepository) GetTeamMembers(ctx context.Context, teamName string) ([]models.User, error) {
	var members []models.User
	if err := r.db.SelectContext(ctx, &members, getTeamMembersQuery, teamName); err != nil {
		return nil, err
	}
	return members, nil
}
