package middleware

import (
	"travel-backend/config"
	"travel-backend/pkg/utils"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// InitJWTMiddleware mengembalikan middleware auth bawaan Echo yang sudah dikonfigurasi
func InitJWTMiddleware(cfg *config.Config) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(utils.JwtCustomClaims)
		},
		SigningKey: []byte(cfg.JWTSecret),
		// Opsional: Custom error response jika token invalid/expired
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(401, map[string]interface{}{
				"success": false,
				"message": "Unauthorized or token invalid",
			})
		},
	})
}