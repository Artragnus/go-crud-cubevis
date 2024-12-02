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
	"strconv"
)

type CreateOrderInput struct {
	ProductId int32     `json:"product_id"`
	Quantity  int32     `json:"quantity"`
	AddressId uuid.UUID `json:"address_id"`
}

type CreateOrderOutput struct {
	ID         uuid.UUID `json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	ProductID  int32     `json:"product_id"`
	Quantity   int32     `json:"quantity"`
	AddressID  uuid.UUID `json:"address_id"`
	TotalValue int32     `json:"total_value"`
}

type GetOrderByIdOutput struct {
	ID         uuid.UUID `json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	ProductID  int32     `json:"product_id"`
	Quantity   int32     `json:"quantity"`
	AddressID  uuid.UUID `json:"address_id"`
	TotalValue int32     `json:"total_value"`
}

type GetOrdersOutput struct {
	ID         uuid.UUID `json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	ProductID  int32     `json:"product_id"`
	Quantity   int32     `json:"quantity"`
	AddressID  uuid.UUID `json:"address_id"`
	TotalValue int32     `json:"total_value"`
}

type GetOrdersByProductOutput struct {
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	TotalValue int32     `json:"total_value"`
	Quantity   int32     `json:"quantity"`
	OrderID    uuid.UUID `json:"order_id"`
	State      string    `json:"state"`
	Address    string    `json:"address"`
	Number     string    `json:"number"`
	ZipCode    string    `json:"zip_code"`
	City       string    `json:"city"`
}

func CreateOrderHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)

	body := new(CreateOrderInput)
	err := c.Bind(&body)

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

	defer conn.Close()

	q := db.New(conn)

	product, err := q.GetProductById(context.Background(), body.ProductId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid product id",
		})
	}

	_, err = q.GetAddressById(context.Background(), db.GetAddressByIdParams{
		ID:     body.AddressId,
		UserID: claims.ID,
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid address id",
		})
	}

	if body.Quantity < 1 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid quantity",
		})
	}

	params := db.CreateOrderParams{
		ID:         uuid.New(),
		UserID:     claims.ID,
		ProductID:  body.ProductId,
		Quantity:   body.Quantity,
		AddressID:  body.AddressId,
		TotalValue: product.Value * body.Quantity,
	}

	err = q.CreateOrder(context.Background(), params)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Internal server error",
		})
	}

	output := CreateOrderOutput{
		ID:         params.ID,
		UserID:     params.UserID,
		ProductID:  params.ProductID,
		Quantity:   params.Quantity,
		AddressID:  params.AddressID,
		TotalValue: params.TotalValue,
	}

	return c.JSON(http.StatusCreated, output)

}

func GetOrderByIdHandler(c echo.Context) error {
	users := c.Get("user").(*jwt.Token)
	claims := users.Claims.(*JwtCustomClaims)

	id := c.Param("id")
	parseUUID, err := uuid.Parse(id)

	if err != nil {
		fmt.Println(id)
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid order id",
		})
	}

	conn, err := sql.Open("postgres", env.DataSourceName)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Internal server error",
		})
	}

	defer conn.Close()

	q := db.New(conn)

	order, err := q.GetOrderById(context.Background(), db.GetOrderByIdParams{
		ID:     parseUUID,
		UserID: claims.ID,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Internal server error",
		})
	}

	output := GetOrderByIdOutput{
		ID:         order.ID,
		UserID:     order.UserID,
		ProductID:  order.ProductID,
		Quantity:   order.Quantity,
		AddressID:  order.AddressID,
		TotalValue: order.TotalValue,
	}

	return c.JSON(http.StatusOK, output)

}

func GetDetailedOrderByIdHandler(c echo.Context) error {
	users := c.Get("user").(*jwt.Token)
	claims := users.Claims.(*JwtCustomClaims)

	id := c.Param("id")
	parseUUID, err := uuid.Parse(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid order id",
		})
	}

	conn, err := sql.Open("postgres", env.DataSourceName)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Internal server error",
		})
	}

	defer conn.Close()

	q := db.New(conn)

	order, err := q.GetDeitaledOrderById(context.Background(), db.GetDeitaledOrderByIdParams{
		ID:     parseUUID,
		UserID: claims.ID,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Internal server error",
		})
	}

	return c.JSON(http.StatusOK, order)
}

func GetOrdersHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)

	conn, err := sql.Open("postgres", env.DataSourceName)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Internal server error",
		})
	}

	defer conn.Close()

	q := db.New(conn)

	orders, err := q.GetOrders(context.Background(), claims.ID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Internal server error",
		})
	}

	output := make([]GetOrdersOutput, len(orders))

	for i, o := range orders {
		output[i] = GetOrdersOutput{
			ID:         o.ID,
			UserID:     o.UserID,
			ProductID:  o.ProductID,
			Quantity:   o.Quantity,
			AddressID:  o.AddressID,
			TotalValue: o.TotalValue,
		}
	}

	return c.JSON(http.StatusOK, output)

}

func GetOrdersByProductIdHandler(c echo.Context) error {
	id := c.Param("id")

	num, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid product id",
		})
	}

	conn, err := sql.Open("postgres", env.DataSourceName)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Internal server error",
		})
	}

	defer conn.Close()

	q := db.New(conn)

	result, err := q.GetUsersByProduct(context.Background(), int32(num))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Internal server error",
		})
	}

	if len(result) == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "No orders found",
		})
	}

	output := make([]GetOrdersByProductOutput, len(result))

	for i, r := range result {
		output[i] = GetOrdersByProductOutput{
			Name:       r.Name,
			Email:      r.Email,
			TotalValue: r.TotalValue,
			Quantity:   r.Quantity,
			OrderID:    r.OrderID,
			State:      r.State,
			Address:    r.Address,
			Number:     r.Number,
			ZipCode:    r.ZipCode,
			City:       r.City,
		}
	}

	return c.JSON(http.StatusOK, output)
}
