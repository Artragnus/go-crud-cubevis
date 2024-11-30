package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Artragnus/go-crud-cubevis/db"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JwtCustomClaims struct {
	ID uuid.UUID `json:"id"`
	jwt.RegisteredClaims
}

func NewUser(name, email, password string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	return &User{

		ID:       uuid.New(),
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}, nil
}

func ToEntity(user db.User) *User {
	return &User{
		ID:       user.ID,
		Name:     user.Name.String,
		Email:    user.Email,
		Password: user.Password,
	}
}

func (u *User) ValidatePassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	if err != nil {
		return err
	}

	return nil
}

func CreateUserHandler(c echo.Context) error {
	body := new(CreateUserInput)

	err := c.Bind(&body)

	if err != nil {
		fmt.Println(err)
	}

	u, err := NewUser(body.Name, body.Email, body.Password)

	if err != nil {
		fmt.Println(err)
	}

	conn, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")

	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()

	q := db.New(conn)

	err = q.CreateUser(context.Background(), db.CreateUserParams{
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

func LoginHandler(c echo.Context) error {
	body := new(LoginInput)

	err := c.Bind(&body)

	if err != nil {
		fmt.Println(err)
	}

	conn, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")

	defer conn.Close()

	q := db.New(conn)

	data, err := q.GetUserByEmail(context.Background(), body.Email)

	user := ToEntity(data)

	err = user.ValidatePassword(body.Password)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "Invalid email or password",
		})
	}

	claims := &JwtCustomClaims{
		user.ID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))

	if err != nil {
		fmt.Println(err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})

}

func UpdateUserHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)

	conn, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")

	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()

	body := new(UpdateUserInput)

	err = c.Bind(&body)

	if err != nil {
		fmt.Println(err)
	}

	q := db.New(conn)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println(err)
	}

	err = q.UpdateUser(context.Background(), db.UpdateUserParams{
		ID:       claims.ID,
		Name:     sql.NullString{String: body.Name, Valid: len(body.Name) > 0},
		Email:    body.Email,
		Password: string(hashedPassword),
	})

	if err != nil {
		fmt.Println(err)
	}

	return nil
}
