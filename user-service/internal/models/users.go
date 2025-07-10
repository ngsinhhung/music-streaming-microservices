package models

type Users struct {
	id       int
	name     string
	email    string
	phone    string
	password string
	roles    []Roles
}
