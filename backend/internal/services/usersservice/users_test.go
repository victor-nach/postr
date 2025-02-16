package usersservice

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/stretchr/testify/require"

	"github.com/victor-nach/postr-backend/internal/domain"
	"github.com/victor-nach/postr-backend/internal/services/usersservice/mocks"
)

func TestService_Get_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockusersRepo(ctrl)
	logger := zap.NewNop()
	svc := New(mockRepo, logger)

	ctx := context.Background()
	userID := uuid.NewString()
	expectedUser := &domain.User{
		ID:       userID,
		Name:     "Bob",
		Username: "Jones",
		Email:    "bob@example.com",
	}

	mockRepo.EXPECT().Get(ctx, userID).Return(expectedUser, nil)

	user, err := svc.Get(ctx, userID)
	require.NoError(t, err)
	require.Equal(t, expectedUser, user)
}

func TestService_Get_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockusersRepo(ctrl)
	logger := zap.NewNop()
	svc := New(mockRepo, logger)

	ctx := context.Background()
	userID := uuid.NewString()

	mockRepo.EXPECT().Get(ctx, userID).Return(nil, gorm.ErrRecordNotFound)

	user, err := svc.Get(ctx, userID)
	require.Error(t, err)
	require.Nil(t, user)
	require.Equal(t, domain.ErrUserNotFound, err)
}

func TestService_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockusersRepo(ctrl)
	logger := zap.NewNop()
	svc := New(mockRepo, logger)

	ctx := context.Background()
	pageNumber, pageSize := 1, 10

	expectedPaginated := domain.PaginatedUsers{
		Pagination: domain.Pagination{
			CurrentPage: 1,
			TotalPages:  2,
			TotalSize:   15,
		},
		Users: []domain.User{
			{
				ID:        uuid.NewString(),
				Name: "Charlie",
				Username:  "Brown",
				Email:     "charlie@example.com",
			},
		},
	}

	mockRepo.EXPECT().List(ctx, pageNumber, pageSize).Return(expectedPaginated, nil)

	result, err := svc.List(ctx, pageNumber, pageSize)
	require.NoError(t, err)
	require.Equal(t, expectedPaginated, result)
}

func TestService_Count(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockusersRepo(ctrl)
	logger := zap.NewNop()
	svc := New(mockRepo, logger)

	ctx := context.Background()
	expectedCount := 42

	mockRepo.EXPECT().Count(ctx).Return(expectedCount, nil)

	count, err := svc.Count(ctx)
	require.NoError(t, err)
	require.Equal(t, expectedCount, count)
}
