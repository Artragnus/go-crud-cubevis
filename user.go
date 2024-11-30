package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Artragnus/go-crud-cubevis/db"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type UserEntity struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserBody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUser(name, email, password string) (*UserEntity, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	return &UserEntity{
		ID:       uuid.New().String(),
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}, nil
}

func createUserHandler(c echo.Context) error {
	body := new(CreateUserBody)

	err := c.Bind(&body)

	if err != nil {
		fmt.Println(err)
	}

	u := NewUser(body.Name, body.Email, body.Password)

	ctx := context.Background()

	conn, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")

	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()

	q := db.New(conn)

	err = q.CreateUser(ctx, db.CreateUserParams{
		ID:       u.ID,
		Name:     sql.NullString{String: u.Name, Valid: true},
		Email:    u.Email,
		Password: u.Password,
	})

	if err != nil {
		fmt.Println(err)
	}

	return c.JSON(http.StatusCreated, u)

}
