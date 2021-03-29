package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

//ValidateToken ...
func ValidateToken(token string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")

			if authHeader == token {
				return next(c)
			}
			return c.NoContent(http.StatusUnauthorized)
		}
	}
}
