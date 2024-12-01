package main

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

var env *Conf

func main() {
	var err error

	env, err = LoadConfig(".")

	if err != nil {
		panic(err)
	}

	e := echo.New()

	e.POST("/user", CreateUserHandler)
	e.POST("/login", LoginHandler)
	e.GET("/product/:id", GetProductByIdHandler)
	e.GET("/product", GetProductsHandler)
	e.GET("/order/product/:id", GetOrdersByProductHandler)

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JwtCustomClaims)
		},
		SigningKey: []byte(env.JWTSecret),
	}

	r := e.Group("/auth")

	r.Use(echojwt.WithConfig(config))

	r.PUT("/user", UpdateUserHandler)
	r.DELETE("/user", DeleteUserHandler)

	r.POST("/address", CreateAddressHandler)
	r.GET("/address/:id", GetAddressByIdHandler)
	r.PUT("/address/:id", UpdateAddressHandler)
	r.GET("/address", GetAddressesHandler)
	r.DELETE("/address/:id", DeleteAddressHandler)

	r.POST("/order", CreateOrderHandler)
	r.GET("/order/:id", GetOrderByIdHandler)
	r.GET("/order/:id/details", GetDetailedOrderByIDHandler)
	r.GET("/order", GetOrdersHandler)

	e.Logger.Fatal(e.Start(":" + env.Port))
}
