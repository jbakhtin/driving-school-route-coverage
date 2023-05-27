package query

const (
	CreateUser = `
		INSERT INTO users (name, lastname, login, email, password, created_at)
		VALUES($1, $2, $3, $4, $5, NOW())
		RETURNING id, name, lastname, login, email, password, created_at, updated_at
	`

	UpdateUser = `
		INSERT INTO users (name, lastname, login, email, password, updated_at)
		VALUES($1, $2, $3, $4, $5, NOW())
		RETURNING id, name, lastname, login, email, password, created_at, updated_at
	`

	GetUserByID = `
		SELECT id, name, lastname, login, email, created_at, updated_at FROM users
		WHERE users.id = $1
	`

	GetUsers = `
		SELECT id, name, lastname, login, email, created_at, updated_at FROM users
	`
)
