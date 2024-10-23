package service

import (
	"context"

	types "github.com/munaiplan/munaiplan-backend/internal/application/types/errors"
	"github.com/munaiplan/munaiplan-backend/internal/application/types/requests"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/domain/repository"
)

type stringsService struct {
	commonRepo repository.CommonRepository
	repo       repository.StringsRepository
}

func NewStringsService(repo repository.StringsRepository, commonRepo repository.CommonRepository) *stringsService {
	return &stringsService{
		repo:       repo,
		commonRepo: commonRepo,
	}
}

func (s *stringsService) GetStrings(ctx context.Context, input *requests.GetStringsRequest) ([]*entities.String, error) {
	if err := s.commonRepo.CheckIfCaseExists(ctx, input.CaseID); err != nil {
		return nil, err
	}

	return s.repo.GetStrings(ctx, input.CaseID)
}

func (s *stringsService) GetStringByID(ctx context.Context, input *requests.GetStringByIDRequest) (*entities.String, error) {
	return s.repo.GetStringByID(ctx, input.ID)
}

func (s *stringsService) CreateString(ctx context.Context, input *requests.CreateStringRequest) error {
	if err := s.commonRepo.CheckIfCaseExists(ctx, input.CaseID); err != nil {
		return err
	}

	if exists, err := s.commonRepo.CheckIfStringExists(ctx, input.CaseID); err != nil {
		return err
	} else if exists {
		return types.ErrAlreadyExists
	}

	newString := s.CreateStringRequestToEntity(&input.Body)
	return s.repo.CreateString(ctx, input.CaseID, newString)
}

func (s *stringsService) UpdateString(ctx context.Context, input *requests.UpdateStringRequest) (*entities.String, error) {
	updatedString := s.UpdateStringRequestToEntity(&input.Body)
	updatedString.ID = input.ID
	return s.repo.UpdateString(ctx, updatedString)
}

func (s *stringsService) DeleteString(ctx context.Context, input *requests.DeleteStringRequest) error {
	return s.repo.DeleteString(ctx, input.ID)
}

func (s *stringsService) CreateStringRequestToEntity(input *requests.CreateStringRequestBody) *entities.String {
	sections := make([]*entities.Section, len(input.Sections))
	for i, section := range input.Sections {
		sections[i] = &entities.Section{
			Description:         section.Description,
			Manufacturer:        section.Manufacturer,
			Type:                section.Type,
			BodyMD:              section.BodyMD,
			BodyLength:          section.BodyLength,
			BodyOD:              section.BodyOD,
			BodyID:              section.BodyID,
			AvgJointLength:      section.AvgJointLength,
			StabilizerLength:    section.StabilizerLength,
			StabilizerOD:        section.StabilizerOD,
			StabilizerID:        section.StabilizerID,
			Weight:              section.Weight,
			Material:            section.Material,
			Grade:               section.Grade,
			Class:               section.Class,
			FrictionCoefficient: section.FrictionCoefficient,
			MinYieldStrength:    section.MinYieldStrength,
		}
	}

	return &entities.String{
		Name:     input.Name,
		Depth:    input.Depth,
		Sections: sections,
	}
}

func (s *stringsService) UpdateStringRequestToEntity(input *requests.UpdateStringRequestBody) *entities.String {
	sections := make([]*entities.Section, len(input.Sections))
	for i, section := range input.Sections {
		sections[i] = &entities.Section{
			ID:                  section.ID,
			Description:         section.Description,
			Manufacturer:        section.Manufacturer,
			Type:                section.Type,
			BodyMD:              section.BodyMD,
			BodyLength:          section.BodyLength,
			BodyOD:              section.BodyOD,
			BodyID:              section.BodyID,
			AvgJointLength:      section.AvgJointLength,
			StabilizerLength:    section.StabilizerLength,
			StabilizerOD:        section.StabilizerOD,
			StabilizerID:        section.StabilizerID,
			Weight:              section.Weight,
			Material:            section.Material,
			Grade:               section.Grade,
			Class:               section.Class,
			FrictionCoefficient: section.FrictionCoefficient,
			MinYieldStrength:    section.MinYieldStrength,
		}
	}

	return &entities.String{
		Name:     input.Name,
		Depth:    input.Depth,
		Sections: sections,
	}
}
