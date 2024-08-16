package postgres

import (
	"context"
	"fmt"

	"github.com/munaiplan/munaiplan-backend/internal/infrastructure/drivers/postgres/models"
	"gorm.io/gorm"
)

type commonRepository struct {
	db *gorm.DB
}

func NewCommonRepository(db *gorm.DB) *commonRepository {
	return &commonRepository{db: db}
}

func (r *commonRepository) CheckIfOrganizationExists(ctx context.Context, organizationId string) error {
    var count int64
    if err := r.db.WithContext(ctx).Model(&models.Organization{}).Where("id = ?", organizationId).Count(&count).Error; err != nil {
        return fmt.Errorf("error checking organization existence: %w", err)
    }
    if count == 0 {
        return fmt.Errorf("organization with id %s does not exist", organizationId)
    }
    return nil
}

func (r *commonRepository) CheckIfUserExistsByEmail(ctx context.Context, email string) error {
    var count int64
    if err := r.db.WithContext(ctx).Model(&models.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
        return fmt.Errorf("error checking user existence: %w", err)
    }
    if count == 0 {
        return fmt.Errorf("user with email %s does not exist", email)
    }
    return nil
}

func (r *commonRepository) CheckIfCompanyExists(ctx context.Context, companyId string) error {
	var count int64
	if err := r.db.WithContext(ctx).Model(&models.Company{}).Where("id = ?", companyId).Count(&count).Error; err != nil {
		return fmt.Errorf("error checking company existence: %w", err)
	}
	if count == 0 {
		return fmt.Errorf("company with id %s does not exist", companyId)
	}
	return nil
}