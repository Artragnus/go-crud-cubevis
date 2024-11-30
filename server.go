package main

import (
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	e := echo.New()

	e.POST("/users", createUserHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
