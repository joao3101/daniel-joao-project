package athletes

import (
	"github.com/joao3101/daniel-joao-project/api/config"
	"github.com/labstack/echo/v4"
)

func PostAthlete(config *config.Config) echo.HandlerFunc {
	return func(c echo.Context) error {

		//TODO fazer a lógica da parada aqui que chama o useCase
		return nil
	}
}
