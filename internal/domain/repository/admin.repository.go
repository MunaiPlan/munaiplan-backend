package repository

import (
	"context"

	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
)

// AdminRepository backs tenant and account provisioning. Implementations must create an
// organization and its first user atomically and report duplicate emails as ErrEmailTaken.
type AdminRepository interface {
	ListOrganizations(ctx context.Context) ([]*entities.OrganizationSummary, error)
	CreateOrganizationWithUser(ctx context.Context, org *entities.Organization, user *entities.User) error
	ListUsers(ctx context.Context, organizationID string) ([]*entities.User, error)
	CreateUser(ctx context.Context, organizationID string, user *entities.User) error
}
