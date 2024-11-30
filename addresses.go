package main

import (
	"context"
	"database/sql"
	"github.com/Artragnus/go-crud-cubevis/db"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CreateAddressInput struct {
	Address string `json:"address"`
	Number  string `json:"number"`
	ZipCode string `json:"zip_code"`
	City    string `json:"city"`
	State   string `json:"state"`
}

func CreateAddressHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)

	body := new(CreateAddressInput)

	err := c.Bind(body)
	if err != nil {
		panic(err)
	}

	address := db.Address{
		ID:      uuid.New(),
		UserID:  claims.ID,
		Address: body.Address,
		Number:  body.Number,
		ZipCode: body.ZipCode,
		City:    body.City,
		State:   body.State,
	}

	conn, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")

	if err != nil {
		panic(err)
	}

	q := db.New(conn)

	err = q.CreateAddress(context.Background(), db.CreateAddressParams{
		ID:      address.ID,
		UserID:  address.UserID,
		Address: address.Address,
		Number:  address.Number,
		ZipCode: address.ZipCode,
		City:    address.City,
		State:   address.State,
	})

	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusCreated, address)

}
