package service

import (
	"context"

	"github.com/munaiplan/munaiplan-backend/internal/application/dto/requests"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/domain/repository"
)

type wellboresService struct {
	commonRepo repository.CommonRepository
	repo       repository.WellboresRepository
}

func NewWellboresService(repo repository.WellboresRepository, commonRepo repository.CommonRepository) *wellboresService {
	return &wellboresService{
		repo:       repo,
		commonRepo: commonRepo,
	}
}

func (s *wellboresService) GetWellbores(ctx context.Context, input *requests.GetWellboresRequest) ([]*entities.Wellbore, error) {
	if err := s.commonRepo.CheckIfWellExists(ctx, input.WellID); err != nil {
		return nil, err
	}

	return s.repo.GetWellbores(ctx, input.WellID)
}

func (s *wellboresService) GetWellboreByID(ctx context.Context, input *requests.GetWellboreByIDRequest) (*entities.Wellbore, error) {
	return s.repo.GetWellboreByID(ctx, input.ID)
}

func (s *wellboresService) CreateWellbore(ctx context.Context, input *requests.CreateWellboreRequest) error {
	if err := s.commonRepo.CheckIfWellExists(ctx, input.WellID); err != nil {
		return err
	}

	wellbore := &entities.Wellbore{
		Name:                           input.Body.Name,
		BottomHoleLocation:             input.Body.BottomHoleLocation,
		WellboreDepth:                  input.Body.WellboreDepth,
		AverageHookLoad:                input.Body.AverageHookLoad,
		RiserPressure:                  input.Body.RiserPressure,
		AverageInletFlow:               input.Body.AverageInletFlow,
		AverageColumnRotationFrequency: input.Body.AverageColumnRotationFrequency,
		MaximumColumnRotationFrequency: input.Body.MaximumColumnRotationFrequency,
		AverageWeightOnBit:             input.Body.AverageWeightOnBit,
		MaximumWeightOnBit:             input.Body.MaximumWeightOnBit,
		AverageTorque:                  input.Body.AverageTorque,
		MaximumTorque:                  input.Body.MaximumTorque,
		DownStaticFriction:             input.Body.DownStaticFriction,
		DepthInterval:                  input.Body.DepthInterval,
	}

	return s.repo.CreateWellbore(ctx, input.WellID, wellbore)
}

func (s *wellboresService) UpdateWellbore(ctx context.Context, input *requests.UpdateWellboreRequest) (*entities.Wellbore, error) {
	wellbore := &entities.Wellbore{
		ID:                             input.ID,
		Name:                           input.Body.Name,
		BottomHoleLocation:             input.Body.BottomHoleLocation,
		WellboreDepth:                  input.Body.WellboreDepth,
		AverageHookLoad:                input.Body.AverageHookLoad,
		RiserPressure:                  input.Body.RiserPressure,
		AverageInletFlow:               input.Body.AverageInletFlow,
		AverageColumnRotationFrequency: input.Body.AverageColumnRotationFrequency,
		MaximumColumnRotationFrequency: input.Body.MaximumColumnRotationFrequency,
		AverageWeightOnBit:             input.Body.AverageWeightOnBit,
		MaximumWeightOnBit:             input.Body.MaximumWeightOnBit,
		AverageTorque:                  input.Body.AverageTorque,
		MaximumTorque:                  input.Body.MaximumTorque,
		DownStaticFriction:             input.Body.DownStaticFriction,
		DepthInterval:                  input.Body.DepthInterval,
	}

	return s.repo.UpdateWellbore(ctx, wellbore)
}

func (s *wellboresService) DeleteWellbore(ctx context.Context, input *requests.DeleteWellboreRequest) error {
	return s.repo.DeleteWellbore(ctx, input.ID)
}
