package query

const (
	CreateRoute = `
		INSERT INTO routes (line, created_at)
		VALUES(ST_AsGeoJSON($1), NOW())
		RETURNING id, ST_AsBinary(line), created_at, updated_at
	`

	GetRouteByID = `
		SELECT Id, ST_AsBinary(line) AS line , created_at, updated_at FROM routes WHERE routes.id = $1
	`

	GetRoutes = `
		SELECT * FROM routes
	`

	UpdateRouteByID = `
		UPDATE routes
		SET line = ST_AsGeoJSON($2), updated_at = now()
		WHERE id = $1
		RETURNING id, ST_AsBinary(line), created_at, updated_at
	`
)