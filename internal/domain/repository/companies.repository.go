package repository

import (
	"context"

	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
)

type CompaniesRepository interface {
	CreateCompany(ctx context.Context, organizationID string, company *entities.Company) error
	GetCompanyByID(ctx context.Context, id, organizationId string) (*entities.Company, error)
	GetCompanyByName(ctx context.Context, name, organizationId string) (*entities.Company, error)
	GetCompanies(ctx context.Context, organizationId string) ([]*entities.Company, error)
	UpdateCompany(ctx context.Context, organizationId string, company *entities.Company) (*entities.Company, error)
	DeleteCompany(ctx context.Context, organizationId, id string) error
}