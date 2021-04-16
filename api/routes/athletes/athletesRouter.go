package athletes

import (
	"github.com/joao3101/daniel-joao-project/api/config"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, config *config.Config) {

	r := e.Group("/athletes")

	r.POST("/", config.PostAthlete)
}
