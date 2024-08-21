package postgres

import (
	"context"
	"reflect"

	"github.com/google/uuid"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/infrastructure/drivers/postgres/models"
	"github.com/munaiplan/munaiplan-backend/internal/infrastructure/types"
	"gorm.io/gorm"
)

type companiesRepository struct {
	db *gorm.DB
}

func NewCompaniesRepository(db *gorm.DB) *companiesRepository {
	return &companiesRepository{db: db}
}

func (r *companiesRepository) CreateCompany(ctx context.Context, organizationID string, company *entities.Company) error {
	gormCompany := r.toGormCompany(company)
	orgId, err := uuid.Parse(organizationID)
	if err != nil {
		return err
	}
	gormCompany.OrganizationID = orgId

	result := r.db.WithContext(ctx).Create(gormCompany)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *companiesRepository) GetCompanyByID(ctx context.Context, id, organizationId string) (*entities.Company, error) {
	var company models.Company
	var res entities.Company
	result := r.db.WithContext(ctx).Where("id = ? AND organization_id = ?", id, organizationId).First(&company)
	if result.Error != nil {
		return nil, result.Error
	}

	res = r.toDomainCompany(&company)
	return &res, nil
}

func (r *companiesRepository) GetCompanies(ctx context.Context, organizationId string) ([]*entities.Company, error) {
	var companies []*models.Company
	var res []*entities.Company
	result := r.db.WithContext(ctx).Where("organization_id = ?", organizationId).Find(&companies)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, company := range companies {
		temp := r.toDomainCompany(company)
		res = append(res, &temp)
	}
	return res, nil
}

func (r *companiesRepository) UpdateCompany(ctx context.Context, organizationId string, company *entities.Company) (*entities.Company, error) {
	gormCompany := r.toGormCompany(company)
	oldCompany := models.Company{}
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		query := tx.WithContext(ctx).Where("id = ? AND organization_id = ?", company.ID, organizationId).First(&oldCompany)
		if query.Error != nil {
			return query.Error
		}

		if reflect.DeepEqual(&gormCompany, &oldCompany) {
			return types.ErrComanyNotChanged
		}

		err := tx.WithContext(ctx).Model(&oldCompany).Updates(gormCompany).Error
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	res := r.toDomainCompany(&oldCompany)

	return &res, nil
}

func (r *companiesRepository) DeleteCompany(ctx context.Context, organizationId string, id string) error {
	result := r.db.WithContext(ctx).Where("id = ? AND organization_id = ?", id, organizationId).Delete(&models.Company{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// toDomainCompany maps the GORM Company model to the domain Company entity.
func (r *companiesRepository) toDomainCompany(companyModel *models.Company) entities.Company {
	return entities.Company{
		ID:             companyModel.ID.String(),
		Name:           companyModel.Name,
		Division:       companyModel.Division,
		Group:          companyModel.Group,
		Representative: companyModel.Representative,
		Address:        companyModel.Address,
		Phone:          companyModel.Phone,
	}
}

// toGormCompany maps the domain Company entity to the GORM Company model.
func (r *companiesRepository) toGormCompany(company *entities.Company) *models.Company {
	return &models.Company{
		Name:           company.Name,
		Division:       company.Division,
		Group:          company.Group,
		Representative: company.Representative,
		Address:        company.Address,
		Phone:          company.Phone,
	}
}
