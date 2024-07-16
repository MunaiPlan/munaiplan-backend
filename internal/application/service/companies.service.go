package service

import (
	"context"

	"github.com/munaiplan/munaiplan-backend/internal/application/dto/requests"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/domain/repository"
)

type companiesService struct {
	repo repository.CompaniesRepository
}

func NewCompaniesService(repo repository.CompaniesRepository) *companiesService {
	return &companiesService{
		repo: repo,
	}
}

func (s *companiesService) GetCompanies(ctx context.Context, input *requests.GetCompaniesRequest) ([]*entities.Company, error) {
	return s.repo.GetCompanies(ctx, input.OrganizationID)
}

func (s *companiesService) GetCompanyByName(ctx context.Context, input *requests.GetCompanyByNameRequest) (*entities.Company, error) {
	return s.repo.GetCompanyByName(ctx, input.Name, input.OrganizationID)
}

func (s *companiesService) CreateCompany(ctx context.Context, input *requests.CreateCompanyRequest) error {
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
	return s.repo.DeleteCompany(ctx, input.OrganizationID, input.ID)
}