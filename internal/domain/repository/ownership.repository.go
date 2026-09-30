package repository

import "context"

// OwnershipRepository resolves which organization owns a hierarchy record.
type OwnershipRepository interface {
	// OrganizationOf returns the owning organization id, or "" when the live record does not exist.
	OrganizationOf(ctx context.Context, resource, id string) (string, error)
}
