package catalog

import "github.com/munaiplan/munaiplan-backend/internal/app/config"

type CatalogCache struct {
	ApiDrillCollar ApiDrillCatalog
}

func NewCatalogCache(cfg config.CatalogCacheConfig) *CatalogCache {
	return &CatalogCache{
		ApiDrillCollar: NewApiDrillCollarCatalogWrapper(cfg.ApiDrillCollarPath),
	}
}