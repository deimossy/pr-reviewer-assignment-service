package postgres

import (
	"errors"
	"fmt"

	customErrors "github.com/deimossy/pr-reviewer-assignment-service/pkg/errors"
	"github.com/jackc/pgx/v5/pgconn"
)

var (
	CodeUniqueViolation     = "23505"
	CodeForeignKeyViolation = "23503"
)

func CheckUnique(err error, entity string) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == CodeUniqueViolation {
		switch entity {
		case "user":
			return fmt.Errorf("%w: user already exists", customErrors.ErrUserNotFound)
		case "team":
			return fmt.Errorf("%w: team already exists", customErrors.ErrTeamAlreadyExists)
		case "pull_request":
			return fmt.Errorf("%w: PR id already exists", customErrors.ErrPRAlreadyExists)
		default:
			return fmt.Errorf("unique violation on %s", entity)
		}
	}
	return err
}

func CheckForeignKey(err error, entity string) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == CodeForeignKeyViolation {
		switch entity {
		case "team":
			return fmt.Errorf("%w: referenced team not found", customErrors.ErrTeamNotFound)
		case "user":
			return fmt.Errorf("%w: referenced user not found", customErrors.ErrUserNotFound)
		default:
			return fmt.Errorf("foreign key violation on %s", entity)
		}
	}
	return err
}
