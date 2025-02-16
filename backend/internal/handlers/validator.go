package handlers

import (
	"strconv"
	"strings"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/go-ozzo/ozzo-validation/v4"
	"go.uber.org/zap"

	"github.com/microcosm-cc/bluemonday"

	"github.com/victor-nach/postr-backend/internal/domain"
)

func sanitizeInput(input string) string {
	p := bluemonday.StrictPolicy()
	return p.Sanitize(input)
}

var compactUUIDRegex = regexp.MustCompile(`^[0-9a-fA-F]{32}$`)

func isCompactUUID(value interface{}) error {
	s, ok := value.(string)
	if !ok {
		return validation.NewError("validation_compact_uuid", "invalid UUID format")
	}
	if !compactUUIDRegex.MatchString(s) {
		return validation.NewError("validation_compact_uuid", "invalid compact UUID format")
	}
	return nil
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

	err := validation.Validate(userId, validation.By(isCompactUUID))
    if err != nil {
        logr.Error("invalid userId format", zap.Error(err))
        return "", domain.ErrInvalidInputWithStr("invalid userId format")
    }

	return userId, nil
}

func (h *PostHandler) validateDeletePost(c *gin.Context) (string, error) {
	logr := h.logger.With(zap.String("method", "validateDeletePost"))

	id := c.Param("id")
	
	err := validation.Validate(id, validation.By(isCompactUUID))
    if err != nil {
        logr.Error("invalid userId format", zap.Error(err))
        return "", domain.ErrInvalidInputWithStr("invalid userId format")
    }

	return id, nil
}

func (h *UserHandler) validateListUsers(c *gin.Context) (*listUsersRequest, error) {
	pageNumber := 1
	pageSize := 10

	if pn := c.Query("pageNumber"); pn != "" {
		if num, err := strconv.Atoi(pn); err == nil {
			pageNumber = num
		} else {
			return nil, domain.ErrInvalidInputWithStr("nvalid pageNumber value")
		}
	}

	if ps := c.Query("pageSize"); ps != "" {
		if num, err := strconv.Atoi(ps); err == nil {
			pageSize = num
		} else {
			return nil, domain.ErrInvalidInputWithStr("nvalid pageSize value")
		}
	}

	req := listUsersRequest{
		PageNumber: pageNumber,
		PageSize:   pageSize,
	}

	if err := req.Validate(); err != nil {
		if verrs, ok := err.(validation.Errors); ok {
			return nil, domain.ErrInvalidInput.WithFieldErrors(verrs)
		}
		return nil, domain.ErrInvalidInputWithStr("Pagination validation failednvalid pageSize value")
	}

	return &req, nil
}

func (h *UserHandler) validateGetUserByID(c *gin.Context) (string, error) {
	logr := h.logger.With(zap.String("method", "GetUserByID"))

	id := c.Param("id")
	
	err := validation.Validate(id, validation.By(isCompactUUID))
    if err != nil {
        logr.Error("invalid userId format", zap.Error(err))
        return "", domain.ErrInvalidInputWithStr("invalid userId format")
    }

	return id, nil
}
