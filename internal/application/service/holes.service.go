package service

import (
	"context"
	"fmt"

	"github.com/munaiplan/munaiplan-backend/internal/application/types/requests"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/domain/repository"
)

type holesService struct {
	commonRepo repository.CommonRepository
	repo       repository.HolesRepository
}

func NewHolesService(repo repository.HolesRepository, commonRepo repository.CommonRepository) *holesService {
	return &holesService{
		repo:       repo,
		commonRepo: commonRepo,
	}
}

func (s *holesService) GetHoles(ctx context.Context, input *requests.GetHolesRequest) ([]*entities.Hole, error) {
	if err := s.commonRepo.CheckIfCaseExists(ctx, input.CaseID); err != nil {
		return nil, err
	}

	return s.repo.GetHoles(ctx, input.CaseID)
}

func (s *holesService) GetHoleByID(ctx context.Context, input *requests.GetHoleByIDRequest) (*entities.Hole, error) {
	return s.repo.GetHoleByID(ctx, input.ID)
}

func (s *holesService) CreateHole(ctx context.Context, input *requests.CreateHoleRequest) error {
	if err := s.commonRepo.CheckIfCaseExists(ctx, input.CaseID); err != nil {
		return err
	}

	hole := &entities.Hole{
		CaseID:                    input.CaseID,
		MDTop:                     input.Body.MDTop,
		MDBase:                    input.Body.MDBase,
		Length:                    input.Body.Length,
		ShoeMD:                    input.Body.ShoeMD,
		OD:                        input.Body.OD,
		CaisingInternalDiameter:   input.Body.CaisingInternalDiameter,
		DriftInternalDiameter:     input.Body.DriftInternalDiameter,
		EffectiveHoleDiameter:     input.Body.EffectiveHoleDiameter,
		Weight:                    input.Body.Weight,
		Grade:                     input.Body.Grade,
		MinYieldStrength:          input.Body.MinYieldStrength,
		BurstRating:               input.Body.BurstRating,
		CollapseRating:            input.Body.CollapseRating,
		FrictionFactorCasing:      input.Body.FrictionFactorCasing,
		LinearCapacityCasing:      input.Body.LinearCapacityCasing,
		DescriptionCasing:         input.Body.DescriptionCasing,
		ManufacturerCasing:        input.Body.ManufacturerCasing,
		ModelCasing:               input.Body.ModelCasing,
		OpenHoleMDTop:             input.Body.OpenHoleMDTop,
		OpenHoleMDBase:            input.Body.OpenHoleMDBase,
		OpenHoleLength:            input.Body.OpenHoleLength,
		OpenHoleInternalDiameter:  input.Body.OpenHoleInternalDiameter,
		EffectiveDiameter:         input.Body.EffectiveDiameter,
		FrictionFactorOpenHole:    input.Body.FrictionFactorOpenHole,
		LinearCapacityOpenHole:    input.Body.LinearCapacityOpenHole,
		VolumeExcess:              input.Body.VolumeExcess,
		DescriptionOpenHole:       input.Body.DescriptionOpenHole,
		TrippingInCasing:          input.Body.TrippingInCasing,
		TrippingOutCasing:         input.Body.TrippingOutCasing,
		RotatingOnBottomCasing:    input.Body.RotatingOnBottomCasing,
		SlideDrillingCasing:       input.Body.SlideDrillingCasing,
		BackReamingCasing:         input.Body.BackReamingCasing,
		RotatingOffBottomCasing:   input.Body.RotatingOffBottomCasing,
		TrippingInOpenHole:        input.Body.TrippingInOpenHole,
		TrippingOutOpenHole:       input.Body.TrippingOutOpenHole,
		RotatingOnBottomOpenHole:  input.Body.RotatingOnBottomOpenHole,
		SlideDrillingOpenHole:     input.Body.SlideDrillingOpenHole,
		BackReamingOpenHole:       input.Body.BackReamingOpenHole,
		RotatingOffBottomOpenHole: input.Body.RotatingOffBottomOpenHole,
	}

	return s.repo.CreateHole(ctx, input.CaseID, hole)
}

func (s *holesService) UpdateHole(ctx context.Context, input *requests.UpdateHoleRequest) (*entities.Hole, error) {
	fmt.Println("service id is " + input.ID)
	hole := &entities.Hole{
		ID:                        input.ID,
		MDTop:                     input.Body.MDTop,
		MDBase:                    input.Body.MDBase,
		Length:                    input.Body.Length,
		ShoeMD:                    input.Body.ShoeMD,
		OD:                        input.Body.OD,
		CaisingInternalDiameter:   input.Body.CaisingInternalDiameter,
		DriftInternalDiameter:     input.Body.DriftInternalDiameter,
		EffectiveHoleDiameter:     input.Body.EffectiveHoleDiameter,
		Weight:                    input.Body.Weight,
		Grade:                     input.Body.Grade,
		MinYieldStrength:          input.Body.MinYieldStrength,
		BurstRating:               input.Body.BurstRating,
		CollapseRating:            input.Body.CollapseRating,
		FrictionFactorCasing:      input.Body.FrictionFactorCasing,
		LinearCapacityCasing:      input.Body.LinearCapacityCasing,
		DescriptionCasing:         input.Body.DescriptionCasing,
		ManufacturerCasing:        input.Body.ManufacturerCasing,
		ModelCasing:               input.Body.ModelCasing,
		OpenHoleMDTop:             input.Body.OpenHoleMDTop,
		OpenHoleMDBase:            input.Body.OpenHoleMDBase,
		OpenHoleLength:            input.Body.OpenHoleLength,
		OpenHoleInternalDiameter:  input.Body.OpenHoleInternalDiameter,
		EffectiveDiameter:         input.Body.EffectiveDiameter,
		FrictionFactorOpenHole:    input.Body.FrictionFactorOpenHole,
		LinearCapacityOpenHole:    input.Body.LinearCapacityOpenHole,
		VolumeExcess:              input.Body.VolumeExcess,
		DescriptionOpenHole:       input.Body.DescriptionOpenHole,
		TrippingInCasing:          input.Body.TrippingInCasing,
		TrippingOutCasing:         input.Body.TrippingOutCasing,
		RotatingOnBottomCasing:    input.Body.RotatingOnBottomCasing,
		SlideDrillingCasing:       input.Body.SlideDrillingCasing,
		BackReamingCasing:         input.Body.BackReamingCasing,
		RotatingOffBottomCasing:   input.Body.RotatingOffBottomCasing,
		TrippingInOpenHole:        input.Body.TrippingInOpenHole,
		TrippingOutOpenHole:       input.Body.TrippingOutOpenHole,
		RotatingOnBottomOpenHole:  input.Body.RotatingOnBottomOpenHole,
		SlideDrillingOpenHole:     input.Body.SlideDrillingOpenHole,
		BackReamingOpenHole:       input.Body.BackReamingOpenHole,
		RotatingOffBottomOpenHole: input.Body.RotatingOffBottomOpenHole,
	}

	return s.repo.UpdateHole(ctx, hole)
}

func (s *holesService) DeleteHole(ctx context.Context, input *requests.DeleteHoleRequest) error {
	return s.repo.DeleteHole(ctx, input.ID)
}
