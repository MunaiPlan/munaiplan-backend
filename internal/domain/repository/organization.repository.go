package repository

import (
	"context"

	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
)

type OrganizationsRepository interface {
	CreateOrganization(ctx context.Context, organization *entities.Organization) error
	GetOrganizationByID(ctx context.Context, id string) (*entities.Organization, error)
	GetOrganizationByName(ctx context.Context, name string) (*entities.Organization, error)
	GetOrganizations(ctx context.Context) ([]*entities.Organization, error)
	UpdateOrganization(ctx context.Context, organization *entities.Organization) (*entities.Organization, error)
	DeleteOrganization(ctx context.Context, id string) error
}