package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Setup(e *echo.Echo) {
	// 1. Recovery - Should be first to recover from panics
	e.Use(middleware.Recover())

	// 2. CORS - Handle cross-origin requests
	e.Use(middleware.CORSWithConfig(CORSConfig()))

	// 3. Logging - Log all requests
	e.Use(middleware.LoggerWithConfig(LoggerConfig()))

	// 4. Body Limit - Prevent large payloads
	e.Use(middleware.BodyLimit(BodyLimitConfig()))

	// 5. Decompress - Decompress incoming requests
	e.Use(middleware.DecompressWithConfig(DecompressConfig()))

	// 6. Compression - Compress responses
	e.Use(middleware.GzipWithConfig(GzipConfig()))

	// 7. Security Headers
	e.Use(middleware.SecureWithConfig(SecurityConfig()))

	// 8. Request ID - Add unique ID to each request
	e.Use(middleware.RequestID())
}

// Development setup (less strict)
func SetupDevelopment(e *echo.Echo) {
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS()) // Allow all origins in dev
}

// Production setup (strict)
func SetupProduction(e *echo.Echo) {
	Setup(e)
}
