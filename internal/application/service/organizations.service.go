package service

import (
	"context"

	"github.com/munaiplan/munaiplan-backend/internal/application/types/requests"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/domain/repository"
)

type organizationsService struct {
	organizationsRepo repository.OrganizationsRepository
}

func NewOrganizationsService(organizationsRepo repository.OrganizationsRepository) *organizationsService {
	return &organizationsService{
		organizationsRepo: organizationsRepo,
	}
}

func (s *organizationsService) CreateOrganization(ctx context.Context, input *requests.CreateOrganizationRequest) error {
	organization := &entities.Organization{
		Name: input.Name,
		Phone: input.Phone,
		Address: input.Address,
		Email: input.Email,
	}
	return s.organizationsRepo.CreateOrganization(ctx, organization)
}

func (s *organizationsService) UpdateOrganization(ctx context.Context, input *requests.UpdateOrganizationRequest) (*entities.Organization, error) {
	organization := &entities.Organization{
		ID: input.ID,
		Name: input.Name,
		Phone: input.Phone,
		Address: input.Address,
		Email: input.Email,
	}
	return s.organizationsRepo.UpdateOrganization(ctx, organization)
}

func (s *organizationsService) DeleteOrganization(ctx context.Context, input *requests.DeleteOrganizationRequest) error {
	return s.organizationsRepo.DeleteOrganization(ctx, input.ID)
}

func (s *organizationsService) GetOrganizations(ctx context.Context) ([]*entities.Organization, error) {
	return s.organizationsRepo.GetOrganizations(ctx)
}

func (s *organizationsService) GetOrganizationByName(ctx context.Context, input *requests.GetOrganizationByNameRequest) (*entities.Organization, error) {
	return s.organizationsRepo.GetOrganizationByName(ctx, input.Name)
}