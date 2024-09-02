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

type wellboresRepository struct {
	db *gorm.DB
}

func NewWellboresRepository(db *gorm.DB) *wellboresRepository {
	return &wellboresRepository{db: db}
}

func (r *wellboresRepository) CreateWellbore(ctx context.Context, wellID string, wellbore *entities.Wellbore) error {
	gormWellbore := toGormWellbore(wellbore)
	wellId, err := uuid.Parse(wellID)
	if err != nil {
		return err
	}
	gormWellbore.WellID = wellId

	result := r.db.WithContext(ctx).Create(gormWellbore)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *wellboresRepository) GetWellboreByID(ctx context.Context, id string) (*entities.Wellbore, error) {
	var wellbore models.Wellbore
	
	result := r.db.WithContext(ctx).Where("id = ?", id).First(&wellbore)
	if result.Error != nil {
		return nil, result.Error
	}

	res := toDomainWellbore(&wellbore)
	return res, nil
}

func (r *wellboresRepository) GetWellbores(ctx context.Context, wellID string) ([]*entities.Wellbore, error) {
	var wellbores []*models.Wellbore
	var res []*entities.Wellbore
	result := r.db.WithContext(ctx).Where("well_id = ?", wellID).Find(&wellbores)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, wellbore := range wellbores {
		temp := toDomainWellbore(wellbore)
		res = append(res, temp)
	}
	return res, nil
}

func (r *wellboresRepository) UpdateWellbore(ctx context.Context, wellbore *entities.Wellbore) (*entities.Wellbore, error) {
	gormWellbore := toGormWellbore(wellbore)
	oldWellbore := models.Wellbore{}
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		query := tx.WithContext(ctx).Where("id = ?", wellbore.ID).First(&oldWellbore)
		if query.Error != nil {
			return query.Error
		}

		if reflect.DeepEqual(&gormWellbore, &oldWellbore) {
			return types.ErrWellboreNotChanged
		}

		err := tx.WithContext(ctx).Model(&oldWellbore).Updates(gormWellbore).Error
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	res := toDomainWellbore(&oldWellbore)

	return res, nil
}

func (r *wellboresRepository) DeleteWellbore(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Wellbore{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}