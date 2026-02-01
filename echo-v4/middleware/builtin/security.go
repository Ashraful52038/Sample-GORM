package middleware

import (
	"github.com/labstack/echo/v4/middleware"
)

func SecurityConfig() middleware.SecureConfig {
	return middleware.SecureConfig{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "DENY",
		HSTSMaxAge:            31536000,
		HSTSExcludeSubdomains: false,
		ContentSecurityPolicy: "default-src 'self'; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline'",
		CSPReportOnly:         false,
		ReferrerPolicy:        "strict-origin-when-cross-origin",
	}
}

func BodyLimitConfig() string {
	return "10M"
}
