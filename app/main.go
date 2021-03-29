package main

import (
	"log"
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joao3101/daniel-joao-project/api/config"
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

	dbConfig, err := config.InitDB(coreDBString, "core")
	if err != nil {
		logger.Panic(err.Error())
	}

	coreDB := dbConfig.DB

	// Sometimes the connection is "lazy-loaded"
	err = coreDB.Ping()
	if err != nil {
		logger.Panic(err.Error())
	}

	gormDB := dbConfig.GormDB

	dbcond, err := gormDB.DB()
	if err != nil {
		logger.Panic(err.Error())
	}

	dbcond.SetMaxIdleConns(10)
	dbcond.SetMaxOpenConns(100)
	dbcond.SetConnMaxLifetime(time.Minute * 2)

	e := echo.New()
	log.Fatal(e.Start(viper.GetString("server.address")))
}
