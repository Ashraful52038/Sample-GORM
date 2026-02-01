package middleware

import (
	"github.com/labstack/echo/v4/middleware"
)

func GzipConfig() middleware.GzipConfig {
	return middleware.GzipConfig{
		Level: 5,
	}
}

func DecompressConfig() middleware.DecompressConfig {
	return middleware.DecompressConfig{}
}
