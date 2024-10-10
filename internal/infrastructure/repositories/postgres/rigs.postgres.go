package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/infrastructure/drivers/postgres/models"
	"gorm.io/gorm"
)

type rigsRepository struct {
	db *gorm.DB
}

// NewRigsRepository creates a new instance of rigsRepository
func NewRigsRepository(db *gorm.DB) *rigsRepository {
	return &rigsRepository{db: db}
}

// CreateRig creates a new rig in the database.
func (r *rigsRepository) CreateRig(ctx context.Context, caseID string, rig *entities.Rig) error {
	gormRig := toGormRig(rig)
	var err error
	if gormRig.CaseID, err = uuid.Parse(caseID); err != nil {
		return err
	}

	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(gormRig).Error; err != nil {
			return err
		}
		return nil
	})
}

// GetRigByID retrieves a rig by its ID from the database.
func (r *rigsRepository) GetRigByID(ctx context.Context, id string) (*entities.Rig, error) {
	var rigModel models.Rig
	result := r.db.WithContext(ctx).Where("id = ?", id).First(&rigModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return toDomainRig(&rigModel), nil
}

// GetRigs retrieves all rigs associated with a specific case ID from the database.
func (r *rigsRepository) GetRigs(ctx context.Context, caseID string) ([]*entities.Rig, error) {
	var rigsModel []models.Rig
	var res []*entities.Rig
	result := r.db.WithContext(ctx).Where("case_id = ?", caseID).Find(&rigsModel)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, rig := range rigsModel {
		res = append(res, toDomainRig(&rig))
	}
	return res, nil
}

// UpdateRig updates an existing rig in the database.
func (r *rigsRepository) UpdateRig(ctx context.Context, rig *entities.Rig) (*entities.Rig, error) {
	var existingRig models.Rig

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", rig.ID).First(&existingRig).Error; err != nil {
			return err
		}

		// Convert to GORM model and update fields
		gormRig := toGormRig(rig)
		if err := tx.Model(&existingRig).Updates(gormRig).Error; err != nil {
			return err
		}

		// Reload updated rig into existingRig
		if err := tx.Where("id = ?", rig.ID).First(&existingRig).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return toDomainRig(&existingRig), nil
}

// DeleteRig deletes a rig from the database by its ID.
func (r *rigsRepository) DeleteRig(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		result := tx.Where("id = ?", id).Delete(&models.Rig{})
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}
		return nil
	})
}
