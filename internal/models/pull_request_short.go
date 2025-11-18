package models

type PullRequestShort struct {
	PullRequestId   string `db:"pull_request_id"`
	PullRequestName string `db:"pull_request_name"`
	AuthorId        string `db:"author_id"`
	Status          string `db:"status"`
}
