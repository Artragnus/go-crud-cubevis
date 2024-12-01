package main

import (
	"context"
	"database/sql"
	"github.com/Artragnus/go-crud-cubevis/db"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type ProductOutput struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Value int32  `json:"value"`
}

func GetProductByIdHandler(c echo.Context) error {
	id := c.Param("id")

	num, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid id",
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

	product, err := q.GetProductById(context.Background(), int32(num))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Internal server error",
		})
	}

	output := ProductOutput{
		ID:    product.ID,
		Name:  product.Name,
		Value: product.Value,
	}

	return c.JSON(http.StatusOK, output)

}

func GetProductsHandler(c echo.Context) error {
	conn, err := sql.Open("postgres", env.DataSourceName)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Internal server error",
		})
	}

	defer conn.Close()

	q := db.New(conn)

	products, err := q.GetProducts(context.Background())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Internal server error",
		})
	}

	output := make([]ProductOutput, len(products))

	for i, p := range products {
		output[i] = ProductOutput{
			ID:    p.ID,
			Name:  p.Name,
			Value: p.Value,
		}
	}

	return c.JSON(http.StatusOK, output)

}
