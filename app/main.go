package main

import (
	_ "log"

	"github.com/joao3101/daniel-joao-project/api/config"
	athlete_repository "github.com/joao3101/daniel-joao-project/app/athletes/repository"
	"github.com/joao3101/daniel-joao-project/models"
	_ "github.com/labstack/echo"
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
	_ = viper.GetString(`database.core.wr`)
	logger := config.GetLogger()

	model, err := models.InitModel("root@/fantasy?parseTime=true&charset=utf8&group_concat_max_len=65535", "core")
	if err != nil {
		logger.Panic(err.Error())
	}

	a := athlete_repository.AthleteModelData{Name: "afjkdlsadh"}
	model.InsertAtlhete(a)

	//e := echo.New()
	//log.Fatal(e.Start(viper.GetString("server.address")))
}
