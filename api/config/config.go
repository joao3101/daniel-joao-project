package config

import (
	"database/sql"
	"sync"

	"go.uber.org/zap"
)

type Config struct {
	Env                   string
	Address               string
	AnalyticsURL          string
	BackendURL            string
	WebappURL             string
	MailInboundDomain     string
	CoreDB                *sql.DB
	NodeDBString          string
	Logger                *zap.Logger
	WaitGroup             *sync.WaitGroup
	RequestLimitPerSecond float64
	CreditsOnCreatedOrg   int
}
