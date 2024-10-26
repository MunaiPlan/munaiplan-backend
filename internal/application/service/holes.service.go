package service

import (
	"context"
	"fmt"

	"github.com/munaiplan/munaiplan-backend/internal/application/types/errors"
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

	if exists, err := s.commonRepo.CheckIfHoleExists(ctx, input.CaseID); err != nil {
		return err
	} else if exists {
		return types.ErrAlreadyExists
	}

	fmt.Println("len of input.Body.Caisings at service", len(input.Body.Caisings))

	hole := s.CreateHoleRequestToEntity(&input.Body)
	return s.repo.CreateHole(ctx, input.CaseID, hole)
}

func (s *holesService) UpdateHole(ctx context.Context, input *requests.UpdateHoleRequest) (*entities.Hole, error) {
	hole := s.UpdateHoleRequestToEntity(&input.Body)
	hole.ID = input.ID
	return s.repo.UpdateHole(ctx, hole)
}

func (s *holesService) DeleteHole(ctx context.Context, input *requests.DeleteHoleRequest) error {
	return s.repo.DeleteHole(ctx, input.ID)
}

func (s *holesService) CreateHoleRequestToEntity(input *requests.CreateHoleRequestBody) *entities.Hole {
	caisings := make([]*entities.Caising, len(input.Caisings))
	for i, caising := range input.Caisings {
		caisings[i] = &entities.Caising{
			MDTop:                 caising.MDTop,
			MDBase:                caising.MDBase,
			Length:                caising.Length,
			ShoeMD:                caising.ShoeMD,
			OD:                    caising.OD,
			VD:                    caising.VD,
			DriftID:               caising.DriftID,
			EffectiveHoleDiameter: caising.EffectiveHoleDiameter,
			Weight:                caising.Weight,
			Grade:                 caising.Grade,
			MinYieldStrength:      caising.MinYieldStrength,
			BurstRating:           caising.BurstRating,
			CollapseRating:        caising.CollapseRating,
			FrictionFactorCaising: caising.FrictionFactorCaising,
			LinearCapacityCaising: caising.LinearCapacityCaising,
			DescriptionCaising:    caising.DescriptionCaising,
			ManufacturerCaising:   caising.ManufacturerCaising,
			ModelCaising:          caising.ModelCaising,
		}
	}

	hole := &entities.Hole{
		OpenHoleMDTop:             input.OpenHoleMDTop,
		OpenHoleMDBase:            input.OpenHoleMDBase,
		OpenHoleLength:            input.OpenHoleLength,
		OpenHoleVD:                input.OpenHoleVD,
		EffectiveDiameter:         input.EffectiveDiameter,
		FrictionFactorOpenHole:    input.FrictionFactorOpenHole,
		LinearCapacityOpenHole:    input.LinearCapacityOpenHole,
		VolumeExcess:              input.VolumeExcess,
		DescriptionOpenHole:       input.DescriptionOpenHole,
		TrippingInCasing:          input.TrippingInCasing,
		TrippingOutCasing:         input.TrippingOutCasing,
		RotatingOnBottomCasing:    input.RotatingOnBottomCasing,
		SlideDrillingCasing:       input.SlideDrillingCasing,
		BackReamingCasing:         input.BackReamingCasing,
		RotatingOffBottomCasing:   input.RotatingOffBottomCasing,
		TrippingInOpenHole:        input.TrippingInOpenHole,
		TrippingOutOpenHole:       input.TrippingOutOpenHole,
		RotatingOnBottomOpenHole:  input.RotatingOnBottomOpenHole,
		SlideDrillingOpenHole:     input.SlideDrillingOpenHole,
		BackReamingOpenHole:       input.BackReamingOpenHole,
		RotatingOffBottomOpenHole: input.RotatingOffBottomOpenHole,
		Caisings:                  caisings,
	}

	return hole
}

func (s *holesService) UpdateHoleRequestToEntity(input *requests.UpdateHoleRequestBody) *entities.Hole {
	caisings := make([]*entities.Caising, len(input.Caisings))
	for i, caising := range input.Caisings {
		caisings[i] = &entities.Caising{
			ID:                    caising.ID,
			MDTop:                 caising.MDTop,
			MDBase:                caising.MDBase,
			Length:                caising.Length,
			ShoeMD:                caising.ShoeMD,
			OD:                    caising.OD,
			VD:                    caising.VD,
			DriftID:               caising.DriftID,
			EffectiveHoleDiameter: caising.EffectiveHoleDiameter,
			Weight:                caising.Weight,
			Grade:                 caising.Grade,
			MinYieldStrength:      caising.MinYieldStrength,
			BurstRating:           caising.BurstRating,
			CollapseRating:        caising.CollapseRating,
			FrictionFactorCaising: caising.FrictionFactorCaising,
			LinearCapacityCaising: caising.LinearCapacityCaising,
			DescriptionCaising:    caising.DescriptionCaising,
			ManufacturerCaising:   caising.ManufacturerCaising,
			ModelCaising:          caising.ModelCaising,
		}
	}

	hole := &entities.Hole{
		OpenHoleMDTop:             input.OpenHoleMDTop,
		OpenHoleMDBase:            input.OpenHoleMDBase,
		OpenHoleLength:            input.OpenHoleLength,
		OpenHoleVD:                input.OpenHoleVD,
		EffectiveDiameter:         input.EffectiveDiameter,
		FrictionFactorOpenHole:    input.FrictionFactorOpenHole,
		LinearCapacityOpenHole:    input.LinearCapacityOpenHole,
		VolumeExcess:              input.VolumeExcess,
		DescriptionOpenHole:       input.DescriptionOpenHole,
		TrippingInCasing:          input.TrippingInCasing,
		TrippingOutCasing:         input.TrippingOutCasing,
		RotatingOnBottomCasing:    input.RotatingOnBottomCasing,
		SlideDrillingCasing:       input.SlideDrillingCasing,
		BackReamingCasing:         input.BackReamingCasing,
		RotatingOffBottomCasing:   input.RotatingOffBottomCasing,
		TrippingInOpenHole:        input.TrippingInOpenHole,
		TrippingOutOpenHole:       input.TrippingOutOpenHole,
		RotatingOnBottomOpenHole:  input.RotatingOnBottomOpenHole,
		SlideDrillingOpenHole:     input.SlideDrillingOpenHole,
		BackReamingOpenHole:       input.BackReamingOpenHole,
		RotatingOffBottomOpenHole: input.RotatingOffBottomOpenHole,
		Caisings:                  caisings,
	}

	return hole
}
