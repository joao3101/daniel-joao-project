package routes

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joao3101/daniel-joao-project/api/config"
	"github.com/joao3101/daniel-joao-project/api/routes/athletes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

func Bootstrap(config *config.Config) {
	e := echo.New()

	//e.Use(echoMiddleware.RateLimit(config, rateLimit))

	e.Debug = true
	e.HideBanner = true

	// Limiting the number of requests per second
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(config.RequestLimitPerSecond))))

	e.Use(middleware.BodyLimit("50M"))

	e.Use(middleware.Recover())

	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		Skipper:               middleware.DefaultSkipper,
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "SAMEORIGIN",
		HSTSMaxAge:            15768000,
		ContentSecurityPolicy: "default-src 'self' *.track.co *.sendgrid.net *.googleapis.com *.gstatic.com ;  img-src 'self' *.track.co *.sendgrid.net data: ; script-src 'self'   *.sendgrid.net cdn.lr-ingest.io 'unsafe-inline' 'unsafe-eval' *.googleapis.com; style-src 'self' *.googleapis.com 'unsafe-inline';",
		ReferrerPolicy:        "strict-origin-when-cross-origin",
		//HSTSPreloadEnabled:    true,
	}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	//auth.SetupRoutes(e, config)

	athletes.SetupRoutes(e, config)

	// Start server
	go func() {
		fmt.Println("start")
		err := e.Start(config.Address)
		fmt.Println(err)
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
}
