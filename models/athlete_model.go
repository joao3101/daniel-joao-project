package models

import (
	"github.com/jinzhu/copier"
	atlheteDomain "github.com/joao3101/daniel-joao-project/app/domain"
	"github.com/joao3101/daniel-joao-project/sqlboiler"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (model ModelImp) GetAtlhete(athlete_id int) (atlheteDomain.Athlete, error) {
	p, _ := sqlboiler.FindPlayer(model.DB, athlete_id)

	var athleteRsp atlheteDomain.Athlete
	copier.Copy(&athleteRsp, p)

	return athleteRsp, nil
}

func (model ModelImp) InsertAtlhete(athlete atlheteDomain.Athlete) (*atlheteDomain.Athlete, error) {
	var athleteBank sqlboiler.Player
	copier.Copy(&athleteBank, athlete)

	err := athleteBank.Insert(model.DB, boil.Infer())

	if err != nil {
		panic(err)
	}

	var athleteRsp atlheteDomain.Athlete
	copier.Copy(&athleteRsp, athleteBank)

	return &athlete, err
}
