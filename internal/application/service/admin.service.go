package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/munaiplan/munaiplan-backend/internal/application/types/requests"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/domain/repository"
	"github.com/munaiplan/munaiplan-backend/internal/helpers"
)

// ValidationError is a client-correctable input problem (HTTP 400).
type ValidationError struct{ Reason string }

func (e *ValidationError) Error() string { return e.Reason }

type adminService struct {
	repo repository.AdminRepository
}

func NewAdminService(repo repository.AdminRepository) *adminService {
	return &adminService{repo: repo}
}

func (s *adminService) ListOrganizations(ctx context.Context) ([]*entities.OrganizationSummary, error) {
	orgs, err := s.repo.ListOrganizations(ctx)
	if orgs == nil {
		orgs = []*entities.OrganizationSummary{}
	}
	return orgs, err
}

func (s *adminService) CreateOrganization(ctx context.Context, in *requests.AdminCreateOrganizationRequest) (*entities.Organization, *entities.User, error) {
	org := &entities.Organization{
		Name:    strings.TrimSpace(in.Organization.Name),
		Email:   normalizeEmail(in.Organization.Email),
		Phone:   strings.TrimSpace(in.Organization.Phone),
		Address: strings.TrimSpace(in.Organization.Address),
	}
	if org.Name == "" {
		return nil, nil, &ValidationError{"organization name is required"}
	}
	user, err := newAccount(&in.User)
	if err != nil {
		return nil, nil, err
	}
	if err := s.repo.CreateOrganizationWithUser(ctx, org, user); err != nil {
		return nil, nil, err
	}
	return org, user, nil
}

func (s *adminService) ListUsers(ctx context.Context, organizationID string) ([]*entities.User, error) {
	users, err := s.repo.ListUsers(ctx, organizationID)
	if users == nil {
		users = []*entities.User{}
	}
	return users, err
}

func (s *adminService) CreateUser(ctx context.Context, organizationID string, in *requests.AdminUserInput) (*entities.User, error) {
	user, err := newAccount(in)
	if err != nil {
		return nil, err
	}
	if err := s.repo.CreateUser(ctx, organizationID, user); err != nil {
		return nil, err
	}
	return user, nil
}

// newAccount normalises input, applies the password policy and hashes the password.
func newAccount(in *requests.AdminUserInput) (*entities.User, error) {
	user := &entities.User{
		Name:    strings.TrimSpace(in.Name),
		Surname: strings.TrimSpace(in.Surname),
		Email:   normalizeEmail(in.Email),
		Phone:   strings.TrimSpace(in.Phone),
		Role:    in.Role,
	}
	if user.Role == "" {
		user.Role = entities.RoleUser
	}
	if user.Role != entities.RoleUser && user.Role != entities.RoleAdmin {
		return nil, &ValidationError{fmt.Sprintf("unknown role %q", in.Role)}
	}
	if user.Name == "" || user.Surname == "" {
		return nil, &ValidationError{"user name and surname are required"}
	}
	if err := helpers.ValidatePassword(in.Password); err != nil {
		return nil, &ValidationError{err.Error()}
	}
	hash, err := helpers.HashPassword(in.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hash
	return user, nil
}

func normalizeEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}
