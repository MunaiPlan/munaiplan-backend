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

type designsRepository struct {
	db *gorm.DB
}

func NewDesignsRepository(db *gorm.DB) *designsRepository {
	return &designsRepository{db: db}
}

func (r *designsRepository) CreateDesign(ctx context.Context, wellboreID string, design *entities.Design) error {
	gormDesign := toGormDesign(design)
	wellboreId, err := uuid.Parse(wellboreID)
	if err != nil {
		return err
	}
	gormDesign.WellboreID = wellboreId

	result := r.db.WithContext(ctx).Create(gormDesign)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *designsRepository) GetDesignByID(ctx context.Context, id string) (*entities.Design, error) {
	var design models.Design
	result := r.db.WithContext(ctx).Where("id = ?", id).First(&design)
	if result.Error != nil {
		return nil, result.Error
	}

	res := toDomainDesign(&design)
	return res, nil
}

func (r *designsRepository) GetDesigns(ctx context.Context, wellboreID string) ([]*entities.Design, error) {
	var designs []*models.Design
	var res []*entities.Design
	result := r.db.WithContext(ctx).Where("wellbore_id = ?", wellboreID).Find(&designs)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, design := range designs {
		temp := toDomainDesign(design)
		res = append(res, temp)
	}
	return res, nil
}

func (r *designsRepository) UpdateDesign(ctx context.Context, design *entities.Design) (*entities.Design, error) {
	gormDesign := toGormDesign(design)
	oldDesign := models.Design{}
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		query := tx.WithContext(ctx).Where("id = ?", design.ID).First(&oldDesign)
		if query.Error != nil {
			return query.Error
		}

		if reflect.DeepEqual(&gormDesign, &oldDesign) {
			return types.ErrDesignNotChanged
		}

		err := tx.WithContext(ctx).Model(&oldDesign).Updates(gormDesign).Error
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	res := toDomainDesign(&oldDesign)

	return res, nil
}

func (r *designsRepository) DeleteDesign(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Design{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
