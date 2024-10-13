package service

import (
	"context"

	types "github.com/munaiplan/munaiplan-backend/internal/application/types/errors"
	"github.com/munaiplan/munaiplan-backend/internal/application/types/requests"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/domain/repository"
)

type rigsService struct {
	commonRepo repository.CommonRepository
	repo       repository.RigsRepository
}

func NewRigsService(repo repository.RigsRepository, commonRepo repository.CommonRepository) *rigsService {
	return &rigsService{
		repo:       repo,
		commonRepo: commonRepo,
	}
}

func (s *rigsService) GetRigs(ctx context.Context, input *requests.GetRigsRequest) ([]*entities.Rig, error) {
	if err := s.commonRepo.CheckIfCaseExists(ctx, input.CaseID); err != nil {
		return nil, err
	}
	return s.repo.GetRigs(ctx, input.CaseID)
}

func (s *rigsService) GetRigByID(ctx context.Context, input *requests.GetRigByIDRequest) (*entities.Rig, error) {
	return s.repo.GetRigByID(ctx, input.ID)
}

func (s *rigsService) CreateRig(ctx context.Context, input *requests.CreateRigRequest) error {
	if err := s.commonRepo.CheckIfCaseExists(ctx, input.CaseID); err != nil {
		return err
	}

	if exists, err := s.commonRepo.CheckIfRigExists(ctx, input.CaseID); err != nil {
		return err
	} else if exists {
		return types.ErrAlreadyExists
	}

	rig := s.CreateRigRequestToEntity(&input.Body)
	return s.repo.CreateRig(ctx, input.CaseID, rig)
}

func (s *rigsService) UpdateRig(ctx context.Context, input *requests.UpdateRigRequest) (*entities.Rig, error) {
	rig := s.UpdateRigRequestToEntity(&input.Body)
	rig.ID = input.ID
	return s.repo.UpdateRig(ctx, rig)
}

func (s *rigsService) DeleteRig(ctx context.Context, input *requests.DeleteRigRequest) error {
	return s.repo.DeleteRig(ctx, input.ID)
}

// CreateRigRequestToEntity converts a CreateRigRequestBody to a Rig entity.
func (s *rigsService) CreateRigRequestToEntity(input *requests.CreateRigRequestBody) *entities.Rig {
	return &entities.Rig{
		BlockRating:                     input.BlockRating,
		TorqueRating:                    input.TorqueRating,
		RatedWorkingPressure:            input.RatedWorkingPressure,
		BopPressureRating:               input.BopPressureRating,
		SurfacePressureLoss:             input.SurfacePressureLoss,
		StandpipeLength:                 input.StandpipeLength,
		StandpipeInternalDiameter:       input.StandpipeInternalDiameter,
		HoseLength:                      input.HoseLength,
		HoseInternalDiameter:            input.HoseInternalDiameter,
		SwivelLength:                    input.SwivelLength,
		SwivelInternalDiameter:          input.SwivelInternalDiameter,
		KellyLength:                     input.KellyLength,
		KellyInternalDiameter:           input.KellyInternalDiameter,
		PumpDischargeLineLength:         input.PumpDischargeLineLength,
		PumpDischargeLineInternalDiameter: input.PumpDischargeLineInternalDiameter,
		TopDriveStackupLength:           input.TopDriveStackupLength,
		TopDriveStackupInternalDiameter: input.TopDriveStackupInternalDiameter,
	}
}

// UpdateRigRequestToEntity converts an UpdateRigRequestBody to a Rig entity.
func (s *rigsService) UpdateRigRequestToEntity(input *requests.UpdateRigRequestBody) *entities.Rig {
	return &entities.Rig{
		BlockRating:                     input.BlockRating,
		TorqueRating:                    input.TorqueRating,
		RatedWorkingPressure:            input.RatedWorkingPressure,
		BopPressureRating:               input.BopPressureRating,
		SurfacePressureLoss:             input.SurfacePressureLoss,
		StandpipeLength:                 input.StandpipeLength,
		StandpipeInternalDiameter:       input.StandpipeInternalDiameter,
		HoseLength:                      input.HoseLength,
		HoseInternalDiameter:            input.HoseInternalDiameter,
		SwivelLength:                    input.SwivelLength,
		SwivelInternalDiameter:          input.SwivelInternalDiameter,
		KellyLength:                     input.KellyLength,
		KellyInternalDiameter:           input.KellyInternalDiameter,
		PumpDischargeLineLength:         input.PumpDischargeLineLength,
		PumpDischargeLineInternalDiameter: input.PumpDischargeLineInternalDiameter,
		TopDriveStackupLength:           input.TopDriveStackupLength,
		TopDriveStackupInternalDiameter: input.TopDriveStackupInternalDiameter,
	}
}
