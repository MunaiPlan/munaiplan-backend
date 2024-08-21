package postgres

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/infrastructure/drivers/postgres/models"
	"github.com/munaiplan/munaiplan-backend/internal/infrastructure/types"
	"gorm.io/gorm"
)

type usersRepository struct {
    db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) *usersRepository {
    return &usersRepository{db: db}
}

func (r *usersRepository) Create(ctx context.Context, organizationId string, user *entities.User) error {
    tempUser := toGormUser(user)
    tempUser.OrganizationID = uuid.MustParse(organizationId)
    return r.db.WithContext(ctx).Create(&tempUser).Error
}

func (r *usersRepository) GetByEmail(ctx context.Context, email string) (*entities.User, error) {
    var user entities.User

    err := r.db.WithContext(ctx).
        Select("users.*, users.organization_id").
        Where("email = ?", email).
        First(&user).Error

    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, types.ErrUserNotFound
    }

    return &user, err
}

// Todo() Decide on need
// func (r *userRepository) Update(ctx context.Context, user domain.User) error {
//     return r.db.WithContext(ctx).Save(&user).Error
// }


// ToGormUser maps the domain User entity to the GORM User model.
func toGormUser(user *entities.User) models.User {
    return models.User{
        Name:      user.Name,
        Email:     user.Email,
        Password:  user.Password,
        Phone:     user.Phone,
    }
}