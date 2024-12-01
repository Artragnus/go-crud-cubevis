package main

import (
	"context"
	"database/sql"
	"github.com/Artragnus/go-crud-cubevis/db"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

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

	return c.JSON(http.StatusOK, product)

}
