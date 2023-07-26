package query

const (
	CreateRoute = `
		INSERT INTO routes (user_id, name, description, linestring, created_at)
		VALUES($1, $2, $3, ST_AsGeoJSON($4), NOW())
		RETURNING id, user_id, name, description, ST_AsBinary(linestring), created_at, updated_at
	`

	GetRouteByID = `
		SELECT Id, user_id, name, description, ST_AsBinary(linestring) AS linestring , created_at, updated_at 
		FROM routes 
		WHERE routes.id = $1 AND routes.user_id = $2
	`

	GetRoutes = `
		SELECT Id, user_id, name, description, created_at, updated_at  FROM routes WHERE user_id = $1 LIMIT $2
	`

	UpdateRouteByID = `
		UPDATE routes
		SET name = $3, linestring = ST_AsGeoJSON($4), updated_at = now()
		WHERE id = $1 AND user_id = $2
		RETURNING id, user_id, description, name, ST_AsBinary(linestring), created_at, updated_at
	`

	DeleteRouteByID = `
		DELETE FROM routes
		WHERE id = $1 AND user_id = $2
	`
)
