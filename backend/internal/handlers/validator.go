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
	logr := h.logger.With(zap.String("method", "validateCreatePost"))
	var req createPostRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		logr.Error("Error binding JSON", zap.Error(err))
		return nil, domain.ErrInvalidInput
	}

	req.UserID = sanitizeInput(strings.TrimSpace(req.UserID))
	req.Title = sanitizeInput(strings.TrimSpace(req.Title))
	req.Body = sanitizeInput(strings.TrimSpace(req.Body))

	if err := req.Validate(); err != nil {
		if verrs, ok := err.(validation.Errors); ok {
			logr.Error("Validation errors", zap.Any("errors", verrs))
			return nil, domain.ErrInvalidInput.WithFieldErrors(verrs)
		}

		logr.Error("Validation error", zap.Error(err))
		return nil, domain.ErrInvalidInput
	}

	return &req, nil
}

func (h *PostHandler) validateListPostsByUserID(c *gin.Context) (string, error) {
	logr := h.logger.With(zap.String("method", "validateListPostsByUserID"))

	userId := c.Query("userId")
	if userId == "" {
		logr.Error("missing userId query parameter")
		return "", domain.ErrInvalidInputWithStr("missing userId query parameter")
		
	}

	err := validation.Validate(userId, is.UUID)
    if err != nil {
        logr.Error("invalid userId format", zap.Error(err))
        return "", domain.ErrInvalidInputWithStr("invalid userId format")
    }

	return userId, nil
}

func (h *PostHandler) validateDeletePost(c *gin.Context) (string, error) {
	logr := h.logger.With(zap.String("method", "validateDeletePost"))

	id := c.Param("id")
	
	err := validation.Validate(id, is.UUID)
    if err != nil {
        logr.Error("invalid userId format", zap.Error(err))
        return "", domain.ErrInvalidInputWithStr("invalid userId format")
    }

	return id, nil
}
