package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Artragnus/go-crud-cubevis/db"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Address struct {
	ID      uuid.UUID `json:"id"`
	UserID  uuid.UUID `json:"user_id"`
	Address string    `json:"address"`
	Number  string    `json:"number"`
	ZipCode string    `json:"zip_code"`
	City    string    `json:"city"`
	State   string    `json:"state"`
}

func NewAddress(userId uuid.UUID, address, number, zipCode, city, state string) (*Address, error) {

	a := Address{
		ID:      uuid.New(),
		UserID:  userId,
		Address: address,
		Number:  number,
		ZipCode: zipCode,
		City:    city,
		State:   state,
	}

	err := a.Validate()

	if err != nil {
		return nil, err
	}

	return &a, nil

}

func (a *Address) Validate() error {
	if a.Address == "" {
		return fmt.Errorf("address is required")
	}
	if a.Number == "" {
		return fmt.Errorf("number is required")
	}
	if a.ZipCode == "" {
		return fmt.Errorf("zip code is required")
	}
	if a.City == "" {
		return fmt.Errorf("city is required")
	}
	if a.State == "" {
		return fmt.Errorf("state is required")
	}

	return nil
}

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

	a, err := NewAddress(claims.ID, body.Address, body.Number, body.ZipCode, body.City, body.State)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	conn, err := sql.Open("postgres", env.DataSourceName)

	if err != nil {
		panic(err)
	}

	q := db.New(conn)

	err = q.CreateAddress(context.Background(), db.CreateAddressParams{
		ID:      a.ID,
		UserID:  a.UserID,
		Address: a.Address,
		Number:  a.Number,
		ZipCode: a.ZipCode,
		City:    a.City,
		State:   a.State,
	})

	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusCreated, a)

}

func GetAddressesHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)

	conn, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	q := db.New(conn)

	addresses, err := q.GetAddresses(context.Background(), claims.ID)

	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, addresses)

}

func GetAddressByIdHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)

	id := c.Param("id")

	conn, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	q := db.New(conn)

	parseUUID, err := uuid.Parse(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "Address not found",
		})
	}

	address, err := q.GetAddressById(context.Background(), db.GetAddressByIdParams{
		ID:     parseUUID,
		UserID: claims.ID,
	})

	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, address)

}
