package config

import (
	"database/sql"
	"sync"

	"go.uber.org/zap"

	// Gorm
	"gorm.io/gorm"
)

type Config struct {
	Env                   string
	Address               string
	AnalyticsURL          string
	BackendURL            string
	WebappURL             string
	MailInboundDomain     string
	CoreDB                *sql.DB
	CoreDBGorm            *gorm.DB
	NodeDBString          string
	Logger                *zap.Logger
	WaitGroup             *sync.WaitGroup
	RequestLimitPerSecond float64
	CreditsOnCreatedOrg   int
}
