package models

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type ModelImp struct {
	DB     *sql.DB
}


func InitModel(dataSourceName string, schema string) (*ModelImp, error) {
	coreDB, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	err = coreDB.Ping()
	if err != nil {
		panic(err.Error())
	}

	coreDB.SetMaxIdleConns(10)
	coreDB.SetMaxOpenConns(100)
	coreDB.SetConnMaxLifetime(time.Minute * 2)

	return &ModelImp{
		DB:     coreDB,
	}, nil
}