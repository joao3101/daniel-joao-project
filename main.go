package main

import (
	"log"
	_ "log"

	"github.com/joao3101/daniel-joao-project/api/config"
	athlete_repository "github.com/joao3101/daniel-joao-project/app/athletes/repository"
	"github.com/joao3101/daniel-joao-project/models"
	"github.com/joao3101/daniel-joao-project/util"
	_ "github.com/labstack/echo"
	"github.com/spf13/viper"
)

func init() {
}

func main() {
	_ = viper.GetString(`database.core.wr`)
	logger := config.GetLogger()

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	model, err := models.InitModel(config.DBDriver, config.DBSource)
	if err != nil {
		logger.Panic(err.Error())
	}

	a := athlete_repository.AthleteModelData{Name: "afjkdlsadh"}
	model.InsertAtlhete(a)

	//e := echo.New()
	//log.Fatal(e.Start(viper.GetString("server.address")))
}
