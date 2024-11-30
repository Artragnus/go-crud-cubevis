package main

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	e := echo.New()

	e.POST("/users", CreateUserHandler)
	e.POST("/login", LoginHandler)

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}

	r := e.Group("/auth")

	r.Use(echojwt.WithConfig(config))

	r.PUT("/users", UpdateUserHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
