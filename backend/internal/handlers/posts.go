package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/victor-nach/postr-backend/internal/domain"
)

type PostHandler struct {
	service domain.PostService
	logger  *zap.Logger
}

func NewPostHandler(service domain.PostService, logger *zap.Logger) *PostHandler {
	logger = logger.With(zap.String("package", "handlers"))

	return &PostHandler{
		service: service,
		logger:  logger,
	}
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	logr := h.logger.With(zap.String("method", "CreatePost"))

	req, err := h.validateCreatePost(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	post := &domain.Post{
		ID:        uuid.NewString(),
		UserID:    req.UserID,
		Title:     req.Title,
		Body:      req.Body,
		CreatedAt: time.Now(),
	}

	if err := h.service.Create(c.Request.Context(), post); err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, err)
			return
		}

		c.JSON(http.StatusInternalServerError, err)
		return
	}

	logr.Info("Post created successfully", zap.Any("post", post))

	resp := APIResponse{
		Status:  successStatus,
		Message: "Post created successfully",
		Data:    post,
	}
	c.JSON(http.StatusOK, resp)
}

func (h *PostHandler) ListPostsByUserID(c *gin.Context) {
	logr := h.logger.With(zap.String("method", "ListPostsByUserID"))

	userId, err := h.validateListPostsByUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	posts, err := h.service.List(c.Request.Context(), userId)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, err)
			return
		}

		c.JSON(http.StatusInternalServerError, domain.ErrInternalServer)
		return
	}

	logr.Info("Posts listed successfully", zap.String("userId", userId), zap.Int("count", len(posts)))

	resp := APIResponse{
		Status:  successStatus,
		Message: "Posts listed successfully",
		Data:    posts,
	}
	c.JSON(http.StatusOK, resp)
}

func (h *PostHandler) DeletePost(c *gin.Context) {
	logr := h.logger.With(zap.String("method", "DeletePost"))

	id, err := h.validateDeletePost(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		if errors.Is(err, domain.ErrPostNotFound) {
			c.JSON(http.StatusNotFound, err)
			return
		}

		c.JSON(http.StatusInternalServerError, err)
		return
	}

	logr.Info("Post deleted successfully", zap.String("id", id))
	c.JSON(http.StatusNoContent, nil)
}
