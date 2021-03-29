package middleware

import (
	"math/rand"
	"time"

	"github.com/labstack/echo/v4"
)

// Sloth is the middleware function
func Sloth(fixedDelay int, delta int) func(next echo.HandlerFunc) echo.HandlerFunc {
	if delta > fixedDelay {
		delta = fixedDelay
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {

		delay := (fixedDelay + (rand.Intn(delta*2) - delta)) * 1000

		time.Sleep(time.Duration(delay) * time.Microsecond)

		return func(c echo.Context) error {
			return next(c)
		}
	}

}
