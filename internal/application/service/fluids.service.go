package service

import (
	"context"

	"github.com/munaiplan/munaiplan-backend/internal/application/types/requests"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/domain/repository"
)

type fluidsService struct {
	commonRepo repository.CommonRepository
	repo       repository.FluidsRepository
}

// NewFluidsService creates a new instance of the fluids service.
func NewFluidsService(repo repository.FluidsRepository, commonRepo repository.CommonRepository) *fluidsService {
	return &fluidsService{
		repo:       repo,
		commonRepo: commonRepo,
	}
}

// GetFluids retrieves all fluids associated with a specific case.
func (s *fluidsService) GetFluids(ctx context.Context, input *requests.GetFluidsRequest) ([]*entities.Fluid, error) {
	if err := s.commonRepo.CheckIfCaseExists(ctx, input.CaseID); err != nil {
		return nil, err
	}

	return s.repo.GetFluids(ctx, input.CaseID)
}

// GetFluidByID retrieves a fluid by its ID.
func (s *fluidsService) GetFluidByID(ctx context.Context, input *requests.GetFluidByIDRequest) (*entities.Fluid, error) {
	return s.repo.GetFluidByID(ctx, input.ID)
}

// CreateFluid creates a new fluid within a specific case.
func (s *fluidsService) CreateFluid(ctx context.Context, input *requests.CreateFluidRequest) error {
	if err := s.commonRepo.CheckIfCaseExists(ctx, input.CaseID); err != nil {
		return err
	}

	fluid := s.CreateFluidRequestToEntity(&input.Body)
	return s.repo.CreateFluid(ctx, input.CaseID, fluid)
}

// UpdateFluid updates an existing fluid.
func (s *fluidsService) UpdateFluid(ctx context.Context, input *requests.UpdateFluidRequest) (*entities.Fluid, error) {
	fluid := s.UpdateFluidRequestToEntity(&input.Body)
	fluid.ID = input.ID
	return s.repo.UpdateFluid(ctx, fluid)
}

// DeleteFluid deletes a fluid by its ID.
func (s *fluidsService) DeleteFluid(ctx context.Context, input *requests.DeleteFluidRequest) error {
	return s.repo.DeleteFluid(ctx, input.ID)
}

// CreateFluidRequestToEntity converts a create request to a fluid entity.
func (s *fluidsService) CreateFluidRequestToEntity(input *requests.CreateFluidRequestBody) *entities.Fluid {
	return &entities.Fluid{
		Name:          input.Name,
		Description:   input.Description,
		Density:       input.Density,
		FluidBaseType: &entities.FluidType{ID: input.FluidBaseTypeID},
		BaseFluid:     &entities.FluidType{ID: input.BaseFluidID},
	}
}

// UpdateFluidRequestToEntity converts an update request to a fluid entity.
func (s *fluidsService) UpdateFluidRequestToEntity(input *requests.UpdateFluidRequestBody) *entities.Fluid {
	return &entities.Fluid{
		ID:            input.ID,
		Name:          input.Name,
		Description:   input.Description,
		Density:       input.Density,
		FluidBaseType: &entities.FluidType{ID: input.FluidBaseTypeID},
		BaseFluid:     &entities.FluidType{ID: input.BaseFluidID},
	}
}
