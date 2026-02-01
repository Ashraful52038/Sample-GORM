package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "healthy",
		"service": "go-backend",
		"version": "1.0.0",
	})
}
