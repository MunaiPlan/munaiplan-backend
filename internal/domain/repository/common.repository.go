package repository

import "context"

type CommonRepository interface {
	CheckIfOrganizationExists(ctx context.Context, organizationId string) error
	CheckIfCompanyExists(ctx context.Context, companyId string) error
	CheckIfUserExistsByEmail(ctx context.Context, email string) error
}