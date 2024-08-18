package repository

import (
	"context"

	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
)

type SitesRepository interface {
	CreateSite(ctx context.Context, fieldID string, site *entities.Site) error
	GetSiteByID(ctx context.Context, id string) (*entities.Site, error)
	GetSites(ctx context.Context, fieldID string) ([]*entities.Site, error)
	UpdateSite(ctx context.Context, site *entities.Site) (*entities.Site, error)
	DeleteSite(ctx context.Context, id string) error
}
