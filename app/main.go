package main

import (
	"log"

	"github.com/joao3101/daniel-joao-project/api/config"
	"github.com/joao3101/daniel-joao-project/models"
	"github.com/joao3101/daniel-joao-project/app/repository"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`../config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	coreDBString := viper.GetString(`database.core.wr`)
	logger := config.GetLogger()

	db, err := models.InitModel(coreDBString, "core")
	if err != nil {
		logger.Panic(err.Error())
	}

	model := &models.ModelImp{}
	model.InsertAtlhete()

	e := echo.New()
	log.Fatal(e.Start(viper.GetString("server.address")))
}
