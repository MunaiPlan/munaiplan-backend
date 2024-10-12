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

type fractureGradientsRepository struct {
	db *gorm.DB
}

func NewFractureGradientsRepository(db *gorm.DB) *fractureGradientsRepository {
	return &fractureGradientsRepository{db: db}
}

// CreateFractureGradient creates a new FractureGradient record in the database.
func (r *fractureGradientsRepository) CreateFractureGradient(ctx context.Context, caseID string, fractureGradient *entities.FractureGradient) error {
	gormFractureGradient := toGormFractureGradient(fractureGradient)
	caseId, err := uuid.Parse(caseID)
	if err != nil {
		return err
	}
	gormFractureGradient.CaseID = caseId

	result := r.db.WithContext(ctx).Create(gormFractureGradient)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// GetFractureGradientByID retrieves a FractureGradient by its ID.
func (r *fractureGradientsRepository) GetFractureGradientByID(ctx context.Context, id string) (*entities.FractureGradient, error) {
	var fractureGradient models.FractureGradient
	result := r.db.WithContext(ctx).Where("id = ?", id).First(&fractureGradient)
	if result.Error != nil {
		return nil, result.Error
	}

	return toDomainFractureGradient(&fractureGradient), nil
}

// GetFractureGradients retrieves all FractureGradients associated with a given case ID.
func (r *fractureGradientsRepository) GetFractureGradients(ctx context.Context, caseID string) ([]*entities.FractureGradient, error) {
	var fractureGradients []*models.FractureGradient
	var res []*entities.FractureGradient

	result := r.db.WithContext(ctx).Where("case_id = ?", caseID).Find(&fractureGradients)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, fg := range fractureGradients {
		res = append(res, toDomainFractureGradient(fg))
	}
	return res, nil
}

// UpdateFractureGradient updates an existing FractureGradient record.
func (r *fractureGradientsRepository) UpdateFractureGradient(ctx context.Context, fractureGradient *entities.FractureGradient) (*entities.FractureGradient, error) {
	gormFractureGradient := toGormFractureGradient(fractureGradient)
	oldFractureGradient := models.FractureGradient{}

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", fractureGradient.ID).First(&oldFractureGradient).Error; err != nil {
			return err
		}

		if reflect.DeepEqual(&gormFractureGradient, &oldFractureGradient) {
			return types.ErrFractureGradientNotChanged
		}

		if err := tx.Model(&oldFractureGradient).Updates(gormFractureGradient).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return toDomainFractureGradient(&oldFractureGradient), nil
}

// DeleteFractureGradient deletes a FractureGradient record by ID.
func (r *fractureGradientsRepository) DeleteFractureGradient(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.FractureGradient{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}