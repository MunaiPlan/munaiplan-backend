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

func (s *companiesService) GetCompanies(ctx context.Context) ([]*entities.Company, error) {
	return s.repo.GetCompanies(ctx)
}

func (s *companiesService) GetCompanyByName(ctx context.Context, name string) (*entities.Company, error) {
	return s.repo.GetCompanyByName(ctx, name)
}

func (s *companiesService) CreateCompany(ctx context.Context, input *requests.CreateCompanyRequest) error {
	company := &entities.Company{
		Name: input.Name,
		Division: input.Division,
		Group: input.Group,
		Representative: input.Representative,
		Address: input.Address,
		Phone: input.Phone,
	}
	return s.repo.CreateCompany(ctx, input.OrganizationID, company)
}

func (s *companiesService) UpdateCompany(ctx context.Context, input *requests.UpdateCompanyRequest) (*entities.Company, error) {
	company := &entities.Company{
		ID: input.ID,
		Name: input.Name,
		Division: input.Division,
		Group: input.Group,
		Representative: input.Representative,
		Address: input.Address,
		Phone: input.Phone,
	}
	return s.repo.UpdateCompany(ctx, company)
}

func (s *companiesService) DeleteCompany(ctx context.Context, input *requests.DeleteCompanyRequest) error {
	return s.repo.DeleteCompany(ctx, input.ID)
}