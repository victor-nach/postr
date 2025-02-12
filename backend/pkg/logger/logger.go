package logger

import (
	"go.uber.org/zap"

	"github.com/victor-nach/postr-backend/internal/config"
)

// NewLogger creates a new zap.Logger instance
func NewLogger(appEnv string) (*zap.Logger, error) {
	var cfg zap.Config
	var err error

	cfg = zap.NewProductionConfig()

	if appEnv == config.DevEnv {
		cfg = zap.NewDevelopmentConfig()
	}

	cfg.DisableStacktrace = true

	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	logger = logger.With(
		zap.String("service", "postr-backend"),
		 zap.String("version", "1.0.0"),
		 zap.String("app_env", appEnv),
		)

	return logger, nil
}
