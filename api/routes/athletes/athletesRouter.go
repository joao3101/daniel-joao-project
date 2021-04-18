package athletes

import (
	athletes "github.com/joao3101/daniel-joao-project/app/athletes"
	config "github.com/joao3101/daniel-joao-project/util"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, config *config.Config) {

	r := e.Group("/athletes")

	r.POST("/", athletes.PostAthlete(config))
}
