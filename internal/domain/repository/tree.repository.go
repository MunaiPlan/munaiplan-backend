package repository

import (
	"context"

	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
)

// TreeRepository reads the lightweight navigation tree of one organization.
type TreeRepository interface {
	Tree(ctx context.Context, organizationID string) ([]*entities.TreeNode, error)
}
