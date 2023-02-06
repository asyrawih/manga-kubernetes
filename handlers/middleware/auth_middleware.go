package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/o1egl/paseto"
)

func AuthMiddleware(key string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set(echo.HeaderServer, "Manga/v1.0")
			h := c.Request().Header
			authHeader := h.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, "Missing Authorization Header")
			}
			s := strings.Split(authHeader, " ")

			if s[0] != "Bearer" {
				return c.JSON(http.StatusUnauthorized, "Missing Authorization Header")
			}

			if err := paseto.NewV2().Decrypt(s[1], []byte(key), nil, nil); err != nil {
				return c.JSON(http.StatusUnauthorized, "Invalid Authorization")
			}

			return next(c)
		}
	}
}
