package main

import (
	"net/http"

	"echo-v4-app/handlers"
	builtin "echo-v4-app/middleware/builtin"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
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

	// e.Use(middleware.Cors)
	// e.Use(middleware.LoggingMiddleware)
	// e.Use(echoMw.Recover())

	// Setup all middlewares
	builtin.Setup(e)

	// Routes
	setupRoutes(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo V4!")
	})

	// All user routes
	e.POST("/users", handlers.CreateUser)
	e.GET("/users/:id", handlers.GetUser)
	e.PUT("/users/:id", handlers.UpdateUser)
	e.DELETE("/users/:id", handlers.DeleteUser)
	e.GET("/users", handlers.GetAllUsers)

	if err := e.Start(":8080"); err != nil {
		e.Logger.Fatal(err)
	}
}

func setupRoutes(e *echo.Echo) {
	// Health check
	e.GET("/health", handlers.HealthCheck)

	// API v1 routes
	api := e.Group("/api/v1")

	// Public routes
	api.POST("/users", handlers.CreateUser)

	// Protected routes
	protected := api.Group("")
	// protected.Use(middleware.JWTMiddleware())
	protected.PUT("/users/:id", handlers.UpdateUser)
	protected.DELETE("/users/:id", handlers.DeleteUser)
}
