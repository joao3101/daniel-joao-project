package models

import (
	"github.com/joao3101/daniel-joao-project/sqlboiler"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (model ModelImp) GetAtlhete(athlete_id int) (athlete sqlboiler.Player, err error) {
	p, _ := sqlboiler.FindPlayer(model.DB, athlete_id)

	return *p, nil
}

func (model ModelImp) InsertAtlhete(athlete sqlboiler.Player) (*sqlboiler.Player, error) {
	p := &athlete

	err := p.Insert(model.DB, boil.Infer())

	if err != nil {
		panic(err)
	}

	return &athlete, err
}
