package postgres

import (
	"context"
	"fmt"
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

// CreateHole creates a new hole and associated caisings in the database.
func (r *holesRepository) CreateHole(ctx context.Context, caseID string, hole *entities.Hole) error {
	gormHole := toGormHole(hole)
	var err error
	if gormHole.CaseID, err = uuid.Parse(caseID); err != nil {
		return err
	}

	fmt.Println("len of input.Body.Caisings at repo", len(hole.Caisings))

	err = r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(gormHole).Error; err != nil {
			return err
		}
		return nil
	})

	return err
}

// GetHoleByID retrieves a hole by its ID from the database, along with associated caisings.
func (r *holesRepository) GetHoleByID(ctx context.Context, id string) (*entities.Hole, error) {
	var hole models.Hole
	result := r.db.WithContext(ctx).Preload("Caisings").Where("id = ?", id).First(&hole)
	if result.Error != nil {
		return nil, result.Error
	}

	return toDomainHole(&hole), nil
}

// GetHoles retrieves all holes associated with a case ID from the database.
func (r *holesRepository) GetHoles(ctx context.Context, caseID string) ([]*entities.Hole, error) {
	var holes []*models.Hole
	var res []*entities.Hole
	result := r.db.WithContext(ctx).Preload("Caisings").Where("case_id = ?", caseID).Find(&holes)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, hole := range holes {
		temp := toDomainHole(hole)
		res = append(res, temp)
	}
	return res, nil
}

// UpdateHole updates a hole and associated caisings in the database.
func (r *holesRepository) UpdateHole(ctx context.Context, hole *entities.Hole) (*entities.Hole, error) {
	var updatedHole models.Hole

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var existingHole models.Hole
		if err := tx.Preload("Caisings").Where("id = ?", hole.ID).First(&existingHole).Error; err != nil {
			return err
		}

		gormHole := toGormHole(hole)
		if err := tx.Model(&existingHole).Updates(gormHole).Error; err != nil {
			return err
		}

		existingCaisingsMap := make(map[uuid.UUID]models.Caising)
		for _, caising := range existingHole.Caisings {
			existingCaisingsMap[caising.ID] = caising
		}

		for _, newCaising := range gormHole.Caisings {
			if newCaising.ID == uuid.Nil {
				newCaising.HoleID = existingHole.ID
				if err := tx.Create(&newCaising).Error; err != nil {
					return err
				}
			} else if existingCaising, exists := existingCaisingsMap[newCaising.ID]; exists {
				if !reflect.DeepEqual(existingCaising, newCaising) {
					if err := tx.Model(&existingCaising).Updates(newCaising).Error; err != nil {
						return err
					}
				}
				delete(existingCaisingsMap, newCaising.ID)
			} else {
				return types.ErrCaisingIdNotFound
			}
		}

		for _, caisingToDelete := range existingCaisingsMap {
			if err := tx.Delete(&caisingToDelete).Error; err != nil {
				return err
			}
		}

		if err := tx.Preload("Caisings").Where("id = ?", hole.ID).First(&updatedHole).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return toDomainHole(&updatedHole), nil
}

// DeleteHole deletes a hole and associated caisings from the database.
func (r *holesRepository) DeleteHole(ctx context.Context, id string) error {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		result := tx.Where("id = ?", id).Delete(&models.Hole{})
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
