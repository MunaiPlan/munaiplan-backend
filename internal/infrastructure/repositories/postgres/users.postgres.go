package postgres

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	domainErrors "github.com/munaiplan/munaiplan-backend/internal/domain/types"
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
    tempUser.OrganizationID = uuid.MustParse(organizationId)
    return r.db.WithContext(ctx).Create(&tempUser).Error
}

// GetByEmail and GetByID only return live accounts: entities.User has no DeletedAt,
// so GORM's soft-delete scope does not apply and the filter must be explicit.
func (r *usersRepository) GetByEmail(ctx context.Context, email string) (*entities.User, error) {
	return r.getOne(ctx, "email = ?", email)
}

func (r *usersRepository) GetByID(ctx context.Context, id string) (*entities.User, error) {
	if _, err := uuid.Parse(id); err != nil {
		return nil, domainErrors.ErrUserNotFound
	}
	return r.getOne(ctx, "id = ?", id)
}

func (r *usersRepository) getOne(ctx context.Context, condition string, value string) (*entities.User, error) {
	var user entities.User
	err := r.db.WithContext(ctx).Table("users").
		Where(condition+" AND deleted_at IS NULL", value).
		First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domainErrors.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
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