package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

const (
	// Environment variable keys
	EnvPort         = "PORT"
	EnvAppEnv       = "APP_ENV"
	EnvApiKeys      = "API_KEYS"
	EnvRateLimitKey = "RATE_LIMIT_RPS"

	// Default values
	DefaultPort         = "8080"
	DefaultAppEnv       = "production"
	DefaultRateLimitEnv = "5"

	// APP envs
	ProdEnv = "production"
	DevEnv = "development"
)

// Config holds the application configuration
type Config struct {
	Port       string
	AppEnv     string
	APIKeys    map[string]string
	RateLimtPS int
}

// Load reads configuration from the environment and loads the .env file in the project root if available
// Sets default values if applicable
func Load(logger *zap.Logger) (*Config, error) {
	// Load environment variables from .env if available.
	err := godotenv.Load()
	if err != nil {
		logger.Warn("error loading .env file, using default values")
	}

	port, ok := os.LookupEnv(EnvPort)
	if !ok {
		port = DefaultPort
	}

	appEnv, ok := os.LookupEnv(EnvAppEnv)
	if !ok {
		appEnv = DefaultAppEnv
	}

	rpsStr, ok := os.LookupEnv(EnvRateLimitKey)
	if !ok {
		rpsStr = DefaultRateLimitEnv
	}
	rps, err := strconv.Atoi(rpsStr)
	if err != nil {
		rpsStr = DefaultRateLimitEnv
	}

	apiKeysStr, ok := os.LookupEnv(EnvApiKeys)
	if !ok {
		logger.Warn("no api keys loaded")
	}
	apiKeys, err := parseAPIKeys(apiKeysStr)
	if err != nil {
		logger.Warn("error retrieving api keys", zap.Error(err))
	}

	cfg := &Config{
		Port:       port,
		AppEnv:     appEnv,
		APIKeys:    apiKeys,
		RateLimtPS: rps,
	}

	logger.Info("Configuration loaded",
		zap.String("port", cfg.Port),
		zap.String("app_env", cfg.AppEnv),
		zap.Int("api_key_count", len(cfg.APIKeys)),
	)

	return cfg, nil
}

func parseAPIKeys(keysStr string) (map[string]string, error) {
	var apiKeys map[string]string

	if err := json.Unmarshal([]byte(keysStr), &apiKeys); err != nil {
		return nil, fmt.Errorf("error unmarshalling json: %w", err)
	}

	return apiKeys, nil
}
