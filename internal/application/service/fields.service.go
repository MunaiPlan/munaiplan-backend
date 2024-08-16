package service

import (
	"context"

	"github.com/munaiplan/munaiplan-backend/internal/application/dto/requests"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/domain/repository"
)

type fieldsService struct {
	commonRepo repository.CommonRepository
	repo       repository.FieldsRepository
}

func NewFieldsService(repo repository.FieldsRepository, commonRepo repository.CommonRepository) *fieldsService {
	return &fieldsService{
		repo:       repo,
		commonRepo: commonRepo,
	}
}

func (s *fieldsService) GetFields(ctx context.Context, input *requests.GetFieldsRequest) ([]*entities.Field, error) {
	if err := s.commonRepo.CheckIfCompanyExists(ctx, input.CompanyID); err != nil {
		return nil, err
	}

	return s.repo.GetFields(ctx, input.CompanyID)
}

func (s *fieldsService) GetFieldByID(ctx context.Context, input *requests.GetFieldByIDRequest) (*entities.Field, error) {
	if err := s.commonRepo.CheckIfCompanyExists(ctx, input.CompanyID); err != nil {
		return nil, err
	}

	return s.repo.GetFieldByID(ctx, input.ID, input.CompanyID)
}

func (s *fieldsService) CreateField(ctx context.Context, input *requests.CreateFieldRequest) error {
	if err := s.commonRepo.CheckIfCompanyExists(ctx, input.CompanyID); err != nil {
		return err
	}

	field := &entities.Field{
		Name:            input.Body.Name,
		Description:     input.Body.Description,
		ReductionLevel:  input.Body.ReductionLevel,
		ActiveFieldUnit: input.Body.ActiveFieldUnit,
	}

	return s.repo.CreateField(ctx, input.CompanyID, field)
}

func (s *fieldsService) UpdateField(ctx context.Context, input *requests.UpdateFieldRequest) (*entities.Field, error) {
	if err := s.commonRepo.CheckIfCompanyExists(ctx, input.CompanyID); err != nil {
		return nil, err
	}

	field := &entities.Field{
		ID:              input.Body.ID,
		Name:            input.Body.Name,
		Description:     input.Body.Description,
		ReductionLevel:  input.Body.ReductionLevel,
		ActiveFieldUnit: input.Body.ActiveFieldUnit,
	}

	return s.repo.UpdateField(ctx, input.CompanyID, field)
}

func (s *fieldsService) DeleteField(ctx context.Context, input *requests.DeleteFieldRequest) error {
	if err := s.commonRepo.CheckIfCompanyExists(ctx, input.CompanyID); err != nil {
		return err
	}

	return s.repo.DeleteField(ctx, input.CompanyID, input.ID)
}
