package service

import (
	"context"

	"github.com/munaiplan/munaiplan-backend/internal/application/types/requests"
	"github.com/munaiplan/munaiplan-backend/internal/application/types/responses"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/domain/repository"
	"github.com/munaiplan/munaiplan-backend/internal/helpers"
	client "github.com/munaiplan/munaiplan-backend/internal/infrastructure/prediction_client"
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
	GetCompaniesWithComponents(ctx context.Context, input *requests.GetCompaniesRequest) ([]*entities.Company, error)
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

type Holes interface {
	GetHoles(ctx context.Context, input *requests.GetHolesRequest) ([]*entities.Hole, error)
	GetHoleByID(ctx context.Context, input *requests.GetHoleByIDRequest) (*entities.Hole, error)
	CreateHole(ctx context.Context, input *requests.CreateHoleRequest) error
	UpdateHole(ctx context.Context, input *requests.UpdateHoleRequest) (*entities.Hole, error)
	DeleteHole(ctx context.Context, input *requests.DeleteHoleRequest) error
}

type Fluids interface {
	GetFluids(ctx context.Context, input *requests.GetFluidsRequest) ([]*entities.Fluid, error)
	GetFluidTypes(ctx context.Context) ([]*entities.FluidType, error)
	GetFluidByID(ctx context.Context, input *requests.GetFluidByIDRequest) (*entities.Fluid, error)
	CreateFluid(ctx context.Context, input *requests.CreateFluidRequest) error
	UpdateFluid(ctx context.Context, input *requests.UpdateFluidRequest) (*entities.Fluid, error)
	DeleteFluid(ctx context.Context, input *requests.DeleteFluidRequest) error
}

type Rigs interface {
	GetRigs(ctx context.Context, input *requests.GetRigsRequest) ([]*entities.Rig, error)
	GetRigByID(ctx context.Context, input *requests.GetRigByIDRequest) (*entities.Rig, error)
	CreateRig(ctx context.Context, input *requests.CreateRigRequest) error
	UpdateRig(ctx context.Context, input *requests.UpdateRigRequest) (*entities.Rig, error)
	DeleteRig(ctx context.Context, input *requests.DeleteRigRequest) error
}

type PorePressures interface {
	GetPorePressures(ctx context.Context, input *requests.GetPorePressuresRequest) ([]*entities.PorePressure, error)
	GetPorePressureByID(ctx context.Context, input *requests.GetPorePressureByIDRequest) (*entities.PorePressure, error)
	CreatePorePressure(ctx context.Context, input *requests.CreatePorePressureRequest) error
	UpdatePorePressure(ctx context.Context, input *requests.UpdatePorePressureRequest) (*entities.PorePressure, error)
	DeletePorePressure(ctx context.Context, input *requests.DeletePorePressureRequest) error
}

type FractureGradients interface {
	GetFractureGradients(ctx context.Context, input *requests.GetFractureGradientsRequest) ([]*entities.FractureGradient, error)
	GetFractureGradientByID(ctx context.Context, input *requests.GetFractureGradientByIDRequest) (*entities.FractureGradient, error)
	CreateFractureGradient(ctx context.Context, input *requests.CreateFractureGradientRequest) error
	UpdateFractureGradient(ctx context.Context, input *requests.UpdateFractureGradientRequest) (*entities.FractureGradient, error)
	DeleteFractureGradient(ctx context.Context, input *requests.DeleteFractureGradientRequest) error
}

type Strings interface {
	GetStrings(ctx context.Context, input *requests.GetStringsRequest) ([]*entities.String, error)
	GetStringByID(ctx context.Context, input *requests.GetStringByIDRequest) (*entities.String, error)
	CreateString(ctx context.Context, input *requests.CreateStringRequest) error
	UpdateString(ctx context.Context, input *requests.UpdateStringRequest) (*entities.String, error)
	DeleteString(ctx context.Context, input *requests.DeleteStringRequest) error
}

type TorqueAndDrag interface {
	CalculateEffectiveTensionFromMLModel(ctx context.Context, caseID string) (*responses.EffectiveTensionFromMLModelResponse, error)
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
	Holes
	Fluids
	Rigs
	PorePressures
	FractureGradients
	Strings
	TorqueAndDrag
}

func NewServices(repos *repository.Repository, jwt helpers.Jwt, mlServiceClientUrl string) *Services {
	return &Services{
		Users:             NewUsersService(repos.Users, repos.Common, jwt),
		Companies:         NewCompaniesService(repos.Companies, repos.Common),
		Organizations:     NewOrganizationsService(repos.Organizations),
		Fields:            NewFieldsService(repos.Fields, repos.Common),
		Sites:             NewSitesService(repos.Sites, repos.Common),
		Wells:             NewWellsService(repos.Wells, repos.Common),
		Wellbores:         NewWellboresService(repos.Wellbores, repos.Common),
		Designs:           NewDesignsService(repos.Designs, repos.Common),
		Trajectories:      NewTrajectoriesService(repos.Trajectories, repos.Common),
		Cases:             NewCasesService(repos.Cases, repos.Common),
		Holes:             NewHolesService(repos.Holes, repos.Common),
		Fluids:            NewFluidsService(repos.Fluids, repos.Common),
		Rigs:              NewRigsService(repos.Rigs, repos.Common),
		PorePressures:     NewPorePressuresService(repos.PorePressures, repos.Common),
		FractureGradients: NewFractureGradientsService(repos.FractureGradients, repos.Common),
		Strings:           NewStringsService(repos.Strings, repos.Common),
		TorqueAndDrag:     NewTorqueAndDragService(
			repos.Strings, 
			repos.Common,  
			client.NewTorqueAndDragClient(mlServiceClientUrl),
		),
		// CatalogCache: deps.CatalogCache,
	}
}
