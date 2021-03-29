package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SimpleHealthCheck(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
