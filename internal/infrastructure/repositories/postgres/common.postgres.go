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

func (r *commonRepository) CheckIfFieldExists(ctx context.Context, fieldId string) error {
	var count int64
	if err := r.db.WithContext(ctx).Model(&models.Field{}).Where("id = ?", fieldId).Count(&count).Error; err != nil {
		return fmt.Errorf("error checking field existence: %w", err)
	}
	if count == 0 {
		return fmt.Errorf("field with id %s does not exist", fieldId)
	}
	return nil
}

func (r *commonRepository) CheckIfSiteExists(ctx context.Context, siteId string) error {
	var count int64
	if err := r.db.WithContext(ctx).Model(&models.Site{}).Where("id = ?", siteId).Count(&count).Error; err != nil {
		return fmt.Errorf("error checking site existence: %w", err)
	}
	if count == 0 {
		return fmt.Errorf("site with id %s does not exist", siteId)
	}
	return nil
}

func (r *commonRepository) CheckIfWellExists(ctx context.Context, wellId string) error {
	var count int64
	if err := r.db.WithContext(ctx).Model(&models.Well{}).Where("id = ?", wellId).Count(&count).Error; err != nil {
		return fmt.Errorf("error checking well existence: %w", err)
	}
	if count == 0 {
		return fmt.Errorf("well with id %s does not exist", wellId)
	}
	return nil
}

func (r *commonRepository) CheckIfWellboreExists(ctx context.Context, wellboreId string) error {
	var count int64
	if err := r.db.WithContext(ctx).Model(&models.Wellbore{}).Where("id = ?", wellboreId).Count(&count).Error; err != nil {
		return fmt.Errorf("error checking wellbore existence: %w", err)
	}
	if count == 0 {
		return fmt.Errorf("wellbore with id %s does not exist", wellboreId)
	}
	return nil
}

func (r *commonRepository) CheckIfDesignExists(ctx context.Context, designId string) error {
	var count int64
	if err := r.db.WithContext(ctx).Model(&models.Design{}).Where("id = ?", designId).Count(&count).Error; err != nil {
		return fmt.Errorf("error checking design existence: %w", err)
	}
	if count == 0 {
		return fmt.Errorf("design with id %s does not exist", designId)
	}
	return nil
}

func (r *commonRepository) CheckIfTrajectoryExists(ctx context.Context, trajectoryId string) error {
	var count int64
	if err := r.db.WithContext(ctx).Model(&models.Trajectory{}).Where("id = ?", trajectoryId).Count(&count).Error; err != nil {
		return fmt.Errorf("error checking trajectory existence: %w", err)
	}
	if count == 0 {
		return fmt.Errorf("trajectory with id %s does not exist", trajectoryId)
	}
	return nil
}