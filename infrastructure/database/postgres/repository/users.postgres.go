package repository

import (
	"context"
	"errors"

	"github.com/munaiplan/munaiplan-backend/infrastructure/database/postgres/models"
	"github.com/munaiplan/munaiplan-backend/internal/domain"
	"gorm.io/gorm"
)

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
    return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *domain.User) error {
    tempUser := ToGormUser(user)
    return r.db.WithContext(ctx).Create(&tempUser).Error
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
    var user models.User
    err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, domain.ErrUserNotFound
    }

    res := ToDomainUser(&user)
    return &res, err
}

// Todo() Decide on need
// func (r *userRepository) Update(ctx context.Context, user domain.User) error {
//     return r.db.WithContext(ctx).Save(&user).Error
// }

// ToDomainUser maps the GORM User model to the domain User entity.
func ToDomainUser(userModel *models.User) domain.User {
    return domain.User{
        ID:        userModel.ID.String(),
        Name:      userModel.Name,
        Email:     userModel.Email,
        Password:  userModel.Password,
        Phone:     userModel.Phone,
        CreatedAt: userModel.CreatedAt,
    }
}

// ToGormUser maps the domain User entity to the GORM User model.
func ToGormUser(user *domain.User) models.User {
    return models.User{
        Name:      user.Name,
        Email:     user.Email,
        Password:  user.Password,
        Phone:     user.Phone,
    }
}