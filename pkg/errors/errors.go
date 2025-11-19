package errors

import "errors"

var (
	ErrPRNotFound      = errors.New("PR_NOT_FOUND")
	ErrPRAlreadyMerged = errors.New("PR_MERGED")
	ErrPRAlreadyExists = errors.New("PR_EXISTS")
	ErrNoCandidate     = errors.New("NO_CANDIDATE")
	ErrNotAssigned     = errors.New("NOT_ASSIGNED")

	ErrTeamAlreadyExists = errors.New("TEAM_EXISTS")
	ErrTeamNotFound      = errors.New("TEAM_NOT_FOUND")
	ErrUserNotFound      = errors.New("USER_NOT_FOUND")
)
