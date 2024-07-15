package repository

import (
	"context"

	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
)

type CompaniesRepository interface {
	CreateCompany(ctx context.Context, organizationID string, company *entities.Company) error
	GetCompanyByID(ctx context.Context, id string) (*entities.Company, error)
	GetCompanyByName(ctx context.Context, name string) (*entities.Company, error)
	GetCompanies(ctx context.Context, ) ([]*entities.Company, error)
	UpdateCompany(ctx context.Context, company *entities.Company) (*entities.Company, error)
	DeleteCompany(ctx context.Context, id string) error
}