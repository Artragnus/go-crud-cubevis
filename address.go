package main

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Artragnus/go-crud-cubevis/db"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

var (
	ErrAddressIsRequired = errors.New("address is required")
	ErrNumberIsRequired  = errors.New("number is required")
	ErrZipCodeIsRequired = errors.New("zip code is required")
	ErrCityIsRequired    = errors.New("city is required")
	ErrStateIsRequired   = errors.New("state is required")
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
		return ErrAddressIsRequired
	}
	if a.Number == "" {
		return ErrNumberIsRequired
	}
	if a.ZipCode == "" {
		return ErrZipCodeIsRequired
	}
	if a.City == "" {
		return ErrCityIsRequired
	}
	if a.State == "" {
		return ErrStateIsRequired
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

type UpdateAddressInput struct {
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
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Internal server error",
		})
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
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Internal server error",
		})
	}

	return c.JSON(http.StatusCreated, a)

}

func GetAddressByIdHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)

	id := c.Param("id")

	conn, err := sql.Open("postgres", env.DataSourceName)

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
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "Address not found",
		})
	}

	return c.JSON(http.StatusOK, address)
}

func GetAddressesHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)

	conn, err := sql.Open("postgres", env.DataSourceName)

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

func UpdateAddressHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)

	id := c.Param("id")

	body := new(UpdateAddressInput)

	err := c.Bind(body)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid request body",
		})
	}

	conn, err := sql.Open("postgres", env.DataSourceName)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Internal server error",
		})
	}

	q := db.New(conn)

	a, err := q.GetAddressById(context.Background(), db.GetAddressByIdParams{
		ID:     uuid.MustParse(id),
		UserID: claims.ID,
	})

	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "Address not found",
		})
	}

	params := db.UpdateAddressParams{
		ID:      a.ID,
		UserID:  a.UserID,
		Address: body.Address,
		Number:  body.Number,
		ZipCode: body.ZipCode,
		City:    body.City,
		State:   body.State,
	}

	if params.Address == "" {
		params.Address = a.Address
	}

	if params.City == "" {
		params.City = a.City
	}

	if params.Number == "" {
		params.Number = a.Number
	}

	if params.State == "" {
		params.State = a.State
	}

	if params.ZipCode == "" {
		params.ZipCode = a.ZipCode
	}

	err = q.UpdateAddress(context.Background(), params)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Internal server error",
		})
	}

	return c.JSON(http.StatusNoContent, params)

}

func DeleteAddressHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	id := c.Param("id")

	conn, err := sql.Open("postgres", env.DataSourceName)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Internal server error",
		})
	}

	defer conn.Close()

	q := db.New(conn)

	err = q.DeleteAddress(context.Background(), db.DeleteAddressParams{
		ID:     uuid.MustParse(id),
		UserID: claims.ID,
	})

	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "Address not found",
		})
	}

	return c.JSON(http.StatusNoContent, nil)
}
