package service

import (
	"context"

	"github.com/munaiplan/munaiplan-backend/internal/application/dto/requests"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/domain/repository"
)

type sitesService struct {
	commonRepo repository.CommonRepository
	repo       repository.SitesRepository
}

func NewSitesService(repo repository.SitesRepository, commonRepo repository.CommonRepository) *sitesService {
	return &sitesService{
		repo:       repo,
		commonRepo: commonRepo,
	}
}

func (s *sitesService) GetSites(ctx context.Context, input *requests.GetSitesRequest) ([]*entities.Site, error) {
	if err := s.commonRepo.CheckIfFieldExists(ctx, input.FieldID); err != nil {
		return nil, err
	}

	return s.repo.GetSites(ctx, input.FieldID)
}

func (s *sitesService) GetSiteByID(ctx context.Context, input *requests.GetSiteByIDRequest) (*entities.Site, error) {
	return s.repo.GetSiteByID(ctx, input.ID)
}

func (s *sitesService) CreateSite(ctx context.Context, input *requests.CreateSiteRequest) error {
	if err := s.commonRepo.CheckIfFieldExists(ctx, input.FieldID); err != nil {
		return err
	}

	site := &entities.Site{
		Name:    input.Body.Name,
		Area:    input.Body.Area,
		Block:   input.Body.Block,
		Azimuth: input.Body.Azimuth,
		Country: input.Body.Country,
		State:   input.Body.State,
		Region:  input.Body.Region,
	}

	return s.repo.CreateSite(ctx, input.FieldID, site)
}

func (s *sitesService) UpdateSite(ctx context.Context, input *requests.UpdateSiteRequest) (*entities.Site, error) {
	site := &entities.Site{
		ID:      input.ID,
		Name:    input.Body.Name,
		Area:    input.Body.Area,
		Block:   input.Body.Block,
		Azimuth: input.Body.Azimuth,
		Country: input.Body.Country,
		State:   input.Body.State,
		Region:  input.Body.Region,
	}

	return s.repo.UpdateSite(ctx, site)
}

func (s *sitesService) DeleteSite(ctx context.Context, input *requests.DeleteSiteRequest) error {
	return s.repo.DeleteSite(ctx, input.ID)
}
