package repository

import (
	"context"

	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/importers/wellplan"
)

// ImportsRepository persists parsed WellPlan cases atomically inside one organization.
type ImportsRepository interface {
	// FindByFingerprint returns a live previous import of the same files, or nil.
	FindByFingerprint(ctx context.Context, organizationID, fingerprint string) (*entities.ImportResult, error)
	// SaveWellPlan reuses hierarchy levels by name, always creates a new trajectory and case,
	// and stores the full report for provenance and reference comparison.
	SaveWellPlan(ctx context.Context, organizationID, userID, fingerprint string, report *wellplan.Report) (*entities.ImportResult, error)
	// GetReport returns the stored report for a case in the organization, or nil if none.
	GetReport(ctx context.Context, organizationID, caseID string) (*wellplan.Report, error)
}
