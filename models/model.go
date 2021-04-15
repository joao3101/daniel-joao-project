package models

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type ModelImp struct {
	DB *sql.DB
}

func InitModel(dbDriver, dataSourceName string) (*ModelImp, error) {
	coreDB, err := sql.Open(dbDriver, dataSourceName)
	if err != nil {
		return nil, err
	}

	boil.DebugMode = true

	return &ModelImp{
		DB: coreDB,
	}, nil
}
