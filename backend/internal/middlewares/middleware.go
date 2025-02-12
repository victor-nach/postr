package middlewares

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/time/rate"

	"github.com/victor-nach/postr-backend/internal/config"
	"github.com/victor-nach/postr-backend/internal/domain"
)

const (
	UserIDKey = "user_id"
)

type Service struct {
	logger       *zap.Logger
	config       *config.Config
	userLimiters map[string]*rate.Limiter
	mu           sync.Mutex
}

func New(logger *zap.Logger, cfg *config.Config) *Service {
	return &Service{
		logger:       logger,
		config:       cfg,
		userLimiters: make(map[string]*rate.Limiter),
	}
}

// AuthMiddleware checks for the X-API-Key header against valid keys
func (m *Service) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// skip checks for dev environment
		if m.config.AppEnv == config.DevEnv {
			m.logger.Info("skipping API key check in development mode")
			c.Set(UserIDKey, "dev-user")
			c.Next()
			return
		}

		apiKey := c.GetHeader("X-API-Key")
		if apiKey == "" {
			m.logger.Error("missing API key")
			c.JSON(http.StatusUnauthorized, domain.ErrMissingAPIKey)
			c.Abort()
			return
		}

		// Validate the provided key
		var userID string
		authorized := false
		for id, keyValue := range m.config.APIKeys {
			if keyValue == apiKey {
				userID = id
				authorized = true
				break
			}
		}

		if !authorized {
			m.logger.Error("invalid API key", zap.String("api_key", apiKey))
			c.JSON(http.StatusUnauthorized, domain.ErrInvalidAPIKey)
			c.Abort()
			return
		}

		c.Set(UserIDKey, userID)
		m.logger.Info("user authenticated", zap.String("user_id", userID))

		c.Next()
	}
}

// RateLimitMiddleware applies a per-user rate limit based on the userID from the context
func (m *Service) RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the userID from the context
		userID, exists := c.Get(UserIDKey)
		if !exists {
			m.logger.Warn("user_id not found in context for rate limiting")
			c.Next()
			return
		}

		limiter := m.getLimiter(userID.(string))
		if !limiter.Allow() {
			// Rate limit exceeded
			m.logger.Warn("rate limit exceeded", zap.String("user_id", userID.(string)))
			c.JSON(http.StatusTooManyRequests, domain.ErrTooManyRequests)
			c.Abort()
			return
		}

		c.Next()
	}
}

// getLimiter retrieves or creates a rate.Limiter for the given user
func (m *Service) getLimiter(userID string) *rate.Limiter {
	m.mu.Lock()
	defer m.mu.Unlock()

	limiter, exists := m.userLimiters[userID]
	if !exists {
		// Create a new limiter if none exists for this user
		r := rate.Limit(m.config.RateLimtPS)
		limiter = rate.NewLimiter(r, m.config.RateLimtPS)
		m.userLimiters[userID] = limiter
	}

	return limiter
}
