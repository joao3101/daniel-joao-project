package main

import (
	"log"

	"github.com/joao3101/daniel-joao-project/api/routes"
	"github.com/joao3101/daniel-joao-project/util"
	_ "github.com/lib/pq"
)

func init() {
}

func main() {
	//logger := config.GetLogger()

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	/*server, err := api.NewServer(config)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	model, err := models.InitModel(config.DBDriver, config.DBSource)
	if err != nil {
		logger.Panic(err.Error())
	}

	a := domainAthlete.Athlete{
		Name:     util.RandomString(8),
		Age:      int(util.RandomInt(18, 33)),
		Position: int(util.RandomInt(1, 11)),
	}
	model.InsertAtlhete(a)

	fmt.Println(model.GetAtlhete(1))*/

	routes.Bootstrap(&config)

	/*err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}*/

	//e := echo.New()
	//log.Fatal(e.Start(viper.GetString("server.address")))
}
