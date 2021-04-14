package models

import (
	"github.com/joao3101/daniel-joao-project/app/athletes/repository"
	"github.com/joao3101/daniel-joao-project/sqlboiler"
	"github.com/volatiletech/sqlboiler/v4/boil"
)


func (model ModelImp) GetAtlhete(athlete_id int) (*athlete_repository.AthleteModelData, error){
	p , _:= sqlboiler.FindPlayer(model.DB, athlete_id)
	a := &athlete_repository.AthleteModelData{
		ID: athlete_id,
		Name: p.Name,
	}

	return a, nil
}


func (model ModelImp) InsertAtlhete(athlete athlete_repository.AthleteModelData) (*athlete_repository.AthleteModelData, error){
	p := &sqlboiler.Player{Name: athlete.Name}
	
	err := p.Insert(model.DB, boil.Infer())

	if err != nil{
		panic(err)
	}

	return &athlete, err
}