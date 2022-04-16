package helper

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Var
var (
	logger *zap.SugaredLogger
)

// InitLogger : Init logger with debug option
func InitLogger(isDebug bool) *zap.SugaredLogger {
	loggerMgr := initZapLog(isDebug)
	zap.ReplaceGlobals(loggerMgr)
	defer func() {
		// flushes buffer, if any
		_ = zap.S().Sync()
	}()
	return zap.S()
}

// GetLogger - call InitLogger first, otherwise, return a default logger
func GetLogger() *zap.SugaredLogger {
	if logger == nil {
		logger = zap.S()
	}
	return logger
}

// Init Zap Logger
func initZapLog(isDebug bool) *zap.Logger {
	config := zap.NewDevelopmentConfig()
	if isDebug {
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	} else {
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
		config.EncoderConfig.LevelKey = ""
		config.EncoderConfig.TimeKey = ""
		config.EncoderConfig.CallerKey = ""
	}
	logger, _ := config.Build()
	return logger
}

// HandleError : Handle an error without exiting (just logging)
func HandleError(e error) {
	if e != nil {
		GetLogger().Error(e)
	}
}

// HandleErrorExit : Handle an error and exit
func HandleErrorExit(e error) {
	if e != nil {
		GetLogger().Info(e)
		os.Exit(1)
	}
}
