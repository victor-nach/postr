package repositories

import (
	"fmt"
	"testing"

	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/victor-nach/postr-backend/internal/domain"
)

func TestUserRepository_Get(t *testing.T) {
	cleanUsers(t)

	user := domain.User{
		ID:       uuid.NewString(),
		Name:     "Get Test User",
		Username: "gettestuser",
		Email:    "get@example.com",
		Phone:    "9876543210",
		Address: domain.Address{
			ID:      uuid.NewString(),
			Street:  "456 Get St",
			City:    "Getville",
			State:   "GT",
			Zipcode: "67890",
		},
	}
	err := usersrepo.Create(testCtx, &user)
	require.NoError(t, err)

	retrieved, err := usersrepo.Get(testCtx, user.ID)
	require.NoError(t, err)
	assert.Equal(t, user.Name, retrieved.Name)
	assert.Equal(t, user.Username, retrieved.Username)
	assert.Equal(t, user.Email, retrieved.Email)
	assert.Equal(t, user.Phone, retrieved.Phone)
	require.NotNil(t, retrieved.Address)
	assert.Equal(t, user.Address.Street, retrieved.Address.Street)
	assert.Equal(t, user.Address.City, retrieved.Address.City)
	assert.Equal(t, user.Address.State, retrieved.Address.State)
	assert.Equal(t, user.Address.Zipcode, retrieved.Address.Zipcode)

	_, err = usersrepo.Get(testCtx, "non-existent-id")
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

func TestUserRepository_Count(t *testing.T) {
	cleanUsers(t)

	count, err := usersrepo.Count(testCtx)
	require.NoError(t, err)
	assert.Equal(t, 0, count)

	users := []domain.User{
		{
			ID:       uuid.NewString(),
			Name:     "User One",
			Username: "userone",
			Email:    "user1@example.com",
			Phone:    "1111111111",
		},
		{
			ID:       uuid.NewString(),
			Name:     "User Two",
			Username: "usertwo",
			Email:    "user2@example.com",
			Phone:    "2222222222",
		},
	}
	err = db.WithContext(testCtx).Create(&users).Error
	require.NoError(t, err)

	count, err = usersrepo.Count(testCtx)
	require.NoError(t, err)
	assert.Equal(t, len(users), count)
}

func TestUserRepository_List(t *testing.T) {
	cleanUsers(t)

	var users []domain.User
	for i := 1; i <= 5; i++ {
		users = append(users, domain.User{
			ID:       uuid.NewString(),
			Name:     fmt.Sprintf("User %d", i),
			Username: fmt.Sprintf("user%d", i),
			Email:    fmt.Sprintf("user%d@example.com", i),
			Phone:    fmt.Sprintf("12345678%02d", i),
		})
	}
	err := db.WithContext(testCtx).Create(&users).Error
	require.NoError(t, err)

	paginated, err := usersrepo.List(testCtx, 1, 2)
	require.NoError(t, err)
	assert.Equal(t, 1, paginated.Pagination.CurrentPage)
	assert.Equal(t, 3, paginated.Pagination.TotalPages)
	assert.Equal(t, 5, paginated.Pagination.TotalSize)
	assert.Len(t, paginated.Users, 2)

	paginated, err = usersrepo.List(testCtx, 2, 2)
	require.NoError(t, err)
	assert.Equal(t, 2, paginated.Pagination.CurrentPage)
	assert.Len(t, paginated.Users, 2)
}

func TestUserRepository_Validate(t *testing.T) {
	cleanUsers(t)

	user := domain.User{
		ID:       uuid.NewString(),
		Name:     "Validate User",
		Username: "validateuser",
		Email:    "validate@example.com",
		Phone:    "0000000000",
	}
	err := usersrepo.Create(testCtx, &user)
	require.NoError(t, err)

	// Validate an existing user.
	err = usersrepo.Validate(testCtx, user.ID)
	require.NoError(t, err)

	// Validate a non-existent user.
	err = usersrepo.Validate(testCtx, "non-existent-id")
	assert.Error(t, err)
	assert.Equal(t, domain.ErrUserNotFound, err)
}

func cleanUsers(t *testing.T) {
	err := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&domain.User{}).Error
	require.NoError(t, err)
}
