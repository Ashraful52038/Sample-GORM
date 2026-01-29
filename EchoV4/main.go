package main

import (
	"net/http"

	"EchoV4/handlers"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Custom Validator
type CustomValidator struct {
	validator *validator.Validate
}

// Validate implements echo.Validator
func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return err
	}
	return nil
}

func main() {
	e := echo.New()

	// Initialize validator
	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo V4!")
	})

	//All user route
	e.POST("/users", handlers.CreateUser)
	e.GET("/users/:id", handlers.GetUser)
	e.PUT("/users/:id", handlers.UpdateUser)
	e.DELETE("/users/:id", handlers.DeleteUser)
	e.GET("/users", handlers.GetAllUsers)

	if err := e.Start(":8080"); err != nil {
		e.Logger.Fatal(err)
	}
}
