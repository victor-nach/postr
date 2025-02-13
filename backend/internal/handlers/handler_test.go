package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"

	"github.com/victor-nach/postr-backend/internal/domain"
	"github.com/victor-nach/postr-backend/internal/domain/mocks"
)

func TestPostHandler_CreatePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostService := mocks.NewMockPostService(ctrl)
	logger := zap.NewNop()
	handler := NewPostHandler(mockPostService, logger)

	reqBody := `{"userId": "b63df572-9bd1-4a4f-9f0d-2a8155a81fde", "title": "Test Title", "body": "Test Body"}`
	req, err := http.NewRequest("POST", "/posts", strings.NewReader(reqBody))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	mockPostService.EXPECT().Create(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, post *domain.Post) error {
			require.Equal(t, "b63df572-9bd1-4a4f-9f0d-2a8155a81fde", post.UserID)
			require.Equal(t, "Test Title", post.Title)
			require.Equal(t, "Test Body", post.Body)
			require.NotEmpty(t, post.ID)
			require.False(t, post.CreatedAt.IsZero())
			return nil
		}).Times(1)

	handler.CreatePost(c)

	require.Equal(t, http.StatusOK, w.Code)

	var resp APIResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	require.Equal(t, "success", resp.Status)
	require.Equal(t, "Post created successfully", resp.Message)

	data, ok := resp.Data.(map[string]interface{})
	require.True(t, ok, "expected Data to be a map")
	require.Equal(t, "b63df572-9bd1-4a4f-9f0d-2a8155a81fde", data["userId"])
	require.Equal(t, "Test Title", data["title"])
	require.Equal(t, "Test Body", data["body"])
}

func TestPostHandler_ListPostsByUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostService := mocks.NewMockPostService(ctrl)
	logger := zap.NewNop()
	handler := NewPostHandler(mockPostService, logger)

	userId := uuid.NewString()
	req, err := http.NewRequest("GET", fmt.Sprintf("/posts?userId=%s", userId), nil)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = gin.Params{gin.Param{Key: "userId", Value: userId}}

	expectedPosts := []domain.Post{
		{
			ID:        "post1",
			UserID:    "12345",
			Title:     "Title 1",
			Body:      "Body 1",
			CreatedAt: time.Now(),
		},
		{
			ID:        "post2",
			UserID:    "12345",
			Title:     "Title 2",
			Body:      "Body 2",
			CreatedAt: time.Now(),
		},
	}

	mockPostService.EXPECT().List(gomock.Any(), userId).Return(expectedPosts, nil).Times(1)

	handler.ListPostsByUserID(c)

	require.Equal(t, http.StatusOK, w.Code)

	var resp APIResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	require.Equal(t, "success", resp.Status)
	require.Equal(t, "Posts listed successfully", resp.Message)

	dataSlice, ok := resp.Data.([]interface{})
	require.True(t, ok, "expected Data to be a slice")
	require.Len(t, dataSlice, len(expectedPosts))
}

func TestPostHandler_DeletePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostService := mocks.NewMockPostService(ctrl)
	logger := zap.NewNop()
	handler := NewPostHandler(mockPostService, logger)

	postID := uuid.NewString()

	req, err := http.NewRequest("DELETE", "/posts/"+postID, nil)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	c.Params = gin.Params{
		{Key: "id", Value: postID},
	}

	mockPostService.EXPECT().Delete(gomock.Any(), postID).Return(nil).Times(1)

	handler.DeletePost(c)

	require.Equal(t, http.StatusNoContent, w.Code)
	require.Empty(t, w.Body.Bytes())
}

func TestUserHandler_ListUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mocks.NewMockUserService(ctrl)
	logger := zap.NewNop()
	handler := NewUserHandler(mockUserService, logger)

	req, err := http.NewRequest("GET", "/users?pageNumber=1&pageSize=10", nil)
	require.NoError(t, err)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	paginatedUsers := domain.PaginatedUsers{
		Pagination: domain.Pagination{
			CurrentPage: 1,
			TotalSize:   1,
			TotalPages:  1,
		},
		Users: []domain.User{
			{ID: uuid.NewString()},
		},
	}
	mockUserService.EXPECT().List(gomock.Any(), 1, 10).Return(paginatedUsers, nil).Times(1)

	handler.ListUsers(c)
	require.Equal(t, http.StatusOK, w.Code)

	var resp APIResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	require.Equal(t, "success", resp.Status)
	require.Equal(t, "Users listed successfully", resp.Message)
	require.NotNil(t, resp.Pagination)
	users, ok := resp.Data.([]interface{})
	require.True(t, ok)
	require.Len(t, users, 1)
}

func TestUserHandler_GetUserByID_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mocks.NewMockUserService(ctrl)
	logger := zap.NewNop()
	handler := NewUserHandler(mockUserService, logger)

	userID := uuid.NewString()
	req, err := http.NewRequest("GET", "/users/"+userID, nil)
	require.NoError(t, err)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = gin.Params{{Key: "id", Value: userID}}

	user := &domain.User{ID: userID}
	mockUserService.EXPECT().Get(gomock.Any(), userID).Return(user, nil).Times(1)

	handler.GetUserByID(c)
	require.Equal(t, http.StatusOK, w.Code)

	var resp APIResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	require.Equal(t, "success", resp.Status)
	require.Equal(t, "User retrieved successfully", resp.Message)
	data, ok := resp.Data.(map[string]interface{})
	require.True(t, ok)
	require.Equal(t, userID, data["id"])
}

func TestUserHandler_CountUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mocks.NewMockUserService(ctrl)
	logger := zap.NewNop()
	handler := NewUserHandler(mockUserService, logger)

	req, err := http.NewRequest("GET", "/users/count", nil)
	require.NoError(t, err)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	mockUserService.EXPECT().Count(gomock.Any()).Return(42, nil).Times(1)

	handler.CountUsers(c)
	require.Equal(t, http.StatusOK, w.Code)

	var resp APIResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	require.Equal(t, "success", resp.Status)
	require.Equal(t, "Users count retrieved successfully", resp.Message)
	
	var countData Count
	dataBytes, err := json.Marshal(resp.Data)
	require.NoError(t, err)
	err = json.Unmarshal(dataBytes, &countData)
	require.NoError(t, err)
	require.Equal(t, 42, countData.Count)
}