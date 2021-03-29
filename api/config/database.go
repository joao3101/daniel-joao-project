package config

import (
	"database/sql"

	// Gorm
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type DatabaseConfig struct {
	DB     *sql.DB
	GormDB *gorm.DB
}

func InitDB(dataSourceName string, schema string) (*DatabaseConfig, error) {
	coreDB, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	var newLogger gormLogger.Interface
	logger := GetLogger()
	logger = logger.Named(schema)

	configGorm := gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 newLogger,
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: coreDB}), &configGorm)
	if err != nil {
		return nil, err
	}

	RegisterCallbacks(gormDB)

	return &DatabaseConfig{
		DB:     coreDB,
		GormDB: gormDB,
	}, nil
}
