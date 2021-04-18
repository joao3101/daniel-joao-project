package athletes

import (
	errReturn "github.com/joao3101/daniel-joao-project/api/error"
	httpReq "github.com/joao3101/daniel-joao-project/app/athletes/http"
	"github.com/joao3101/daniel-joao-project/models"
	config "github.com/joao3101/daniel-joao-project/util"
	"github.com/labstack/echo/v4"
)

func PostAthlete(config *config.Config) echo.HandlerFunc {
	return func(c echo.Context) error {

		var postReq httpReq.AthleteReq
		err := c.Bind(&postReq)
		if err != nil {
			return errReturn.ResponseInternalServerError(c)
		}
		model, err := models.InitModel(config.DBDriver, config.DBSource)
		if err != nil {
			return errReturn.ResponseInternalServerError(c)
		}

		response, err := CreateAthlete(model, postReq)
		if err != nil {
			return errReturn.ResponseInternalServerError(c)
		}

		//TODO fazer a lógica da parada aqui que chama o useCase
		return c.JSON(200, errReturn.BaseResponse{
			Data:    response,
			Success: 1,
			Message: "",
		})

	}
}
