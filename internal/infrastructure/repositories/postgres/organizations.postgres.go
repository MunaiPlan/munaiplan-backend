package postgres

import (
	"context"
	"reflect"

	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/infrastructure/drivers/postgres/models"
	"github.com/munaiplan/munaiplan-backend/internal/infrastructure/types"
	"gorm.io/gorm"
)

type organizationsRepository struct {
	db *gorm.DB
}

func NewOrganizationsRepository(db *gorm.DB) *organizationsRepository {
	return &organizationsRepository{db: db}
}

func (r *organizationsRepository) CreateOrganization(ctx context.Context, organization *entities.Organization) error {
	gormOrganization := r.toGormOrganization(organization)
	result := r.db.Where("email = ? AND deleted_at IS NULL", organization.Email).FirstOrCreate(&gormOrganization)

	if result.Error != nil {
        return types.ErrGettingOrganizationByEmail
    }

    if result.RowsAffected == 0 {
        return types.ErrOrganizationExistsWithEmail
    }
	return nil
}

func (r *organizationsRepository) GetOrganizationByID(ctx context.Context, id string) (*entities.Organization, error) {
	var organization models.Organization
	query := r.db.WithContext(ctx).Where("id = ?", id).First(&organization)
	if query.Error != nil {
		return nil, query.Error
	}

	res := r.toDomainOrganization(&organization)
	return &res, nil
}

func (r *organizationsRepository) GetOrganizationByName(ctx context.Context, name string) (*entities.Organization, error) {
	var organization models.Organization
	query := r.db.WithContext(ctx).Where("name = ?", name).First(&organization)
	if query.Error != nil {
		return nil, query.Error
	}

	res := r.toDomainOrganization(&organization)
	return &res, nil
}

func (r *organizationsRepository) GetOrganizations(ctx context.Context) ([]*entities.Organization, error) {
	var organizations []*models.Organization
	res := []*entities.Organization{}
	query := r.db.WithContext(ctx).Find(&organizations)
	if query.Error != nil {
		return nil, query.Error
	}

	for _, organization := range organizations {
		temp := r.toDomainOrganization(organization)
		res = append(res, &temp)
	}
	return res, nil
}

func (r *organizationsRepository) UpdateOrganization(ctx context.Context, organization *entities.Organization) (*entities.Organization, error) {
    gormOrg := r.toGormOrganization(organization)

    err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
        var oldOrg models.Organization
        if err := tx.Where("id = ?", organization.ID).First(&oldOrg).Error; err != nil {
            return err
        }

		if reflect.DeepEqual(gormOrg, &oldOrg) {
            return nil
        }

        if gormOrg.Email != oldOrg.Email {
            var count int64
            if err := tx.Model(&models.Organization{}).
                Where("email = ? AND id != ? AND deleted_at IS NULL", gormOrg.Email, organization.ID).
                Count(&count).Error; err != nil {
                return err
            }
            if count > 0 {
                return types.ErrOrganizationExistsWithEmail
            }
        }

        return tx.Model(&oldOrg).Updates(gormOrg).Error
    })

    if err != nil {
        return nil, err
    }

    return r.GetOrganizationByID(ctx, organization.ID)
}

func (r *organizationsRepository) DeleteOrganization(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		query := tx.WithContext(ctx).Where("id = ?", id).Delete(&models.Organization{})
		if query.Error != nil {
			return query.Error
		}

		return nil
	})
}

func (r *organizationsRepository) toGormOrganization(organization *entities.Organization) *models.Organization {
	return &models.Organization{
		Name:    organization.Name,
		Email:   organization.Email,
		Phone:   organization.Phone,
		Address: organization.Address,
	}
}

func (r *organizationsRepository) toDomainOrganization(organization *models.Organization) entities.Organization {
	return entities.Organization{
		ID:      organization.ID.String(),
		Name:    organization.Name,
		Email:   organization.Email,
		Phone:   organization.Phone,
		Address: organization.Address,
	}
}
