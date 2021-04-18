package athletes

import (
	"github.com/jinzhu/copier"
	httpReq "github.com/joao3101/daniel-joao-project/app/athletes/http"
	"github.com/joao3101/daniel-joao-project/app/domain"
	"github.com/joao3101/daniel-joao-project/models"
)

func CreateAthlete(model *models.ModelImp, request httpReq.AthleteReq) (httpReq.AthleteRes, error) {
	var athlete domain.Athlete
	copier.Copy(&athlete, request)
	insertedAthlete, err := model.InsertAtlhete(athlete)
	if err != nil {
		return httpReq.AthleteRes{}, err
	}
	var response httpReq.AthleteRes
	copier.Copy(&response, insertedAthlete)

	return response, nil
}
