package handlers

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"go.uber.org/zap"

	"github.com/microcosm-cc/bluemonday"

	"github.com/victor-nach/postr-backend/internal/domain"
)

func sanitizeInput(input string) string {
	p := bluemonday.StrictPolicy()
	return p.Sanitize(input)
}

func (r createPostRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.UserID, validation.Required, is.UUID),
		validation.Field(&r.Title, validation.Required),
		validation.Field(&r.Body, validation.Required),
	)
}

func (h *PostHandler) validateCreatePost(c *gin.Context) (*createPostRequest, error) {
	var req createPostRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Error binding JSON", zap.Error(err))
		return nil, domain.ErrInvalidInput
	}

	req.UserID = sanitizeInput(strings.TrimSpace(req.UserID))
	req.Title = sanitizeInput(strings.TrimSpace(req.Title))
	req.Body = sanitizeInput(strings.TrimSpace(req.Body))

	if err := req.Validate(); err != nil {
		if verrs, ok := err.(validation.Errors); ok {
			h.logger.Error("Validation errors", zap.Any("errors", verrs))
			return nil, domain.ErrInvalidInput.WithFieldErrors(verrs)
		}

		h.logger.Error("Validation error", zap.Error(err))
		return nil, domain.ErrInvalidInput
	}

	return &req, nil
}
