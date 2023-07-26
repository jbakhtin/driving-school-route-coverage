package models

type User struct {
	ID        int
	Name      string
	Lastname  string
	Login     string
	Email     string
	Password  string
	CreatedAt string
	UpdatedAt *string
}
