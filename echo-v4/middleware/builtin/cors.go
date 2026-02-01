package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CORSConfig() middleware.CORSConfig {
	return middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:3000",  // Frontend development
			"http://localhost:5173",  // Vite dev server
			"https://yourdomain.com", // Production
			"https://www.yourdomain.com",
		},
		AllowMethods: []string{
			echo.GET,
			echo.POST,
			echo.PUT,
			echo.DELETE,
			echo.PATCH,
			echo.OPTIONS,
			echo.HEAD,
		},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
			"X-Requested-With",
			"X-CSRF-Token",
			"X-API-Key",
		},
		ExposeHeaders: []string{
			"Content-Length",
			"Content-Disposition",
			"X-Total-Count",
		},
		AllowCredentials: true,
		MaxAge:           86400, // 24 hours
	}
}
