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

type stringsRepository struct {
	db *gorm.DB
}

func NewStringsRepository(db *gorm.DB) *stringsRepository {
	return &stringsRepository{db: db}
}

// CreateString creates a new string and its associated sections.
func (r *stringsRepository) CreateString(ctx context.Context, caseID string, stringEntity *entities.String) error {
	gormString := toGormString(stringEntity)
	caseUUID, err := uuid.Parse(caseID)
	if err != nil {
		return err
	}
	gormString.CaseID = caseUUID

	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(gormString).Error; err != nil {
			return err
		}
		return nil
	})
}

// GetStringByID retrieves a string by its ID, along with associated sections.
func (r *stringsRepository) GetStringByID(ctx context.Context, id string) (*entities.String, error) {
	var gormString models.String
	result := r.db.WithContext(ctx).Preload("Sections").Where("id = ?", id).First(&gormString)
	if result.Error != nil {
		return nil, result.Error
	}

	return toDomainString(&gormString), nil
}

// GetStrings retrieves all strings associated with a case ID, along with associated sections.
func (r *stringsRepository) GetStrings(ctx context.Context, caseID string) ([]*entities.String, error) {
	var gormStrings []models.String
	var res []*entities.String
	result := r.db.WithContext(ctx).Preload("Sections").Where("case_id = ?", caseID).Find(&gormStrings)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, gormString := range gormStrings {
		res = append(res, toDomainString(&gormString))
	}
	return res, nil
}

// UpdateString updates a string and its associated sections.
func (r *stringsRepository) UpdateString(ctx context.Context, stringEntity *entities.String) (*entities.String, error) {
	var updatedString models.String

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var existingString models.String
		if err := tx.Preload("Sections").Where("id = ?", stringEntity.ID).First(&existingString).Error; err != nil {
			return err
		}

		gormString := toGormString(stringEntity)
		if err := tx.Model(&existingString).Updates(gormString).Error; err != nil {
			return err
		}

		existingSectionsMap := make(map[uuid.UUID]models.Section)
		for _, section := range existingString.Sections {
			existingSectionsMap[section.ID] = section
		}

		for _, newSection := range gormString.Sections {
			if newSection.ID == uuid.Nil {
				newSection.StringID = existingString.ID
				if err := tx.Create(&newSection).Error; err != nil {
					return err
				}
			} else if existingSection, exists := existingSectionsMap[newSection.ID]; exists {
				if !reflect.DeepEqual(existingSection, newSection) {
					if err := tx.Model(&existingSection).Updates(newSection).Error; err != nil {
						return err
					}
				}
				delete(existingSectionsMap, newSection.ID)
			} else {
				return types.ErrSectionIdNotFound
			}
		}

		for _, sectionToDelete := range existingSectionsMap {
			if err := tx.Delete(&sectionToDelete).Error; err != nil {
				return err
			}
		}

		if err := tx.Preload("Sections").Where("id = ?", stringEntity.ID).First(&updatedString).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return toDomainString(&updatedString), nil
}

// DeleteString deletes a string and its associated sections from the database.
func (r *stringsRepository) DeleteString(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		result := tx.Where("id = ?", id).Delete(&models.String{})
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}
		return nil
	})
}
