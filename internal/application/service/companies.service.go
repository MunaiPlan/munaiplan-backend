package service

import (
	"context"

	"github.com/munaiplan/munaiplan-backend/internal/application/dto/requests"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/domain/repository"
)

type companiesService struct {
	commonRepo repository.CommonRepository
	repo repository.CompaniesRepository
}

func NewCompaniesService(repo repository.CompaniesRepository, commonRepo repository.CommonRepository) *companiesService {
	return &companiesService{
		repo: repo,
		commonRepo: commonRepo,
	}
}

func (s *companiesService) GetCompanies(ctx context.Context, input *requests.GetCompaniesRequest) ([]*entities.Company, error) {
	if err := s.commonRepo.CheckIfOrganizationExists(ctx, input.OrganizationID); err != nil {
		return nil, err
	}

	return s.repo.GetCompanies(ctx, input.OrganizationID)
}

func (s *companiesService) GetCompanyByName(ctx context.Context, input *requests.GetCompanyByNameRequest) (*entities.Company, error) { 
	if err := s.commonRepo.CheckIfOrganizationExists(ctx, input.OrganizationID); err != nil {
		return nil, err
	}

	return s.repo.GetCompanyByName(ctx, input.Name, input.OrganizationID)
}

func (s *companiesService) CreateCompany(ctx context.Context, input *requests.CreateCompanyRequest) error {
	if err := s.commonRepo.CheckIfOrganizationExists(ctx, input.OrganizationID); err != nil {
		return err
	}

	company := &entities.Company{
		Name: input.Body.Name,
		Division: input.Body.Division,
		Group: input.Body.Group,
		Representative: input.Body.Representative,
		Address: input.Body.Address,
		Phone: input.Body.Phone,
	}
	return s.repo.CreateCompany(ctx, input.OrganizationID, company)
}

func (s *companiesService) UpdateCompany(ctx context.Context, input *requests.UpdateCompanyRequest) (*entities.Company, error) {
	if err := s.commonRepo.CheckIfOrganizationExists(ctx, input.OrganizationID); err != nil {
		return nil, err
	}

	company := &entities.Company{
		ID: input.Body.ID,
		Name: input.Body.Name,
		Division: input.Body.Division,
		Group: input.Body.Group,
		Representative: input.Body.Representative,
		Address: input.Body.Address,
		Phone: input.Body.Phone,
	}
	return s.repo.UpdateCompany(ctx, input.OrganizationID, company)
}

func (s *companiesService) DeleteCompany(ctx context.Context, input *requests.DeleteCompanyRequest) error {
	if err := s.commonRepo.CheckIfOrganizationExists(ctx, input.OrganizationID); err != nil {
		return err
	}
	return s.repo.DeleteCompany(ctx, input.OrganizationID, input.ID)
}