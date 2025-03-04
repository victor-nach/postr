package domain

import (
	"context"
)

//go:generate mockgen -destination=./mocks/user_mock.go -package=mocks github.com/victor-nach/postr-backend/internal/domain UserService
type UserService interface {
	Get(ctx context.Context, id string) (*User, error)
	List(ctx context.Context, pageNumber int, pageSize int) (PaginatedUsers, error)
	Count(ctx context.Context) (int, error)
}

//go:generate mockgen -destination=./mocks/post_mock.go -package=mocks github.com/victor-nach/postr-backend/internal/domain PostService
type PostService interface {
	Create(ctx context.Context, post *Post) error
	List(ctx context.Context, userId string) ([]Post, error)
	Delete(ctx context.Context, id string) error
}
