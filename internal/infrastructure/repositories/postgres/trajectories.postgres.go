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

type trajectoriesRepository struct {
	db *gorm.DB
}

func NewTrajectoriesRepository(db *gorm.DB) *trajectoriesRepository {
	return &trajectoriesRepository{db: db}
}

// CreateTrajectory creates a new trajectory in the database.
func (r *trajectoriesRepository) CreateTrajectory(ctx context.Context, designID string, trajectory *entities.Trajectory) error {
	gormTrajectory := toGormTrajectory(trajectory)
	var err error
	if gormTrajectory.DesignID, err = uuid.Parse(designID); err != nil {
		return err
	}

	err = r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(gormTrajectory).Error; err != nil {
			return err
		}
		return nil
	})

	return err
}

// GetTrajectoryByID retrieves a trajectory by its ID from the database.
func (r *trajectoriesRepository) GetTrajectoryByID(ctx context.Context, id string) (*entities.Trajectory, error) {
	var trajectory models.Trajectory
	result := r.db.WithContext(ctx).Preload("Headers").Preload("Units").Where("id = ?", id).First(&trajectory)
	if result.Error != nil {
		return nil, result.Error
	}

	res := toDomainTrajectory(&trajectory)
	return res, nil
}

func (r *trajectoriesRepository) GetTrajectories(ctx context.Context, designID string) ([]*entities.Trajectory, error) {
	var trajectories []*models.Trajectory
	var res []*entities.Trajectory
	result := r.db.WithContext(ctx).Preload("Headers").Preload("Units").Where("design_id = ?", designID).Find(&trajectories)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, trajectory := range trajectories {
		temp := toDomainTrajectory(trajectory)
		res = append(res, temp)
	}
	return res, nil
}

// UpdateTrajectory updates a trajectory in the database.
func (r *trajectoriesRepository) UpdateTrajectory(ctx context.Context, trajectory *entities.Trajectory) (*entities.Trajectory, error) {
	var updatedTrajectory models.Trajectory

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var existingTrajectory models.Trajectory
		if err := tx.Preload("Headers").Preload("Units").Where("id = ?", trajectory.ID).First(&existingTrajectory).Error; err != nil {
			return err
		}

		gormTrajectory := toGormTrajectory(trajectory)
		if err := tx.Model(&existingTrajectory).Updates(gormTrajectory).Error; err != nil {
			return err
		}

		existingHeadersMap := make(map[uuid.UUID]models.TrajectoryHeader)
		for _, header := range existingTrajectory.Headers {
			existingHeadersMap[header.ID] = header
		}

		existingUnitsMap := make(map[uuid.UUID]models.TrajectoryUnit)
		for _, unit := range existingTrajectory.Units {
			existingUnitsMap[unit.ID] = unit
		}

		for _, newHeader := range gormTrajectory.Headers {
			if newHeader.ID == uuid.Nil {
				newHeader.TrajectoryID = existingTrajectory.ID
				if err := tx.Create(&newHeader).Error; err != nil {
					return err
				}
			} else if existingHeader, exists := existingHeadersMap[newHeader.ID]; exists {
				if !reflect.DeepEqual(existingHeader, newHeader) {
					if err := tx.Model(&existingHeader).Updates(newHeader).Error; err != nil {
						return err
					}
				}
				delete(existingHeadersMap, newHeader.ID)
			} else {
				return types.ErrTrajectoryHeaderIdNotFound
			}
		}

		for _, headerToDelete := range existingHeadersMap {
			if err := tx.Delete(&headerToDelete).Error; err != nil {
				return err
			}
		}

		for _, newUnit := range gormTrajectory.Units {
			if newUnit.ID == uuid.Nil {
				newUnit.TrajectoryID = existingTrajectory.ID
				if err := tx.Create(&newUnit).Error; err != nil {
					return err
				}
			} else if existingUnit, exists := existingUnitsMap[newUnit.ID]; exists {
				if !reflect.DeepEqual(existingUnit, newUnit) {
					if err := tx.Model(&existingUnit).Updates(newUnit).Error; err != nil {
						return err
					}
				}
				delete(existingUnitsMap, newUnit.ID)
			} else {
				return types.ErrTrajectoryUnitIdNotFound
			}
		}

		for _, unitToDelete := range existingUnitsMap {
			if err := tx.Delete(&unitToDelete).Error; err != nil {
				return err
			}
		}

		if err := tx.Preload("Headers").Preload("Units").Where("id = ?", trajectory.ID).First(&updatedTrajectory).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	res := toDomainTrajectory(&updatedTrajectory)
	return res, nil
}

// DeleteTrajectory deletes a trajectory from the database.
func (r *trajectoriesRepository) DeleteTrajectory(ctx context.Context, id string) error {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		result := tx.Where("id = ?", id).Delete(&models.Trajectory{})
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
