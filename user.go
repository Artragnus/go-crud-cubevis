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
	if password == "" {
		return nil, fmt.Errorf("password is required")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	user := User{
		ID:       uuid.New(),
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	err = user.Validate()

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) ValidatePassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	if err != nil {
		return err
	}

	return nil
}

func (u *User) Validate() error {
	if u.Name == "" {
		return fmt.Errorf("name is required")
	}

	if u.Email == "" {
		return fmt.Errorf("email is required")
	}

	if u.Password == "" {
		return fmt.Errorf("password is required")
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
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	conn, err := sql.Open("postgres", env.DataSourceName)

	if err != nil {

		fmt.Println(err)
	}

	defer conn.Close()

	q := db.New(conn)

	_, err = q.GetUserByEmail(context.Background(), u.Email)

	if err == nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Email already in use",
		})
	}

	err = q.CreateUser(context.Background(), db.CreateUserParams{
		ID:       u.ID,
		Name:     u.Name,
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

	conn, err := sql.Open("postgres", env.DataSourceName)

	defer conn.Close()

	q := db.New(conn)

	u, err := q.GetUserByEmail(context.Background(), body.Email)

	if err != nil {
		panic(err)
	}

	user := User{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}

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

	t, err := token.SignedString([]byte(env.JWTSecret))

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

	body := new(UpdateUserInput)

	err := c.Bind(&body)

	if err != nil {
		fmt.Println(err)
	}

	var hashedPassword []byte

	if body.Password != "" {
		hashedPassword, err = bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

		if err != nil {
			panic(err)
		}
	}

	params := db.UpdateUserParams{
		ID:       claims.ID,
		Name:     body.Name,
		Email:    body.Email,
		Password: string(hashedPassword),
	}

	conn, err := sql.Open("postgres", env.DataSourceName)

	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()

	q := db.New(conn)

	u, err := q.GetUserById(context.Background(), claims.ID)

	if err != nil {
		panic(err)
	}

	if params.Name == "" {
		params.Name = u.Name
	}

	if params.Email == "" {
		params.Email = u.Email
	}

	if params.Password == "" {
		params.Password = u.Password
	}

	err = q.UpdateUser(context.Background(), db.UpdateUserParams{
		ID:       claims.ID,
		Name:     body.Name,
		Email:    body.Email,
		Password: string(hashedPassword),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Internal server error",
		})
	}

	return c.JSON(http.StatusNoContent, nil)
}

func DeleteUserHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)

	conn, err := sql.Open("postgres", env.DataSourceName)

	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()

	q := db.New(conn)

	err = q.DeleteUser(context.Background(), claims.ID)

	if err != nil {
		fmt.Println(err)
	}

	return c.JSON(http.StatusNoContent, nil)
}
