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

type casesRepository struct {
	db *gorm.DB
}

func NewCasesRepository(db *gorm.DB) *casesRepository {
	return &casesRepository{db: db}
}

// CreateCase creates a new case in the database
func (r *casesRepository) CreateCase(ctx context.Context, trajectoryID string, caseEntity *entities.Case) error {
	gormCase := r.toGormCase(caseEntity)
	trajectoryId, err := uuid.Parse(trajectoryID)
	if err != nil {
		return err
	}
	gormCase.TrajectoryID = trajectoryId

	err = r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(gormCase).Error; err != nil {
			return err
		}
		// Related components like Strings, Holes, and Fluids would be created in their respective repositories
		return nil
	})

	return err
}

// GetCaseByID fetches a case by its ID from the database
func (r *casesRepository) GetCaseByID(ctx context.Context, id string) (*entities.Case, error) {
	var gormCase models.Case
	var res entities.Case
	result := r.db.WithContext(ctx).Where("id = ?", id).First(&gormCase)
	if result.Error != nil {
		return nil, result.Error
	}

	res = r.toDomainCase(&gormCase)
	return &res, nil
}

// GetCases fetches all cases for a given trajectory ID from the database
func (r *casesRepository) GetCases(ctx context.Context, trajectoryID string) ([]*entities.Case, error) {
	var gormCases []*models.Case
	var res []*entities.Case
	result := r.db.WithContext(ctx).Where("trajectory_id = ?", trajectoryID).Find(&gormCases)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, gormCase := range gormCases {
		temp := r.toDomainCase(gormCase)
		res = append(res, &temp)
	}
	return res, nil
}

// UpdateCase updates an existing case in the database
func (r *casesRepository) UpdateCase(ctx context.Context, caseEntity *entities.Case) (*entities.Case, error) {
	gormCase := r.toGormCase(caseEntity)
	oldCase := models.Case{}
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		query := tx.WithContext(ctx).Where("id = ?", caseEntity.ID).First(&oldCase)
		if query.Error != nil {
			return query.Error
		}

		if reflect.DeepEqual(&gormCase, &oldCase) {
			return types.ErrCaseNotChanged
		}

		err := tx.WithContext(ctx).Model(&oldCase).Updates(gormCase).Error
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	res := r.toDomainCase(&oldCase)
	return &res, nil
}

// DeleteCase deletes a case by its ID from the database
func (r *casesRepository) DeleteCase(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Case{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// toDomainCase maps the GORM Case model to the domain Case entity.
func (r *casesRepository) toDomainCase(caseModel *models.Case) entities.Case {
	return entities.Case{
		ID:              caseModel.ID.String(),
		CaseName:        caseModel.CaseName,
		CaseDescription: caseModel.CaseDescription,
		DrillDepth:      caseModel.DrillDepth,
		PipeSize:        caseModel.PipeSize,
		CreatedAt:       caseModel.CreatedAt,
	}
}

// toGormCase maps the domain Case entity to the GORM Case model.
func (r *casesRepository) toGormCase(caseEntity *entities.Case) *models.Case {
	return &models.Case{
		CaseName:        caseEntity.CaseName,
		CaseDescription: caseEntity.CaseDescription,
		DrillDepth:      caseEntity.DrillDepth,
		PipeSize:        caseEntity.PipeSize,
	}
}
