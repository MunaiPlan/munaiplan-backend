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

type wellsRepository struct {
	db *gorm.DB
}

func NewWellsRepository(db *gorm.DB) *wellsRepository {
	return &wellsRepository{db: db}
}

func (r *wellsRepository) CreateWell(ctx context.Context, siteID string, well *entities.Well) error {
	gormWell := toGormWell(well)
	siteId, err := uuid.Parse(siteID)
	if err != nil {
		return err
	}
	gormWell.SiteID = siteId

	result := r.db.WithContext(ctx).Create(gormWell)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *wellsRepository) GetWellByID(ctx context.Context, id string) (*entities.Well, error) {
	var well models.Well
	result := r.db.WithContext(ctx).Where("id = ?", id).First(&well)
	if result.Error != nil {
		return nil, result.Error
	}

	res := toDomainWell(&well)
	return res, nil
}

func (r *wellsRepository) GetWells(ctx context.Context, siteID string) ([]*entities.Well, error) {
	var wells []*models.Well
	var res []*entities.Well
	result := r.db.WithContext(ctx).Where("site_id = ?", siteID).Find(&wells)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, well := range wells {
		temp := toDomainWell(well)
		res = append(res, temp)
	}
	return res, nil
}

func (r *wellsRepository) UpdateWell(ctx context.Context, well *entities.Well) (*entities.Well, error) {
	gormWell := toGormWell(well)
	oldWell := models.Well{}
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		query := tx.WithContext(ctx).Where("id = ?", well.ID).First(&oldWell)
		if query.Error != nil {
			return query.Error
		}

		if reflect.DeepEqual(&gormWell, &oldWell) {
			return types.ErrWellNotChanged
		}

		err := tx.WithContext(ctx).Model(&oldWell).Updates(gormWell).Error
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	res := toDomainWell(&oldWell)

	return res, nil
}

func (r *wellsRepository) DeleteWell(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Well{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
