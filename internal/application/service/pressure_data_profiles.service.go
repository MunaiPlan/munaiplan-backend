package service

import (
	"context"
	"github.com/munaiplan/munaiplan-backend/internal/application/types/requests"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/domain/repository"
)

type pressureDataProfilesService struct {
	repo repository.PressureDataProfilesRepository
}

func NewPressureDataProfilesService(repo repository.PressureDataProfilesRepository) *pressureDataProfilesService {
	return &pressureDataProfilesService{repo: repo}
}

func (s *pressureDataProfilesService) CreatePressureDataProfile(ctx context.Context, input *requests.CreatePressureDataProfileRequest) error {
	profile := &entities.PressureDataProfile{
		TVD:      input.Body.TVD,
		Pressure: input.Body.Pressure,
		EMW:      input.Body.EMW,
	}
	return s.repo.CreatePressureDataProfile(ctx, input.CaseID, profile)
}

func (s *pressureDataProfilesService) GetPressureDataProfileByID(ctx context.Context, input *requests.GetPressureDataProfileByIDRequest) (*entities.PressureDataProfile, error) {
	return s.repo.GetPressureDataProfileByID(ctx, input.ID)
}

func (s *pressureDataProfilesService) GetPressureDataProfiles(ctx context.Context, input *requests.GetPressureDataProfilesRequest) ([]*entities.PressureDataProfile, error) {
	return s.repo.GetPressureDataProfiles(ctx, input.CaseID)
}

func (s *pressureDataProfilesService) UpdatePressureDataProfile(ctx context.Context, input *requests.UpdatePressureDataProfileRequest) error {
	profile := &entities.PressureDataProfile{
		ID:       input.ID,
		TVD:      input.Body.TVD,
		Pressure: input.Body.Pressure,
		EMW:      input.Body.EMW,
	}
	return s.repo.UpdatePressureDataProfile(ctx, profile)
}

func (s *pressureDataProfilesService) DeletePressureDataProfile(ctx context.Context, input *requests.DeletePressureDataProfileRequest) error {
	return s.repo.DeletePressureDataProfile(ctx, input.ID)
}
