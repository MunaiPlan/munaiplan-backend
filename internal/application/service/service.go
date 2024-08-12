package service

import (
	"context"

	"github.com/munaiplan/munaiplan-backend/internal/application/dto/requests"
	"github.com/munaiplan/munaiplan-backend/internal/application/dto/responses"
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
	GetCompanyByName(ctx context.Context, input *requests.GetCompanyByNameRequest) (*entities.Company, error)
	CreateCompany(ctx context.Context, input *requests.CreateCompanyRequest) error
	UpdateCompany(ctx context.Context, input *requests.UpdateCompanyRequest) (*entities.Company, error)
	DeleteCompany(ctx context.Context, input *requests.DeleteCompanyRequest) error
}

type Services struct {
	// TODO() Implement cache
	// CatalogCache *catalog.CatalogCache
	Users
	Companies
	Organizations
}

func NewServices(repos *repository.Repository, jwt helpers.Jwt) *Services {
	return &Services{
		Users:         NewUsersService(repos.Users, repos.Common, jwt),
		Companies:     NewCompaniesService(repos.Companies, repos.Common),
		Organizations: NewOrganizationsService(repos.Organizations),
		// CatalogCache: deps.CatalogCache,
	}
}
