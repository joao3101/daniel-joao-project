package middleware

import (
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerOpts struct {
	IgnoreURIs []string
}

//Logger ...
func Logger(logger *zap.Logger, opts ...LoggerOpts) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)
			if err != nil {
				c.Error(err)
			}

			req := c.Request()
			res := c.Response()

			fields := []zapcore.Field{
				zap.Int("status", res.Status),
				zap.String("latency", time.Since(start).String()),
				zap.String("method", req.Method),
				zap.String("uri", req.RequestURI),
				zap.String("host", req.Host),
				zap.String("remote_ip", c.RealIP()),
				zap.String("user_agent", req.UserAgent()),
			}

			if len(opts) > 0 {
				loggerOpts := opts[0]
				for _, uri := range loggerOpts.IgnoreURIs {
					if req.RequestURI == uri {
						return nil
					}
				}
			}

			namedLogger := logger
			routePath := c.Path()

			if strings.Contains(routePath, "/organizations/:orgUUID") {
				schema := "unknown"
				orgUUID := c.Param("orgUUID")

				user := c.Get("user")
				if user != nil {
					userToken, ok := c.Get("user").(*jwt.Token)
					if ok {
						claims := userToken.Claims.(jwt.MapClaims)
						orgList := claims["Data"].(map[string]interface{})
						orgConfig := orgList[orgUUID]
						orgConfigStr := orgConfig.(string)
						orgConfigData := strings.Split(orgConfigStr, ",")
						if len(orgConfigData) >= 2 {
							schema = orgConfigData[1]
						}
					}
				}

				namedLogger = logger.Named(schema)
			} else {
				namedLogger = logger.Named("core")
			}

			n := res.Status
			switch {
			case n >= 500:
				namedLogger.Error("Server error", fields...)
			case n >= 400:
				namedLogger.Warn("Client error", fields...)
			case n >= 300:
				namedLogger.Info("Redirection", fields...)
			default:
				namedLogger.Info("Success", fields...)
			}

			return nil
		}
	}
}
