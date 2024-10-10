
package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/infrastructure/drivers/postgres/models"
	"gorm.io/gorm"
)

type porePressuresRepository struct {
	db *gorm.DB
}

// NewPorePressureRepository initializes a new PorePressure repository.
func NewPorePressuresRepository(db *gorm.DB) *porePressuresRepository {
	return &porePressuresRepository{db: db}
}

// CreatePorePressure creates a new PorePressure entry in the database.
func (r *porePressuresRepository) CreatePorePressure(ctx context.Context, caseID string, porePressure *entities.PorePressure) error {
	gormPorePressure := toGormPorePressure(porePressure)
	var err error
	if gormPorePressure.CaseID, err = uuid.Parse(caseID); err != nil {
		return err
	}

	return r.db.WithContext(ctx).Create(gormPorePressure).Error
}

// GetPorePressureByID retrieves a PorePressure entry by its ID.
func (r *porePressuresRepository) GetPorePressureByID(ctx context.Context, id string) (*entities.PorePressure, error) {
	var porePressureModel models.PorePressure
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&porePressureModel).Error
	if err != nil {
		return nil, err
	}

	return toDomainPorePressure(&porePressureModel), nil
}

// GetPorePressures retrieves all PorePressure entries associated with a specific Case ID.
func (r *porePressuresRepository) GetPorePressures(ctx context.Context, caseID string) ([]*entities.PorePressure, error) {
	var porePressureModels []models.PorePressure
	var porePressures []*entities.PorePressure

	err := r.db.WithContext(ctx).Where("case_id = ?", caseID).Find(&porePressureModels).Error
	if err != nil {
		return nil, err
	}

	for _, ppModel := range porePressureModels {
		porePressures = append(porePressures, toDomainPorePressure(&ppModel))
	}

	return porePressures, nil
}

// UpdatePorePressure updates an existing PorePressure entry in the database.
func (r *porePressuresRepository) UpdatePorePressure(ctx context.Context, porePressure *entities.PorePressure) (*entities.PorePressure, error) {
	var existingPorePressure models.PorePressure

	// Find the existing record
	err := r.db.WithContext(ctx).Where("id = ?", porePressure.ID).First(&existingPorePressure).Error
	if err != nil {
		return nil, err
	}

	// Update the record
	gormPorePressure := toGormPorePressure(porePressure)
	if err := r.db.WithContext(ctx).Model(&existingPorePressure).Updates(gormPorePressure).Error; err != nil {
		return nil, err
	}

	return toDomainPorePressure(&existingPorePressure), nil
}

// DeletePorePressure deletes a PorePressure entry from the database by its ID.
func (r *porePressuresRepository) DeletePorePressure(ctx context.Context, id string) error {
	err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.PorePressure{}).Error
	if err != nil {
		return err
	}

	return nil
}
