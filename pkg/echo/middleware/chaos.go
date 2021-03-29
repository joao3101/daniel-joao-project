package middleware

import (
	"math/rand"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Chaos is the middleware function.
func Chaos(chance int) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		if rand.Intn(100) > 100-chance {
			return func(c echo.Context) error {
				return c.HTML(http.StatusInternalServerError, "Chaos Middleware")
			}
		}

		return func(c echo.Context) error {
			return next(c)
		}
	}
}
