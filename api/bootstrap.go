package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/joao3101/daniel-joao-project/api/config"
	echoMiddleware "github.com/joao3101/daniel-joao-project/pkg/echo/middleware"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs"
	"golang.org/x/time/rate"
)

func Bootstrap(config *config.Config) {
	e := echo.New()

	//e.Use(echoMiddleware.RateLimit(config, rateLimit))

	e.Debug = false
	e.HideBanner = true

	// Limiting the number of requests per second
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(config.RequestLimitPerSecond))))

	e.Use(middleware.BodyLimit("50M"))

	e.Use(middleware.Recover())

	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		Skipper:            middleware.DefaultSkipper,
		XSSProtection:      "1; mode=block",
		ContentTypeNosniff: "nosniff",
		XFrameOptions:      "SAMEORIGIN",
		HSTSMaxAge:         15768000,
		ReferrerPolicy:     "strict-origin-when-cross-origin",
		//HSTSPreloadEnabled:    true,
	}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	if config.Env == "dev" {
		e.Debug = true
		//e.Use(echoMiddleware.Chaos(10))
		//e.Use(echoMiddleware.Sloth(50, 100))
	}

	e.GET("/healthz", func(c echo.Context) error {
		return c.HTML(200, "Healthy!")
	})

	/*auth.SetupRoutes(e, config)
	finishRegistration.SetupRoutes(e, config)

	// oldSurvey.SetupRoutes(e, config)
	organizations.SetupRoutes(e, config)
	profile.SetupRoutes(e, config)
	support.SetupRoutes(e, config)
	widget.SetupRoutes(e, config)
	// thirdparty.SetupRoutes(e, config)

	//Changed to clean arch
	notifications.SetupRoutes(e, config)
	passwordReset.SetupRoutes(e, config)
	passwordRecovery.SetupRoutes(e, config)
	survey.SetupRoutes(e, config)
	unsubscribe.SetupRoutes(e, config)
	orgSegments.SetupRoutes(e, config)
	webhooks.SetupRoutes(e, config)
	live.SetupRoutes(e, config)
	image.SetupRoutes(e, config)
	*/

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	loggerOpts := echoMiddleware.LoggerOpts{
		IgnoreURIs: []string{
			"/healthz",
		},
	}
	e.Use(echoMiddleware.Logger(config.Logger, loggerOpts))

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

	timeout := 1 * time.Minute

	// Wait all goroutines finishes with timeout of 1 minute
	if waitTimeout(config.WaitGroup, timeout) {
		fmt.Println("Error: Background processes timeout")
	}

	// Timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

// waitTimeout waits for the waitgroup for the specified max timeout.
// Returns true if waiting timed out.
func waitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return false // completed normally
	case <-time.After(timeout):
		return true // timed out
	}
}
