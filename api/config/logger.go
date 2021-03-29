package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var appLogger *zap.Logger

func initLogger() *zap.Logger {
	env := viper.GetString(`env`)

	logger := initZap(env)
	logger = logger.Named(env)

	defer logger.Sync()

	return logger
}

func initZap(env string) *zap.Logger {
	consoleCore := loadConsoleCore(env)

	return zap.New(consoleCore)
}

func loadConsoleCore(env string) zapcore.Core {
	if env == "prod" {
		return prodConsoleCore()
	} else {
		return devConsoleCore()
	}
}

func devConsoleCore() zapcore.Core {
	level := zap.NewAtomicLevelAt(zapcore.DebugLevel)
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeLevel = customLevelEncoder
	encoderConfig.EncodeName = customEncodeName
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	consoleDebugging := zapcore.Lock(os.Stdout)
	return zapcore.NewCore(encoder, consoleDebugging, level)
}

func prodConsoleCore() zapcore.Core {
	level := zap.NewAtomicLevelAt(zapcore.InfoLevel)
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeLevel = customLevelEncoder
	encoderConfig.EncodeName = customEncodeName
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	consoleDebugging := zapcore.Lock(os.Stdout)
	return zapcore.NewCore(encoder, consoleDebugging, level)
}

func customLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

func customEncodeName(loggerName string, enc zapcore.PrimitiveArrayEncoder) {
	data := strings.Split(loggerName, ".")
	if len(data) > 0 {
		data[0] = "env:" + data[0]
	}

	if len(data) > 1 {
		data[1] = "schema:" + data[1]
	}

	enc.AppendString(strings.Join(data, " "))
}

func GetLogger() *zap.Logger {
	if appLogger != nil {
		return appLogger
	}

	appLogger = initLogger()

	return appLogger
}
