// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package internal

import (
	"database/sql"
)

type Address struct {
	ID      string
	UserID  string
	Address string
	Number  sql.NullString
	ZipCode sql.NullString
	City    string
	State   string
}

type User struct {
	ID       string
	Name     sql.NullString
	Email    string
	Password string
}
