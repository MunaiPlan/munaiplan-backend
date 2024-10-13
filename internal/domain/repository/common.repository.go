package repository

import (
	"context"

	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
)

type CommonRepository interface {
	CheckIfOrganizationExists(ctx context.Context, organizationId string) error
	CheckIfCompanyExists(ctx context.Context, companyId string) error
	CheckIfUserExistsByEmail(ctx context.Context, email string) error
	CheckIfFieldExists(ctx context.Context, fieldId string) error
	CheckIfSiteExists(ctx context.Context, siteId string) error
	CheckIfWellExists(ctx context.Context, wellId string) error
	CheckIfWellboreExists(ctx context.Context, wellboreId string) error
	CheckIfDesignExists(ctx context.Context, designId string) error
	CheckIfTrajectoryExists(ctx context.Context, trajectoryId string) error
	CheckIfCaseExists(ctx context.Context, caseId string) error
	CheckCaseCompleteness(ctx context.Context, caseID string) (bool, error)
	CheckIfFluidExists(ctx context.Context, fluidId string) (bool, error)
	CheckIfRigExists(ctx context.Context, rigId string) (bool, error)
	CheckIfHoleExists(ctx context.Context, holeId string) (bool, error)
	CheckIfStringExists(ctx context.Context, stringId string) (bool, error)
	CheckIfPorePressureExists(ctx context.Context, porePressureId string) (bool, error)
	CheckIfFractureGradientExists(ctx context.Context, fractureGradientId string) (bool, error)
	GetTrajectoryByCaseID(ctx context.Context, caseID string) (*entities.Trajectory, error)
}
