package query

const (
	CreateRoute = `
		INSERT INTO routes (line, created_at)
		VALUES(ST_GeomFromWKB($1), NOW())
		RETURNING id, ST_AsBinary(line), created_at, updated_at
	`
)