package teamrepo

const (
	createTeamQuery = `
	INSERT INTO teams (team_name)
	VALUES ($1)
	`

	getByNameQuery = `
	SELECT team_name
	FROM teams
	WHERE team_name = $1;
	`
)
