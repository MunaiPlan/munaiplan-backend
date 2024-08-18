package repository

import "context"

type CommonRepository interface {
	CheckIfOrganizationExists(ctx context.Context, organizationId string) error
	CheckIfCompanyExists(ctx context.Context, companyId string) error
	CheckIfUserExistsByEmail(ctx context.Context, email string) error
	CheckIfFieldExists(ctx context.Context, fieldId string) error
	CheckIfSiteExists(ctx context.Context, siteId string) error
	CheckIfWellExists(ctx context.Context, wellId string) error
	CheckIfWellboreExists(ctx context.Context, wellboreId string) error
}