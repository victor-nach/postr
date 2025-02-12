package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/victor-nach/postr-backend/internal/domain"
)

type UserHandler struct {
	service domain.UserService
	logger  *zap.Logger
}

func NewUserHandler(service domain.UserService, logger *zap.Logger) *UserHandler {
	logger = logger.With(zap.String("package", "handlers"))

	return &UserHandler{
		service: service,
		logger:  logger,
	}
}

func (h *UserHandler) ListUsers(c *gin.Context) {
	logr := h.logger.With(zap.String("method", "ListUsers"))

	req, err := h.validateListUsers(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	paginatedUsers, err := h.service.List(c.Request.Context(), req.PageNumber, req.PageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	logr.Info("Users listed successfully", zap.Any("paginated", paginatedUsers))

	resp := APIResponse{
		Status:     successStatus,
		Message:    "Users listed successfully",
		Pagination: &paginatedUsers.Pagination,
		Data:       paginatedUsers.Users,
	}
	c.JSON(http.StatusOK, resp)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	logr := h.logger.With(zap.String("method", "GetUserByID"))

	id, err := h.validateGetUserByID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	user, err := h.service.Get(c.Request.Context(), id)

	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, err)
			return
		}

		c.JSON(http.StatusInternalServerError, err)
		return
	}

	logr.Info("User retrieved successfully", zap.Any("user", user))

	resp := APIResponse{
		Status:  successStatus,
		Message: "User retrieved successfully",
		Data:    user,
	}
	c.JSON(http.StatusOK, resp)
}

func (h *UserHandler) CountUsers(c *gin.Context) {
	logr := h.logger.With(zap.String("method", "CountUsers"))

	count, err := h.service.Count(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	logr.Info("Users count retrieved successfully", zap.Int("count", count))

	resp := APIResponse{
		Status:  successStatus,
		Message: "Users count retrieved successfully",
		Data: Count{
			Count: count,
		},
	}
	c.JSON(http.StatusOK, resp)
}
