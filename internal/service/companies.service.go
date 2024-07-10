package service

import (
	"github.com/munaiplan/munaiplan-backend/internal/domain"
	"github.com/munaiplan/munaiplan-backend/internal/repository"
)

type CompaniesService struct {
	repo repository.CompaniesRepository
}

func NewCompaniesService(repo repository.CompaniesRepository) *CompaniesService {
	return &CompaniesService{
		repo: repo,
	}
}

func (s *CompaniesService) CreateCompany(company *domain.Company) error {
	return s.repo.CreateCompany(company)
}

func (s *CompaniesService) UpdateCompany(company *domain.Company) error {
	oldCompany, err := s.repo.GetCompanyByID(company.ID)
	if err != nil {
		return err
	}

	if s.compareCompanies(oldCompany, company) {
		return domain.ErrCompanyWasNotUpdated
	}

	return s.repo.UpdateCompany(company)
}

func (s *CompaniesService) DeleteCompany(id string) error {
	return s.repo.DeleteCompany(id)
}


func (s *CompaniesService) compareCompanies(oldCompany *domain.Company, newCompany *domain.Company) bool {
	if (len(oldCompany.Fields) != len(newCompany.Fields)) {
		return false
	}

	for i := 0; i < len(oldCompany.Fields); i++ {
		if oldCompany.Fields[i].ID != newCompany.Fields[i].ID {
			return false
		}
	}

	switch {
	case oldCompany.Name != newCompany.Name:
		return false
	case oldCompany.Division != newCompany.Division:
		return false
	case oldCompany.Group != newCompany.Group:
		return false
	case oldCompany.Address != newCompany.Address:
		return false
	case oldCompany.Representative != newCompany.Representative:
		return false
	case oldCompany.Phone != newCompany.Phone:
		return false
	default:
		return true
	}
}