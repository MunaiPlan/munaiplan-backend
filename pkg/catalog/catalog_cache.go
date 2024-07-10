package catalog

import (
	"github.com/munaiplan/munaiplan-backend/infrastructure/configs"
	catalogParser "github.com/munaiplan/munaiplan-backend/pkg/catalog/parser"
	"github.com/xuri/excelize/v2"
)

type CatalogCache struct {
	ApiDrillCollar             catalogParser.ApiDrillCollarCatalog
	ApiDrillPipe               catalogParser.ApiDrillPipeCatalog
	Additional                 catalogParser.AdditionalCatalog
	AdjustableGaugeStabilizers catalogParser.AdjustableGaugeStabilizersCatalog
}

func NewCatalogCache(cfg configs.CatalogConfig, file *excelize.File) *CatalogCache {
	return &CatalogCache{
		ApiDrillCollar:             catalogParser.NewApiDrillCollarCatalogWrapper(cfg.ApiDrillCollar.CatalogPath, cfg.CatalogCode, cfg.ApiDrillCollar.CatalogItemCode, file),
		ApiDrillPipe:               catalogParser.NewApiDrillPipeCatalogWrapper(cfg.ApiDrillPipe.CatalogPath, cfg.CatalogCode, cfg.ApiDrillPipe.CatalogItemCode, file),
		Additional:                 catalogParser.NewAdditionalCatalogWrapper(cfg.Additional.CatalogPath, cfg.CatalogCode, cfg.Additional.CatalogItemCode, file),
		AdjustableGaugeStabilizers: catalogParser.NewAdjustableGaugeStabilizersCatalogWrapper(cfg.AdjustableGaugeStablilizers.CatalogPath, cfg.CatalogCode, cfg.AdjustableGaugeStablilizers.CatalogItemCode, file),
	}
}
