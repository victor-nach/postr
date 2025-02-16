package handlers

import (
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/victor-nach/postr-backend/internal/domain"
)

const (
	successStatus = "success"
	errorStatus   = "error"
)

type APIResponse struct {
	Status     string             `json:"status"`
	Message    string             `json:"message"`
	Pagination *domain.Pagination `json:"pagination,omitempty"`
	Data       any                `json:"data"`
}

type Count struct {
	Count int `json:"count"`
}

type createPostRequest struct {
	UserID string `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type listUsersRequest struct {
	PageNumber int `json:"pageNumber"`
	PageSize   int `json:"pageSize"`
}

func (r listUsersRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.PageNumber, validation.Required, validation.Min(1)),
		validation.Field(&r.PageSize, validation.Required, validation.Min(1), validation.Max(100)),
	)
}

func (r createPostRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.UserID, validation.Required, validation.By(isCompactUUID)),
		validation.Field(&r.Title, validation.Required, validation.Length(1, 255)),
		validation.Field(&r.Body, validation.Required, validation.Length(1, 2000)),
	)
}
