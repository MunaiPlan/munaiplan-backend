package postgres

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	domainErrors "github.com/munaiplan/munaiplan-backend/internal/domain/errors"
	"github.com/munaiplan/munaiplan-backend/internal/infrastructure/drivers/postgres/models"
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
    tempUser.OrganizationID, _ = uuid.Parse(organizationId)
    return r.db.WithContext(ctx).Create(&tempUser).Error
}

func (r *usersRepository) GetByEmail(ctx context.Context, organizationId string, email string) (*entities.User, error) {
    var user models.User
    err := r.db.WithContext(ctx).Where("email = ? AND organization_id = ?", email, organizationId).First(&user).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, domainErrors.ErrUserNotFound
    }

    res := toDomainUser(&user)
    return &res, err
}

// Todo() Decide on need
// func (r *userRepository) Update(ctx context.Context, user domain.User) error {
//     return r.db.WithContext(ctx).Save(&user).Error
// }

// ToDomainUser maps the GORM User model to the domain User entity.
func toDomainUser(userModel *models.User) entities.User {
    return entities.User{
        ID:        userModel.ID.String(),
        Name:      userModel.Name,
        Email:     userModel.Email,
        Password:  userModel.Password,
        Phone:     userModel.Phone,
        CreatedAt: userModel.CreatedAt,
    }
}

// ToGormUser maps the domain User entity to the GORM User model.
func toGormUser(user *entities.User) models.User {
    return models.User{
        Name:      user.Name,
        Email:     user.Email,
        Password:  user.Password,
        Phone:     user.Phone,
    }
}