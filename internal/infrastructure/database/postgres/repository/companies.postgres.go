package repository

import (
	"github.com/munaiplan/munaiplan-backend/internal/infrastructure/database/postgres/infra/models"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"gorm.io/gorm"
)

type companiesRepository struct {
	db *gorm.DB
}

func NewcompaniesRepository(db *gorm.DB) *companiesRepository {
	return &companiesRepository{db: db}
}

func (r *companiesRepository) CreateCompany(company *domain.Company) error {
	gormCompany := toGormCompany(company)
	result := r.db.Create(gormCompany)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *companiesRepository) GetCompanyByID(id string) (*domain.Company, error) {
	var company models.Company
	var res domain.Company
	result := r.db.Where("id = ?", id).First(&company)
	if result.Error != nil {
		return nil, result.Error
	}

	res = toDomainCompany(&company)
	return &res, nil
}

func (r *companiesRepository) GetCompanyByName(name string) (*domain.Company, error) {
	var company models.Company
	var res domain.Company
	result := r.db.Where("name = ?", name).First(&company)
	if result.Error != nil {
		return nil, result.Error
	}

	res = toDomainCompany(&company)
	return &res, nil
}

func (r *companiesRepository) GetCompanies() ([]*domain.Company, error) {
	var companies []*models.Company
	var res []*domain.Company
	result := r.db.Find(&companies)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, company := range companies {
		temp := toDomainCompany(company)
		res = append(res, &temp)
	}
	return res, nil
}

func (r *companiesRepository) UpdateCompany(company *domain.Company) error {
	gormCompany := toGormCompany(company)
	result := r.db.Save(gormCompany)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *companiesRepository) DeleteCompany(id string) error {
	result := r.db.Where("id = ?", id).Delete(&models.Company{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// toDomainCompany maps the GORM Company model to the domain Company entity.
func toDomainCompany(companyModel *models.Company) domain.Company {
	return domain.Company{
		ID:             companyModel.ID.String(),
		Name:           companyModel.Name,
		Division:       companyModel.Division,
		Group:          companyModel.Group,
		Representative: companyModel.Representative,
		Address:        companyModel.Address,
		Phone:          companyModel.Phone,
		CreatedAt:      companyModel.CreatedAt,
		UpdatedAt:      companyModel.UpdatedAt,
	}
}

// toGormCompany maps the domain Company entity to the GORM Company model.
func toGormCompany(company *domain.Company) models.Company {
	return models.Company{
		Name:           company.Name,
		Division:       company.Division,
		Group:          company.Group,
		Representative: company.Representative,
		Address:        company.Address,
		Phone:          company.Phone,
	}
}
