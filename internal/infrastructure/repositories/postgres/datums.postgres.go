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

type datumsRepository struct {
	db *gorm.DB
}

func NewDatumsRepository(db *gorm.DB) *datumsRepository {
	return &datumsRepository{db: db}
}

func (r *datumsRepository) CreateDatum(ctx context.Context, caseID string, datum *entities.Datum) error {
	gormDatum := r.toGormDatum(datum)
	caseId, err := uuid.Parse(caseID)
	if err != nil {
		return err
	}
	gormDatum.CaseID = caseId

	result := r.db.WithContext(ctx).Create(gormDatum)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *datumsRepository) GetDatumByID(ctx context.Context, id string) (*entities.Datum, error) {
	var datum models.Datum
	result := r.db.WithContext(ctx).Where("id = ?", id).First(&datum)
	if result.Error != nil {
		return nil, result.Error
	}

	res := r.toDomainDatum(&datum)
	return &res, nil
}

func (r *datumsRepository) GetDatumsByCaseID(ctx context.Context, caseID string) ([]*entities.Datum, error) {
	var datums []*models.Datum
	var res []*entities.Datum
	result := r.db.WithContext(ctx).Where("case_id = ?", caseID).Find(&datums)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, datum := range datums {
		temp := r.toDomainDatum(datum)
		res = append(res, &temp)
	}
	return res, nil
}

func (r *datumsRepository) UpdateDatum(ctx context.Context, datum *entities.Datum) (*entities.Datum, error) {
	gormDatum := r.toGormDatum(datum)
	oldDatum := models.Datum{}
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		query := tx.WithContext(ctx).Where("id = ?", datum.ID).First(&oldDatum)
		if query.Error != nil {
			return query.Error
		}

		if reflect.DeepEqual(&gormDatum, &oldDatum) {
			return types.ErrDatumNotChanged
		}

		err := tx.WithContext(ctx).Model(&oldDatum).Updates(gormDatum).Error
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	res := r.toDomainDatum(&oldDatum)
	return &res, nil
}

func (r *datumsRepository) DeleteDatum(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Datum{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// toDomainDatum maps the GORM Datum model to the domain Datum entity.
func (r *datumsRepository) toDomainDatum(datumModel *models.Datum) entities.Datum {
	return entities.Datum{
		ID:                datumModel.ID.String(),
		Name:              datumModel.Name,
		SystemDescription: datumModel.SystemDescription,
		SystemElevation:   datumModel.SystemElevation,
		DatumDescription:  datumModel.DatumDescription,
		WellheadElevation: datumModel.WellheadElevation,
		DatumElevation:    datumModel.DatumElevation,
		AirGap:            datumModel.AirGap,
		GroundElevation:   datumModel.GroundElevation,
		Type:              datumModel.Type,
		CreatedAt:         datumModel.CreatedAt,
	}
}

// toGormDatum maps the domain Datum entity to the GORM Datum model.
func (r *datumsRepository) toGormDatum(datum *entities.Datum) *models.Datum {
	return &models.Datum{
		Name:              datum.Name,
		SystemDescription: datum.SystemDescription,
		SystemElevation:   datum.SystemElevation,
		DatumDescription:  datum.DatumDescription,
		WellheadElevation: datum.WellheadElevation,
		DatumElevation:    datum.DatumElevation,
		AirGap:            datum.AirGap,
		GroundElevation:   datum.GroundElevation,
		Type:              datum.Type,
	}
}
