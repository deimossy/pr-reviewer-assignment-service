-- +goose Up
-- +goose StatementBegin

CREATE TABLE teams (
    team_name TEXT PRIMARY KEY
);

CREATE TABLE users (
    user_id TEXT PRIMARY KEY,
    username TEXT NOT NULL,
    team_name TEXT REFERENCES teams(team_name) ON DELETE CASCADE,
    is_active BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE TABLE pull_requests (
    pull_request_id TEXT PRIMARY KEY,
    pull_request_name TEXT NOT NULL,
    author_id TEXT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    status TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    merged_at TIMESTAMPTZ
);

CREATE TABLE pull_request_reviewers (
    pull_request_id TEXT NOT NULL REFERENCES pull_requests(pull_request_id) ON DELETE CASCADE,
    reviewer_id TEXT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    PRIMARY KEY (pull_request_id, reviewer_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS pull_request_reviewers;
DROP TABLE IF EXISTS pull_requests;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS teams;

-- +goose StatementEnd
