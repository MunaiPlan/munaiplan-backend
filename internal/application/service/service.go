package service

import (
	"context"

	"github.com/munaiplan/munaiplan-backend/internal/application/types/requests"
	"github.com/munaiplan/munaiplan-backend/internal/application/types/responses"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/domain/repository"
	"github.com/munaiplan/munaiplan-backend/internal/helpers"
)

type Users interface {
	SignIn(ctx context.Context, input *requests.UserSignInRequest) (*responses.TokenResponse, error)
	SignUp(ctx context.Context, input *requests.UserSignUpRequest) error
}

type Organizations interface {
	CreateOrganization(ctx context.Context, input *requests.CreateOrganizationRequest) error
	UpdateOrganization(ctx context.Context, input *requests.UpdateOrganizationRequest) (*entities.Organization, error)
	DeleteOrganization(ctx context.Context, input *requests.DeleteOrganizationRequest) error
	GetOrganizations(ctx context.Context) ([]*entities.Organization, error)
	GetOrganizationByName(ctx context.Context, input *requests.GetOrganizationByNameRequest) (*entities.Organization, error)
}

type Companies interface {
	GetCompanies(ctx context.Context, input *requests.GetCompaniesRequest) ([]*entities.Company, error)
	GetCompanyByID(ctx context.Context, input *requests.GetCompanyByIDRequest) (*entities.Company, error)
	CreateCompany(ctx context.Context, input *requests.CreateCompanyRequest) error
	UpdateCompany(ctx context.Context, input *requests.UpdateCompanyRequest) (*entities.Company, error)
	DeleteCompany(ctx context.Context, input *requests.DeleteCompanyRequest) error
}

type Fields interface {
	GetFields(ctx context.Context, input *requests.GetFieldsRequest) ([]*entities.Field, error)
	GetFieldByID(ctx context.Context, input *requests.GetFieldByIDRequest) (*entities.Field, error)
	CreateField(ctx context.Context, input *requests.CreateFieldRequest) error
	UpdateField(ctx context.Context, input *requests.UpdateFieldRequest) (*entities.Field, error)
	DeleteField(ctx context.Context, input *requests.DeleteFieldRequest) error
}

type Sites interface {
	GetSites(ctx context.Context, input *requests.GetSitesRequest) ([]*entities.Site, error)
	GetSiteByID(ctx context.Context, input *requests.GetSiteByIDRequest) (*entities.Site, error)
	CreateSite(ctx context.Context, input *requests.CreateSiteRequest) error
	UpdateSite(ctx context.Context, input *requests.UpdateSiteRequest) (*entities.Site, error)
	DeleteSite(ctx context.Context, input *requests.DeleteSiteRequest) error
}

type Wells interface {
	GetWells(ctx context.Context, input *requests.GetWellsRequest) ([]*entities.Well, error)
	GetWellByID(ctx context.Context, input *requests.GetWellByIDRequest) (*entities.Well, error)
	CreateWell(ctx context.Context, input *requests.CreateWellRequest) error
	UpdateWell(ctx context.Context, input *requests.UpdateWellRequest) (*entities.Well, error)
	DeleteWell(ctx context.Context, input *requests.DeleteWellRequest) error
}

type Wellbores interface {
	GetWellbores(ctx context.Context, input *requests.GetWellboresRequest) ([]*entities.Wellbore, error)
	GetWellboreByID(ctx context.Context, input *requests.GetWellboreByIDRequest) (*entities.Wellbore, error)
	CreateWellbore(ctx context.Context, input *requests.CreateWellboreRequest) error
	UpdateWellbore(ctx context.Context, input *requests.UpdateWellboreRequest) (*entities.Wellbore, error)
	DeleteWellbore(ctx context.Context, input *requests.DeleteWellboreRequest) error
}

type Designs interface {
	GetDesigns(ctx context.Context, input *requests.GetDesignsRequest) ([]*entities.Design, error)
	GetDesignByID(ctx context.Context, input *requests.GetDesignByIDRequest) (*entities.Design, error)
	CreateDesign(ctx context.Context, input *requests.CreateDesignRequest) error
	UpdateDesign(ctx context.Context, input *requests.UpdateDesignRequest) (*entities.Design, error)
	DeleteDesign(ctx context.Context, input *requests.DeleteDesignRequest) error
}

type Trajectories interface {
	GetTrajectories(ctx context.Context, input *requests.GetTrajectoriesRequest) ([]*entities.Trajectory, error)
	GetTrajectoryByID(ctx context.Context, input *requests.GetTrajectoryByIDRequest) (*entities.Trajectory, error)
	CreateTrajectory(ctx context.Context, input *requests.CreateTrajectoryRequest) error
	UpdateTrajectory(ctx context.Context, input *requests.UpdateTrajectoryRequest) (*entities.Trajectory, error)
	DeleteTrajectory(ctx context.Context, input *requests.DeleteTrajectoryRequest) error
}

type Cases interface {
	GetCases(ctx context.Context, input *requests.GetCasesRequest) ([]*entities.Case, error)
	GetCaseByID(ctx context.Context, input *requests.GetCaseByIDRequest) (*entities.Case, error)
	CreateCase(ctx context.Context, input *requests.CreateCaseRequest) error
	UpdateCase(ctx context.Context, input *requests.UpdateCaseRequest) (*entities.Case, error)
	DeleteCase(ctx context.Context, input *requests.DeleteCaseRequest) error
}

type Datums interface {
	GetDatumsByCaseID(ctx context.Context, input *requests.GetDatumsByCaseIDRequest) ([]*entities.Datum, error)
	GetDatumByID(ctx context.Context, input *requests.GetDatumByIDRequest) (*entities.Datum, error)
	CreateDatum(ctx context.Context, input *requests.CreateDatumRequest) error
	UpdateDatum(ctx context.Context, input *requests.UpdateDatumRequest) (*entities.Datum, error)
	DeleteDatum(ctx context.Context, input *requests.DeleteDatumRequest) error
}

type Services struct {
	// TODO() Implement cache
	// CatalogCache *catalog.CatalogCache
	Users
	Companies
	Organizations
	Fields
	Sites
	Wells
	Wellbores
	Designs
	Trajectories
	Cases
	Datums
}

func NewServices(repos *repository.Repository, jwt helpers.Jwt) *Services {
	return &Services{
		Users:         NewUsersService(repos.Users, repos.Common, jwt),
		Companies:     NewCompaniesService(repos.Companies, repos.Common),
		Organizations: NewOrganizationsService(repos.Organizations),
		Fields:        NewFieldsService(repos.Fields, repos.Common),
		Sites:         NewSitesService(repos.Sites, repos.Common),
		Wells:         NewWellsService(repos.Wells, repos.Common),
		Wellbores:     NewWellboresService(repos.Wellbores, repos.Common),
		Designs:       NewDesignsService(repos.Designs, repos.Common),
		Trajectories:  NewTrajectoriesService(repos.Trajectories, repos.Common),
		Cases:         NewCasesService(repos.Cases, repos.Common),
		Datums:        NewDatumsService(repos.Datums, repos.Common),
		// CatalogCache: deps.CatalogCache,
	}
}
