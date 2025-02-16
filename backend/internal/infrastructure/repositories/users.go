package repositories

import (
	"context"
	"math"

	"gorm.io/gorm"

	"github.com/victor-nach/postr-backend/internal/domain"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

type userAddressJoin  struct {
	ID        string  `gorm:"column:id"`
	Name      string  `gorm:"column:name"`
	Username  string  `gorm:"column:username"`
	Email     string  `gorm:"column:email"`
	Phone     string  `gorm:"column:phone"`
	AddressID *string `gorm:"column:address_id"`
	Street    *string `gorm:"column:street"`
	City      *string `gorm:"column:city"`
	State     *string `gorm:"column:state"`
	Zipcode   *string `gorm:"column:zipcode"`
}

func (r *userRepository) Get(ctx context.Context, id string) (*domain.User, error) {
	var result userAddressJoin
	if err := r.db.WithContext(ctx).
		Table("users").
		Joins("LEFT JOIN addresses ON addresses.user_id = users.id").
		Where("users.id = ?", id).
		Select(`
			users.id,
			users.name,
			users.username,
			users.email,
			users.phone,
			addresses.id as address_id,
			addresses.street,
			addresses.city,
			addresses.state,
			addresses.zipcode
		`).
		First(&result).Error; err != nil {
		return nil, err
	}

	user := &domain.User{
		ID:       result.ID,
		Name:     result.Name,
		Username: result.Username,
		Email:    result.Email,
		Phone:    result.Phone,
	}

	if result.AddressID != nil {
		user.Address = domain.Address{
			ID:      *result.AddressID,
			UserID:  result.ID,
			Street:  *result.Street,
			City:    *result.City,
			State:   *result.State,
			Zipcode: *result.Zipcode,
		}
	}

	return user, nil
}

func (r *userRepository) Count(ctx context.Context) (int, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&domain.User{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

// func (r *userRepository) List(ctx context.Context, pageNumber int, pageSize int) (domain.PaginatedUsers, error) {
// 	var users []domain.User
// 	var total int64

// 	if err := r.db.WithContext(ctx).Model(&domain.User{}).Count(&total).Error; err != nil {
// 		return domain.PaginatedUsers{}, err
// 	}

// 	offset := (pageNumber - 1) * pageSize
// 	if err := r.db.WithContext(ctx).
// 		Order("created_at DESC").
// 		Offset(offset).
// 		Limit(pageSize).
// 		Find(&users).Error; err != nil {
// 		return domain.PaginatedUsers{}, err
// 	}

// 	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
// 	paginated := domain.PaginatedUsers{
// 		Pagination: domain.Pagination{
// 			CurrentPage: pageNumber,
// 			TotalPages:  totalPages,
// 			TotalSize:   int(total),
// 		},
// 		Users: users,
// 	}

// 	return paginated, nil
// }

func (r *userRepository) List(ctx context.Context, pageNumber int, pageSize int) (domain.PaginatedUsers, error) {
	var total int64
	if err := r.db.WithContext(ctx).
		Table("users").
		Count(&total).Error; err != nil {
		return domain.PaginatedUsers{}, err
	}

	offset := (pageNumber - 1) * pageSize

	var results []userAddressJoin
	if err := r.db.WithContext(ctx).
		Table("users").
		Joins("LEFT JOIN addresses ON addresses.user_id = users.id").
		Order("users.id DESC").
		Offset(offset).
		Limit(pageSize).
		Select(`
			users.id,
			users.name,
			users.username,
			users.email,
			users.phone,
			addresses.id as address_id,
			addresses.street,
			addresses.city,
			addresses.state,
			addresses.zipcode
		`).
		Scan(&results).Error; err != nil {
		return domain.PaginatedUsers{}, err
	}

	var users []domain.User
	for _, res := range results {
		user := domain.User{
			ID:       res.ID,
			Name:     res.Name,
			Username: res.Username,
			Email:    res.Email,
			Phone:    res.Phone,
		}
		if res.AddressID != nil {
			user.Address = domain.Address{
				ID:      *res.AddressID,
				UserID:  res.ID,
				Street:  *res.Street,
				City:    *res.City,
				State:   *res.State,
				Zipcode: *res.Zipcode,
			}
		}
		users = append(users, user)
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	paginated := domain.PaginatedUsers{
		Pagination: domain.Pagination{
			CurrentPage: pageNumber,
			TotalPages:  totalPages,
			TotalSize:   int(total),
		},
		Users: users,
	}

	return paginated, nil
}


func (r *userRepository) Validate(ctx context.Context, userID string) error {
	var count int64
	if err := r.db.WithContext(ctx).Model(&domain.User{}).Where("id = ?", userID).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		return domain.ErrUserNotFound
	}
	return nil
}

func (r *userRepository) Create(ctx context.Context, user *domain.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}
