package userrepo

const (
	createUserQuery = `
	INSERT INTO users (user_id, username, team_name, is_active)
	VALUES ($1, $2, $3, $4)
	ON CONFLICT (user_id) DO NOTHING;
	`

	setActiveQuery = `
	UPDATE users
	SET is_active = $2
	WHERE user_id = $1;
	`

	getByIDQuery = `
	SELECT user_id, username, team_name, is_active
	FROM users
	WHERE user_id = $1;
	`

	getTeamMembersQuery = `
	SELECT user_id, username, team_name, is_active
	FROM users
	WHERE team_name = $1;
	`
)
