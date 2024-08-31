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

type holesRepository struct {
	db *gorm.DB
}

func NewHolesRepository(db *gorm.DB) *holesRepository {
	return &holesRepository{db: db}
}

func (r *holesRepository) CreateHole(ctx context.Context, caseID string, hole *entities.Hole) error {
	gormHole := r.toGormHole(hole)
	caseId, err := uuid.Parse(caseID)
	if err != nil {
		return err
	}
	gormHole.CaseID = caseId

	result := r.db.WithContext(ctx).Create(gormHole)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *holesRepository) GetHoleByID(ctx context.Context, id string) (*entities.Hole, error) {
	var hole models.Hole
	var res entities.Hole
	result := r.db.WithContext(ctx).Where("id = ?", id).First(&hole)
	if result.Error != nil {
		return nil, result.Error
	}

	res = r.toDomainHole(&hole)
	return &res, nil
}

func (r *holesRepository) GetHoles(ctx context.Context, caseID string) ([]*entities.Hole, error) {
	var holes []*models.Hole
	var res []*entities.Hole
	result := r.db.WithContext(ctx).Where("case_id = ?", caseID).Find(&holes)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, hole := range holes {
		temp := r.toDomainHole(hole)
		res = append(res, &temp)
	}
	return res, nil
}

func (r *holesRepository) UpdateHole(ctx context.Context, hole *entities.Hole) (*entities.Hole, error) {
	gormHole := r.toGormHole(hole)
	oldHole := models.Hole{}
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		query := tx.WithContext(ctx).Where("id = ?", hole.ID).First(&oldHole)
		if query.Error != nil {
			return query.Error
		}

		if reflect.DeepEqual(&gormHole, &oldHole) {
			return types.ErrHoleNotChanged
		}

		err := tx.WithContext(ctx).Model(&oldHole).Updates(gormHole).Error
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	res := r.toDomainHole(&oldHole)

	return &res, nil
}

func (r *holesRepository) DeleteHole(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Hole{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// toDomainHole maps the GORM Hole model to the domain Hole entity.
func (r *holesRepository) toDomainHole(holeModel *models.Hole) entities.Hole {
	return entities.Hole{
		ID:                        holeModel.ID.String(),
		CaseID:                    holeModel.CaseID.String(),
		CreatedAt:                 holeModel.CreatedAt,
		MDTop:                     holeModel.MDTop,
		MDBase:                    holeModel.MDBase,
		Length:                    holeModel.Length,
		ShoeMD:                    holeModel.ShoeMD,
		OD:                        holeModel.OD,
		CaisingInternalDiameter:   holeModel.CaisingInternalDiameter,
		DriftInternalDiameter:     holeModel.DriftInternalDiameter,
		EffectiveHoleDiameter:     holeModel.EffectiveHoleDiameter,
		Weight:                    holeModel.Weight,
		Grade:                     holeModel.Grade,
		MinYieldStrength:          holeModel.MinYieldStrength,
		BurstRating:               holeModel.BurstRating,
		CollapseRating:            holeModel.CollapseRating,
		FrictionFactorCasing:      holeModel.FrictionFactorCasing,
		LinearCapacityCasing:      holeModel.LinearCapacityCasing,
		DescriptionCasing:         holeModel.DescriptionCasing,
		ManufacturerCasing:        holeModel.ManufacturerCasing,
		ModelCasing:               holeModel.ModelCasing,
		OpenHoleMDTop:             holeModel.OpenHoleMDTop,
		OpenHoleMDBase:            holeModel.OpenHoleMDBase,
		OpenHoleLength:            holeModel.OpenHoleLength,
		OpenHoleInternalDiameter:  holeModel.OpenHoleInternalDiameter,
		EffectiveDiameter:         holeModel.EffectiveDiameter,
		FrictionFactorOpenHole:    holeModel.FrictionFactorOpenHole,
		LinearCapacityOpenHole:    holeModel.LinearCapacityOpenHole,
		VolumeExcess:              holeModel.VolumeExcess,
		DescriptionOpenHole:       holeModel.DescriptionOpenHole,
		TrippingInCasing:          holeModel.TrippingInCasing,
		TrippingOutCasing:         holeModel.TrippingOutCasing,
		RotatingOnBottomCasing:    holeModel.RotatingOnBottomCasing,
		SlideDrillingCasing:       holeModel.SlideDrillingCasing,
		BackReamingCasing:         holeModel.BackReamingCasing,
		RotatingOffBottomCasing:   holeModel.RotatingOffBottomCasing,
		TrippingInOpenHole:        holeModel.TrippingInOpenHole,
		TrippingOutOpenHole:       holeModel.TrippingOutOpenHole,
		RotatingOnBottomOpenHole:  holeModel.RotatingOnBottomOpenHole,
		SlideDrillingOpenHole:     holeModel.SlideDrillingOpenHole,
		BackReamingOpenHole:       holeModel.BackReamingOpenHole,
		RotatingOffBottomOpenHole: holeModel.RotatingOffBottomOpenHole,
	}
}

// toGormHole maps the domain Hole entity to the GORM Hole model.
func (r *holesRepository) toGormHole(hole *entities.Hole) *models.Hole {
	return &models.Hole{
		MDTop:                     hole.MDTop,
		MDBase:                    hole.MDBase,
		Length:                    hole.Length,
		ShoeMD:                    hole.ShoeMD,
		OD:                        hole.OD,
		CaisingInternalDiameter:   hole.CaisingInternalDiameter,
		DriftInternalDiameter:     hole.DriftInternalDiameter,
		EffectiveHoleDiameter:     hole.EffectiveHoleDiameter,
		Weight:                    hole.Weight,
		Grade:                     hole.Grade,
		MinYieldStrength:          hole.MinYieldStrength,
		BurstRating:               hole.BurstRating,
		CollapseRating:            hole.CollapseRating,
		FrictionFactorCasing:      hole.FrictionFactorCasing,
		LinearCapacityCasing:      hole.LinearCapacityCasing,
		DescriptionCasing:         hole.DescriptionCasing,
		ManufacturerCasing:        hole.ManufacturerCasing,
		ModelCasing:               hole.ModelCasing,
		OpenHoleMDTop:             hole.OpenHoleMDTop,
		OpenHoleMDBase:            hole.OpenHoleMDBase,
		OpenHoleLength:            hole.OpenHoleLength,
		OpenHoleInternalDiameter:  hole.OpenHoleInternalDiameter,
		EffectiveDiameter:         hole.EffectiveDiameter,
		FrictionFactorOpenHole:    hole.FrictionFactorOpenHole,
		LinearCapacityOpenHole:    hole.LinearCapacityOpenHole,
		VolumeExcess:              hole.VolumeExcess,
		DescriptionOpenHole:       hole.DescriptionOpenHole,
		TrippingInCasing:          hole.TrippingInCasing,
		TrippingOutCasing:         hole.TrippingOutCasing,
		RotatingOnBottomCasing:    hole.RotatingOnBottomCasing,
		SlideDrillingCasing:       hole.SlideDrillingCasing,
		BackReamingCasing:         hole.BackReamingCasing,
		RotatingOffBottomCasing:   hole.RotatingOffBottomCasing,
		TrippingInOpenHole:        hole.TrippingInOpenHole,
		TrippingOutOpenHole:       hole.TrippingOutOpenHole,
		RotatingOnBottomOpenHole:  hole.RotatingOnBottomOpenHole,
		SlideDrillingOpenHole:     hole.SlideDrillingOpenHole,
		BackReamingOpenHole:       hole.BackReamingOpenHole,
		RotatingOffBottomOpenHole: hole.RotatingOffBottomOpenHole,
	}
}
