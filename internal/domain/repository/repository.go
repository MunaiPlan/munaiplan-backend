package repository

import (
	"github.com/munaiplan/munaiplan-backend/internal/infrastructure/repositories/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	Users        UsersRepository
	Companies    CompaniesRepository
	Organizations OrganizationsRepository
}

func NewRepositories(db *gorm.DB) *Repository {
	return &Repository{
		Users:        postgres.NewUsersRepository(db),
		Companies:    postgres.NewCompaniesRepository(db),
		Organizations: postgres.NewOrganizationsRepository(db),
	}
}
