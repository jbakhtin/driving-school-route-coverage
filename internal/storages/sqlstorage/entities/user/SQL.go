package user

const (
	create = `
		INSERT INTO users (name, lastname, login, email, password, created_at)
		VALUES($1, $2, $3, $4, $5, $6)
		RETURNING id, name, lastname, login, email, password, created_at, updated_at
	`

	update = `
		INSERT INTO users (name, lastname, login, email, password, created_at)
		VALUES($1, $2, $3, $4, $5, $6)
		RETURNING id, name, lastname, login, email, password
	`

	getByID = `
		SELECT id, name, lastname, login, email, created_at, updated_at FROM users
		WHERE users.id = $1
	`

	getAll = `
		SELECT id, name, lastname, login, email, created_at, updated_at FROM users
	`
)
