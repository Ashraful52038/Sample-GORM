package main

import (
	"net/http"

	"EchoV4/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo V4!")
	})

	e.POST("/users", handlers.CreateUser)

	if err := e.Start(":8080"); err != nil {
		e.Logger.Fatal(err)
	}
}
