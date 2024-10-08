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

type fieldsRepository struct {
	db *gorm.DB
}

func NewFieldsRepository(db *gorm.DB) *fieldsRepository {
	return &fieldsRepository{db: db}
}

func (r *fieldsRepository) CreateField(ctx context.Context, companyID string, field *entities.Field) error {
	gormField := toGormField(field)
	companyId, err := uuid.Parse(companyID)
	if err != nil {
		return err
	}
	gormField.CompanyID = companyId

	result := r.db.WithContext(ctx).Create(gormField)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *fieldsRepository) GetFieldByID(ctx context.Context, id string) (*entities.Field, error) {
	var field models.Field
	result := r.db.WithContext(ctx).Where("id = ?", id).First(&field)
	if result.Error != nil {
		return nil, result.Error
	}

	res := toDomainField(&field)
	return res, nil
}

func (r *fieldsRepository) GetFields(ctx context.Context, companyID string) ([]*entities.Field, error) {
	var fields []*models.Field
	var res []*entities.Field
	result := r.db.WithContext(ctx).Where("company_id = ?", companyID).Find(&fields)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, field := range fields {
		temp := toDomainField(field)
		res = append(res, temp)
	}
	return res, nil
}

func (r *fieldsRepository) UpdateField(ctx context.Context, field *entities.Field) (*entities.Field, error) {
	gormField := toGormField(field)
	oldField := models.Field{}
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		query := tx.WithContext(ctx).Where("id = ?", field.ID).First(&oldField)
		if query.Error != nil {
			return query.Error
		}

		if reflect.DeepEqual(&gormField, &oldField) {
			return types.ErrFieldNotChanged
		}

		err := tx.WithContext(ctx).Model(&oldField).Updates(gormField).Error
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	res := toDomainField(&oldField)

	return res, nil
}

func (r *fieldsRepository) DeleteField(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Field{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
