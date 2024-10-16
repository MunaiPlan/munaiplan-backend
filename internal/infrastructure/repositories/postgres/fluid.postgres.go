package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/infrastructure/drivers/postgres/models"
	"gorm.io/gorm"
)

type fluidsRepository struct {
	db *gorm.DB
}

func NewFluidsRepository(db *gorm.DB) *fluidsRepository {
	return &fluidsRepository{db: db}
}

// CreateFluid creates a new fluid in the database.
func (r *fluidsRepository) CreateFluid(ctx context.Context, caseID string, fluid *entities.Fluid) error {
	gormFluid := toGormFluid(fluid)
	var err error
	if gormFluid.CaseID, err = uuid.Parse(caseID); err != nil {
		return err
	}

	err = r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(gormFluid).Error; err != nil {
			return err
		}
		return nil
	})

	return err
}

// GetFluidByID retrieves a fluid by its ID from the database.
func (r *fluidsRepository) GetFluidByID(ctx context.Context, id string) (*entities.Fluid, error) {
	var fluid models.Fluid
	result := r.db.WithContext(ctx).Preload("FluidBaseType").Preload("BaseFluid").Where("id = ?", id).First(&fluid)
	if result.Error != nil {
		return nil, result.Error
	}

	return toDomainFluid(&fluid), nil
}

// GetFluids retrieves all fluids for a specific case ID.
func (r *fluidsRepository) GetFluids(ctx context.Context, caseID string) ([]*entities.Fluid, error) {
	var fluids []*models.Fluid
	var res []*entities.Fluid
	result := r.db.WithContext(ctx).Preload("FluidBaseType").Preload("BaseFluid").Where("case_id = ?", caseID).Find(&fluids)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, fluid := range fluids {
		res = append(res, toDomainFluid(fluid))
	}

	return res, nil
}

// UpdateFluid updates a fluid in the database.
func (r *fluidsRepository) UpdateFluid(ctx context.Context, fluid *entities.Fluid) (*entities.Fluid, error) {
	var updatedFluid models.Fluid

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var existingFluid models.Fluid
		if err := tx.Preload("FluidBaseType").Preload("BaseFluid").Where("id = ?", fluid.ID).First(&existingFluid).Error; err != nil {
			return err
		}

		gormFluid := toGormFluid(fluid)
		if err := tx.Model(&existingFluid).Updates(gormFluid).Error; err != nil {
			return err
		}

		if err := tx.Preload("FluidBaseType").Preload("BaseFluid").Where("id = ?", fluid.ID).First(&updatedFluid).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return toDomainFluid(&updatedFluid), nil
}

// DeleteFluid deletes a fluid from the database.
func (r *fluidsRepository) DeleteFluid(ctx context.Context, id string) error {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		result := tx.Where("id = ?", id).Delete(&models.Fluid{})
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}
		return nil
	})

	return err
}

// GetFluidTypes retrieves all fluid types from the database.
func (r *fluidsRepository) GetFluidTypes(ctx context.Context) ([]*entities.FluidType, error) {
	var fluidTypes []*models.FluidType
	var res []*entities.FluidType
	result := r.db.WithContext(ctx).Find(&fluidTypes)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, fluidType := range fluidTypes {
		res = append(res, toDomainFluidType(fluidType))
	}

	return res, nil
}