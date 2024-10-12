package service

import (
	"context"
	"fmt"

	"github.com/munaiplan/munaiplan-backend/internal/application/types/requests"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/domain/repository"
)

type fractureGradientsService struct {
	commonRepo repository.CommonRepository
	repo       repository.FractureGradientsRepository
}

func NewFractureGradientsService(repo repository.FractureGradientsRepository, commonRepo repository.CommonRepository) *fractureGradientsService {
	return &fractureGradientsService{
		repo:       repo,
		commonRepo: commonRepo,
	}
}

func (s *fractureGradientsService) GetFractureGradients(ctx context.Context, input *requests.GetFractureGradientsRequest) ([]*entities.FractureGradient, error) {
	if err := s.commonRepo.CheckIfCaseExists(ctx, input.CaseID); err != nil {
		return nil, err
	}

	fmt.Println(input.CaseID)

	return s.repo.GetFractureGradients(ctx, input.CaseID)
}

func (s *fractureGradientsService) GetFractureGradientByID(ctx context.Context, input *requests.GetFractureGradientByIDRequest) (*entities.FractureGradient, error) {
	return s.repo.GetFractureGradientByID(ctx, input.ID)
}

func (s *fractureGradientsService) CreateFractureGradient(ctx context.Context, input *requests.CreateFractureGradientRequest) error {
	if err := s.commonRepo.CheckIfCaseExists(ctx, input.CaseID); err != nil {
		return err
	}

	fractureGradient := &entities.FractureGradient{
		TemperatureAtSurface: input.Body.TemperatureAtSurface,
		TemperatureAtWellTVD: input.Body.TemperatureAtWellTVD,
		TemperatureGradient:  input.Body.TemperatureGradient,
		WellTVD:              input.Body.WellTVD,
	}

	return s.repo.CreateFractureGradient(ctx, input.CaseID, fractureGradient)
}

func (s *fractureGradientsService) UpdateFractureGradient(ctx context.Context, input *requests.UpdateFractureGradientRequest) (*entities.FractureGradient, error) {
	fractureGradient := &entities.FractureGradient{
		ID:                   input.ID,
		TemperatureAtSurface: input.Body.TemperatureAtSurface,
		TemperatureAtWellTVD: input.Body.TemperatureAtWellTVD,
		TemperatureGradient:  input.Body.TemperatureGradient,
		WellTVD:              input.Body.WellTVD,
	}

	return s.repo.UpdateFractureGradient(ctx, fractureGradient)
}

func (s *fractureGradientsService) DeleteFractureGradient(ctx context.Context, input *requests.DeleteFractureGradientRequest) error {
	return s.repo.DeleteFractureGradient(ctx, input.ID)
}
