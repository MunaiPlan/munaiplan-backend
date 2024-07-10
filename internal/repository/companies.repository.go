package repository

import "github.com/munaiplan/munaiplan-backend/internal/domain"

type CompaniesRepository interface {
	CreateCompany(company *domain.Company) error
	GetCompanyByID(id string) (*domain.Company, error)
	GetCompanyByName(name string) (*domain.Company, error)
	GetCompanies() ([]*domain.Company, error)
	UpdateCompany(company *domain.Company) error
	DeleteCompany(id string) error
}