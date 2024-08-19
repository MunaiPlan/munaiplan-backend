package service

import (
	"context"
	"fmt"

	"github.com/munaiplan/munaiplan-backend/internal/application/dto/requests"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/domain/repository"
)

type designsService struct {
	commonRepo repository.CommonRepository
	repo       repository.DesignsRepository
}

func NewDesignsService(repo repository.DesignsRepository, commonRepo repository.CommonRepository) *designsService {
	return &designsService{
		repo:       repo,
		commonRepo: commonRepo,
	}
}

func (s *designsService) GetDesigns(ctx context.Context, input *requests.GetDesignsRequest) ([]*entities.Design, error) {
	if err := s.commonRepo.CheckIfWellboreExists(ctx, input.WellboreID); err != nil {
		return nil, err
	}

	fmt.Println(input.WellboreID)

	return s.repo.GetDesigns(ctx, input.WellboreID)
}

func (s *designsService) GetDesignByID(ctx context.Context, input *requests.GetDesignByIDRequest) (*entities.Design, error) {
	return s.repo.GetDesignByID(ctx, input.ID)
}

func (s *designsService) CreateDesign(ctx context.Context, input *requests.CreateDesignRequest) error {
	if err := s.commonRepo.CheckIfWellboreExists(ctx, input.WellboreID); err != nil {
		return err
	}

	design := &entities.Design{
		PlanName:   input.Body.PlanName,
		Stage:      input.Body.Stage,
		Version:    input.Body.Version,
		ActualDate: input.Body.ActualDate,
	}

	return s.repo.CreateDesign(ctx, input.WellboreID, design)
}

func (s *designsService) UpdateDesign(ctx context.Context, input *requests.UpdateDesignRequest) (*entities.Design, error) {
	design := &entities.Design{
		ID:         input.ID,
		PlanName:   input.Body.PlanName,
		Stage:      input.Body.Stage,
		Version:    input.Body.Version,
		ActualDate: input.Body.ActualDate,
	}

	return s.repo.UpdateDesign(ctx, design)
}

func (s *designsService) DeleteDesign(ctx context.Context, input *requests.DeleteDesignRequest) error {
	return s.repo.DeleteDesign(ctx, input.ID)
}
