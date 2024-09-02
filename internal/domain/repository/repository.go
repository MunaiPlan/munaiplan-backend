package repository

import (
	"github.com/munaiplan/munaiplan-backend/internal/infrastructure/repositories/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	Common        CommonRepository
	Users         UsersRepository
	Companies     CompaniesRepository
	Organizations OrganizationsRepository
	Fields        FieldsRepository
	Sites         SitesRepository
	Wells         WellsRepository
	Wellbores     WellboresRepository
	Designs       DesignsRepository
	Trajectories  TrajectoriesRepository
	Cases         CasesRepository
	Holes         HolesRepository
}

func NewRepositories(db *gorm.DB) *Repository {
	return &Repository{
		Common:        postgres.NewCommonRepository(db),
		Users:         postgres.NewUsersRepository(db),
		Companies:     postgres.NewCompaniesRepository(db),
		Organizations: postgres.NewOrganizationsRepository(db),
		Fields:        postgres.NewFieldsRepository(db),
		Sites:         postgres.NewSitesRepository(db),
		Wells:         postgres.NewWellsRepository(db),
		Wellbores:     postgres.NewWellboresRepository(db),
		Designs:       postgres.NewDesignsRepository(db),
		Trajectories:  postgres.NewTrajectoriesRepository(db),
		Cases:         postgres.NewCasesRepository(db),
		Holes:         postgres.NewHolesRepository(db),
	}
}
