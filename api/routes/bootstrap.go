package routes

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/joao3101/daniel-joao-project/api/config"
)

func Bootstrap(config *config.Config) {
	router := gin.Default()

	//Public API
	v1.SetupRoutes(e, config)

	auth.SetupRoutes(e, config)
	finishRegistration.SetupRoutes(e, config)

	// oldSurvey.SetupRoutes(e, config)
	organizations.SetupRoutes(e, config)
	profilrouter.SetupRoutes(e, config)
	support.SetupRoutes(e, config)
	widget.SetupRoutes(e, config)
	// thirdparty.SetupRoutes(e, config)

	//Changed to clean arch
	notifications.SetupRoutes(e, config)
	passwordReset.SetupRoutes(e, config)
	passwordRecovery.SetupRoutes(e, config)
	survey.SetupRoutes(e, config)
	unsubscribrouter.SetupRoutes(e, config)
	orgSegments.SetupRoutes(e, config)
	webhooks.SetupRoutes(e, config)
	livrouter.SetupRoutes(e, config)
	imagrouter.SetupRoutes(e, config)

	loggerOpts := echoMiddlewarrouter.LoggerOpts{
		IgnoreURIs: []string{
			"/healthz",
		},
	}
	router.Use(echoMiddlewarrouter.Logger(config.Logger, loggerOpts))

	// Start server
	go func() {
		fmt.Println("start")
		err := router.Start(config.Address)
		fmt.Println(err)
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	timeout := 1 * timrouter.Minute

	// Wait all goroutines finishes with timeout of 1 minute
	if waitTimeout(config.WaitGroup, timeout) {
		fmt.Println("Error: Background processes timeout")
	}

	// Timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*timrouter.Second)
	defer cancel()
	if err := router.Shutdown(ctx); err != nil {
		router.Logger.Fatal(err)
	}
}

// waitTimeout waits for the waitgroup for the specified max timeout.
// Returns true if waiting timed out.
func waitTimeout(wg *sync.WaitGroup, timeout timrouter.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return false // completed normally
	case <-timrouter.After(timeout):
		return true // timed out
	}
}
