package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// TODO: parser subnets net.ParseCIDR(ip)

//IPBlacklist ...
func IPBlacklist(blockedIPs map[string]bool) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			ip := c.RealIP()
			_, exists := blockedIPs[ip]

			if exists {
				return c.NoContent(http.StatusForbidden)
			}

			return next(c)
		}
	}
}
