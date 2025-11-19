package prrepo

const (
	createPRQuery = `
	INSERT INTO pull_requests (pull_request_id, pull_request_name, author_id, status, created_at)
	VALUES ($1, $2, $3, $4, $5);
	`

	getPRByIDQuery = `
	SELECT pull_request_id, pull_request_name, author_id, status, created_at, merged_at
	FROM pull_requests
	WHERE pull_request_id = $1;
	`

	listPRsByUserQuery = `
	SELECT pull_request_id, pull_request_name, author_id, status
	FROM pull_requests
	WHERE author_id = $1;
	`

	updatePRStatusQuery = `
	UPDATE pull_requests
	SET status = $2, merged_at = $3
	WHERE pull_request_id = $1;
	`

	insertReviewerQuery = `
	INSERT INTO pull_request_reviewers (pull_request_id, reviewer_id)
	VALUES ($1, $2)
	ON CONFLICT (pull_request_id, reviewer_id) DO NOTHING;
	`

	deleteReviewerQuery = `
	DELETE FROM pull_request_reviewers
	WHERE pull_request_id = $1 AND reviewer_id = $2;
	`

	getReviewersQuery = `
	SELECT reviewer_id
	FROM pull_request_reviewers
	WHERE pull_request_id = $1;
	`

	getAssignmentsCountQuery = `
	SELECT reviewer_id, COUNT(*) as count
    FROM pull_request_reviewers
    GROUP BY reviewer_id;
	`
)
