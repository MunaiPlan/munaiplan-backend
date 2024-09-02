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

type holesRepository struct {
	db *gorm.DB
}

func NewHolesRepository(db *gorm.DB) *holesRepository {
	return &holesRepository{db: db}
}

func (r *holesRepository) CreateHole(ctx context.Context, caseID string, hole *entities.Hole) error {
	gormHole := toGormHole(hole)
	caseId, err := uuid.Parse(caseID)
	if err != nil {
		return err
	}
	gormHole.CaseID = caseId

	result := r.db.WithContext(ctx).Create(gormHole)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *holesRepository) GetHoleByID(ctx context.Context, id string) (*entities.Hole, error) {
	var hole models.Hole
	var res entities.Hole
	result := r.db.WithContext(ctx).Where("id = ?", id).First(&hole)
	if result.Error != nil {
		return nil, result.Error
	}

	res = toDomainHole(&hole)
	return &res, nil
}

func (r *holesRepository) GetHoles(ctx context.Context, caseID string) ([]*entities.Hole, error) {
	var holes []*models.Hole
	var res []*entities.Hole
	result := r.db.WithContext(ctx).Where("case_id = ?", caseID).Find(&holes)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, hole := range holes {
		temp := toDomainHole(hole)
		res = append(res, &temp)
	}
	return res, nil
}

func (r *holesRepository) UpdateHole(ctx context.Context, hole *entities.Hole) (*entities.Hole, error) {
	gormHole := toGormHole(hole)
	oldHole := models.Hole{}
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		query := tx.WithContext(ctx).Where("id = ?", hole.ID).First(&oldHole)
		if query.Error != nil {
			return query.Error
		}

		if reflect.DeepEqual(&gormHole, &oldHole) {
			return types.ErrHoleNotChanged
		}

		err := tx.WithContext(ctx).Model(&oldHole).Updates(gormHole).Error
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	res := toDomainHole(&oldHole)

	return &res, nil
}

func (r *holesRepository) DeleteHole(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Hole{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}