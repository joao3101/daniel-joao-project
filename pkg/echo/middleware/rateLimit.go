package middleware

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/joao3101/daniel-joao-project/api/config"
	"github.com/labstack/echo/v4"
)

const secondsToLive = 60

// RateLimit is the middleware function.
func RateLimit(config *config.Config, limit int) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			exists, err := config.Cache.Exists(c.RealIP())
			if err != nil {
				return errors.New("CANNOT connect to redis")
			}

			accessCount := 1
			createdTime := time.Now()
			if exists {
				accessCountString, err := config.Cache.Get(c.RealIP())
				if err != nil {
					return errors.New("CANNOT connect to redis")
				}

				ttl, err := config.Cache.TimeToLive(c.RealIP())
				if err != nil {
					return errors.New("CANNOT connect to redis")
				}

				accessCount, _ = strconv.Atoi(accessCountString)
				accessCount++

				//check if key has TTL
				if ttl == -1 {
					err = config.Cache.Del(c.RealIP())
					if err != nil {
						return errors.New("CANNOT connect to redis")
					}
					return errors.New("Restarting IP " + c.RealIP() + "on Redis")
				}

				if accessCount > limit {
					return c.JSON(http.StatusTooManyRequests, "Too many requests")
				}

				// Generate when this key was inserted, to set generate X-RateLimit-Reset
				createdTime = createdTime.Add((-secondsToLive + time.Duration(ttl)) * time.Second)

				err = config.Cache.Increase(c.RealIP())
				if err != nil {
					return errors.New("CANNOT connect to redis")
				}
			} else {
				err = config.Cache.SetWithExpiration(c.RealIP(), strconv.Itoa(1), secondsToLive)
				if err != nil {
					return errors.New("CANNOT connect to redis")
				}
			}

			resetTime := createdTime.Add(60 * time.Second).Unix()

			// Add X-Limit Header
			c.Response().Header().Set("X-RateLimit-Limit", strconv.Itoa(limit))
			c.Response().Header().Set("X-RateLimit-Remaining", strconv.Itoa(limit-accessCount))
			c.Response().Header().Set("X-RateLimit-Reset", strconv.FormatInt(resetTime, 10))

			return next(c)
		}
	}
}
