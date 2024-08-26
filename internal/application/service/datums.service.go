package service

import (
	"context"

	"github.com/munaiplan/munaiplan-backend/internal/application/types/requests"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/domain/repository"
)

type datumsService struct {
	commonRepo repository.CommonRepository
	repo       repository.DatumsRepository
}

func NewDatumsService(repo repository.DatumsRepository, commonRepo repository.CommonRepository) *datumsService {
	return &datumsService{
		repo:       repo,
		commonRepo: commonRepo,
	}
}

func (s *datumsService) GetDatumsByCaseID(ctx context.Context, input *requests.GetDatumsByCaseIDRequest) ([]*entities.Datum, error) {
	if err := s.commonRepo.CheckIfCaseExists(ctx, input.CaseID); err != nil {
		return nil, err
	}

	return s.repo.GetDatumsByCaseID(ctx, input.CaseID)
}

func (s *datumsService) GetDatumByID(ctx context.Context, input *requests.GetDatumByIDRequest) (*entities.Datum, error) {
	return s.repo.GetDatumByID(ctx, input.ID)
}

func (s *datumsService) CreateDatum(ctx context.Context, input *requests.CreateDatumRequest) error {
	if err := s.commonRepo.CheckIfCaseExists(ctx, input.CaseID); err != nil {
		return err
	}

	datum := &entities.Datum{
		Name:              input.Body.Name,
		SystemDescription: input.Body.SystemDescription,
		SystemElevation:   input.Body.SystemElevation,
		DatumDescription:  input.Body.DatumDescription,
		WellheadElevation: input.Body.WellheadElevation,
		DatumElevation:    input.Body.DatumElevation,
		AirGap:            input.Body.AirGap,
		GroundElevation:   input.Body.GroundElevation,
		Type:              input.Body.Type,
	}

	return s.repo.CreateDatum(ctx, input.CaseID, datum)
}

func (s *datumsService) UpdateDatum(ctx context.Context, input *requests.UpdateDatumRequest) (*entities.Datum, error) {
	datum := &entities.Datum{
		ID:                input.ID,
		Name:              input.Body.Name,
		SystemDescription: input.Body.SystemDescription,
		SystemElevation:   input.Body.SystemElevation,
		DatumDescription:  input.Body.DatumDescription,
		WellheadElevation: input.Body.WellheadElevation,
		DatumElevation:    input.Body.DatumElevation,
		AirGap:            input.Body.AirGap,
		GroundElevation:   input.Body.GroundElevation,
		Type:              input.Body.Type,
	}

	return s.repo.UpdateDatum(ctx, datum)
}

func (s *datumsService) DeleteDatum(ctx context.Context, input *requests.DeleteDatumRequest) error {
	return s.repo.DeleteDatum(ctx, input.ID)
}
