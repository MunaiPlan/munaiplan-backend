package postgres

import (
	"context"
	"reflect"

	"github.com/google/uuid"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/infrastructure/drivers/postgres/models"
	"github.com/munaiplan/munaiplan-backend/internal/infrastructure/types"
	"gorm.io/gorm"
)

type sitesRepository struct {
	db *gorm.DB
}

func NewSitesRepository(db *gorm.DB) *sitesRepository {
	return &sitesRepository{db: db}
}

func (r *sitesRepository) CreateSite(ctx context.Context, fieldID string, site *entities.Site) error {
	gormSite := toGormSite(site)
	fieldId, err := uuid.Parse(fieldID)
	if err != nil {
		return err
	}
	gormSite.FieldID = fieldId

	result := r.db.WithContext(ctx).Create(gormSite)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *sitesRepository) GetSiteByID(ctx context.Context, id string) (*entities.Site, error) {
	var site models.Site
	result := r.db.WithContext(ctx).Where("id = ?", id).First(&site)
	if result.Error != nil {
		return nil, result.Error
	}

	res := toDomainSite(&site)
	return res, nil
}

func (r *sitesRepository) GetSites(ctx context.Context, fieldID string) ([]*entities.Site, error) {
	var sites []*models.Site
	var res []*entities.Site
	result := r.db.WithContext(ctx).Where("field_id = ?", fieldID).Find(&sites)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, site := range sites {
		temp := toDomainSite(site)
		res = append(res, temp)
	}
	return res, nil
}

func (r *sitesRepository) UpdateSite(ctx context.Context, site *entities.Site) (*entities.Site, error) {
	gormSite := toGormSite(site)
	oldSite := models.Site{}
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		query := tx.WithContext(ctx).Where("id = ?", site.ID).First(&oldSite)
		if query.Error != nil {
			return query.Error
		}

		if reflect.DeepEqual(&gormSite, &oldSite) {
			return types.ErrSiteNotChanged
		}

		err := tx.WithContext(ctx).Model(&oldSite).Updates(gormSite).Error
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	res := toDomainSite(&oldSite)

	return res, nil
}

func (r *sitesRepository) DeleteSite(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Site{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}