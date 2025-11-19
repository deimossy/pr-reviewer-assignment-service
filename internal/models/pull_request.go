package models

import "time"

type PullRequest struct {
	PullRequestId     string     `db:"pull_request_id"`
	PullRequestName   string     `db:"pull_request_name"`
	AuthorId          string     `db:"author_id"`
	Status            string     `db:"status"`
	CreatedAt         *time.Time `db:"created_at"`
	MergedAt          *time.Time `db:"merged_at,omitempty"`
	AssignedReviewers []string   `db:"-"`
}
