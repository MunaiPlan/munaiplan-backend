package service

import (
	"context"

	"github.com/munaiplan/munaiplan-backend/internal/application/dto/requests"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/domain/repository"
)

type wellsService struct {
	commonRepo repository.CommonRepository
	repo       repository.WellsRepository
}

func NewWellsService(repo repository.WellsRepository, commonRepo repository.CommonRepository) *wellsService {
	return &wellsService{
		repo:       repo,
		commonRepo: commonRepo,
	}
}

func (s *wellsService) GetWells(ctx context.Context, input *requests.GetWellsRequest) ([]*entities.Well, error) {
	if err := s.commonRepo.CheckIfSiteExists(ctx, input.SiteID); err != nil {
		return nil, err
	}

	return s.repo.GetWells(ctx, input.SiteID)
}

func (s *wellsService) GetWellByID(ctx context.Context, input *requests.GetWellByIDRequest) (*entities.Well, error) {
	return s.repo.GetWellByID(ctx, input.ID)
}

func (s *wellsService) CreateWell(ctx context.Context, input *requests.CreateWellRequest) error {
	if err := s.commonRepo.CheckIfSiteExists(ctx, input.SiteID); err != nil {
		return err
	}

	well := &entities.Well{
		Name:                    input.Body.Name,
		Description:             input.Body.Description,
		Location:                input.Body.Location,
		UniversalWellIdentifier: input.Body.UniversalWellIdentifier,
		Type:                    input.Body.Type,
		WellNumber:              input.Body.WellNumber,
		WorkingGroup:            input.Body.WorkingGroup,
		ActiveWellUnit:          input.Body.ActiveWellUnit,
	}

	return s.repo.CreateWell(ctx, input.SiteID, well)
}

func (s *wellsService) UpdateWell(ctx context.Context, input *requests.UpdateWellRequest) (*entities.Well, error) {
	well := &entities.Well{
		ID:                      input.ID,
		Name:                    input.Body.Name,
		Description:             input.Body.Description,
		Location:                input.Body.Location,
		UniversalWellIdentifier: input.Body.UniversalWellIdentifier,
		Type:                    input.Body.Type,
		WellNumber:              input.Body.WellNumber,
		WorkingGroup:            input.Body.WorkingGroup,
		ActiveWellUnit:          input.Body.ActiveWellUnit,
	}

	return s.repo.UpdateWell(ctx, well)
}

func (s *wellsService) DeleteWell(ctx context.Context, input *requests.DeleteWellRequest) error {
	return s.repo.DeleteWell(ctx, input.ID)
}
