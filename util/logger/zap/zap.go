package zap

import (
	"errors"

	"gitlab.com/renodesper/gokit-microservices/util/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// CreateLogger ...
func CreateLogger(env string, level string) (logger.Logger, error) {
	logger, err := createConfig(env, level).Build()
	if err != nil {
		return nil, err
	}

	return logger.Sugar(), nil
}

// SyncLogger ...
func SyncLogger(l logger.Logger) error {
	z, ok := l.(*zap.SugaredLogger)
	if !ok {
		return errors.New("Unexpected logger type")
	}

	return z.Sync()
}

func createConfig(env string, level string) *zap.Config {
	var cfg zap.Config

	if env == "development" {
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		cfg.EncoderConfig.EncodeTime = nil
	} else {
		cfg = zap.NewProductionConfig()
	}

	switch level {
	case logger.Debug:
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case logger.Info:
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case logger.Warn:
		cfg.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case logger.Error:
		cfg.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	}

	return &cfg
}
