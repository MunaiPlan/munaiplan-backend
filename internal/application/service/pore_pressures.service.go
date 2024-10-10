package service

import (
	"context"

	"github.com/munaiplan/munaiplan-backend/internal/application/types/requests"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/domain/repository"
)

type porePressuresService struct {
	commonRepo       repository.CommonRepository
	porePressuresRepo repository.PorePressuresRepository
}

func NewPorePressuresService(porePressureRepo repository.PorePressuresRepository, commonRepo repository.CommonRepository) *porePressuresService {
	return &porePressuresService{
		porePressuresRepo: porePressureRepo,
		commonRepo:       commonRepo,
	}
}

func (s *porePressuresService) CreatePorePressure(ctx context.Context, input *requests.CreatePorePressureRequest) error {
	if err := s.commonRepo.CheckIfCaseExists(ctx, input.CaseID); err != nil {
		return err
	}

	porePressure := &entities.PorePressure{
		TVD:      input.TVD,
		Pressure: input.Pressure,
		EMW:      input.EMW,
	}

	return s.porePressuresRepo.CreatePorePressure(ctx, input.CaseID, porePressure)
}

func (s *porePressuresService) GetPorePressureByID(ctx context.Context, input *requests.GetPorePressureByIDRequest) (*entities.PorePressure, error) {
	return s.porePressuresRepo.GetPorePressureByID(ctx, input.ID)
}

func (s *porePressuresService) GetPorePressures(ctx context.Context, input *requests.GetPorePressuresRequest) ([]*entities.PorePressure, error) {
	if err := s.commonRepo.CheckIfCaseExists(ctx, input.CaseID); err != nil {
		return nil, err
	}

	return s.porePressuresRepo.GetPorePressures(ctx, input.CaseID)
}

func (s *porePressuresService) UpdatePorePressure(ctx context.Context, input *requests.UpdatePorePressureRequest) (*entities.PorePressure, error) {
	porePressure := &entities.PorePressure{
		ID:       input.ID,
		TVD:      input.TVD,
		Pressure: input.Pressure,
		EMW:      input.EMW,
	}

	return s.porePressuresRepo.UpdatePorePressure(ctx, porePressure)
}

func (s *porePressuresService) DeletePorePressure(ctx context.Context, input *requests.DeletePorePressureRequest) error {
	return s.porePressuresRepo.DeletePorePressure(ctx, input.ID)
}
