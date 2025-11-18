package models

type TeamMember struct {
	UserId   string `db:"user_id"`
	Username string `db:"username"`
	IsActive bool   `db:"is_active"`
}
