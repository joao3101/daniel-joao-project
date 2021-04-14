package main

import (
	"log"
	_ "log"

	"github.com/joao3101/daniel-joao-project/api"
	"github.com/joao3101/daniel-joao-project/util"
	_ "github.com/labstack/echo"
)

func init() {
}

func main() {
	//logger := config.GetLogger()

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	server, err := api.NewServer(config)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

	/*model, err := models.InitModel(config.DBDriver, config.DBSource)
	if err != nil {
		logger.Panic(err.Error())
	}

	a := athlete_repository.AthleteModelData{Name: "afjkdlsadh"}
	model.InsertAtlhete(a)*/

	//e := echo.New()
	//log.Fatal(e.Start(viper.GetString("server.address")))
}
