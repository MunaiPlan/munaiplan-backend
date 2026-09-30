package postgres

import (
	"context"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

// parentOf describes the well hierarchy: table -> (parent table, foreign key column).
// Companies hold organization_id directly. Only these whitelisted names reach SQL.
var parentOf = map[string][2]string{
	"fields":             {"companies", "company_id"},
	"sites":              {"fields", "field_id"},
	"wells":              {"sites", "site_id"},
	"wellbores":          {"wells", "well_id"},
	"designs":            {"wellbores", "wellbore_id"},
	"trajectories":       {"designs", "design_id"},
	"cases":              {"trajectories", "trajectory_id"},
	"holes":              {"cases", "case_id"},
	"strings":            {"cases", "case_id"},
	"fluids":             {"cases", "case_id"},
	"rigs":               {"cases", "case_id"},
	"pore_pressures":     {"cases", "case_id"},
	"fracture_gradients": {"cases", "case_id"},
}

type ownershipRepository struct {
	db      *gorm.DB
	queries map[string]string
}

func NewOwnershipRepository(db *gorm.DB) *ownershipRepository {
	queries := map[string]string{"companies": ownershipQuery("companies")}
	for table := range parentOf {
		queries[table] = ownershipQuery(table)
	}
	return &ownershipRepository{db: db, queries: queries}
}

// ownershipQuery joins from the table up to companies; every level must be live.
func ownershipQuery(table string) string {
	var joins strings.Builder
	alias, current := "t0", table
	for i := 1; current != "companies"; i++ {
		link := parentOf[current]
		next := fmt.Sprintf("t%d", i)
		fmt.Fprintf(&joins, " JOIN %s %s ON %s.id = %s.%s AND %s.deleted_at IS NULL", link[0], next, next, alias, link[1], next)
		alias, current = next, link[0]
	}
	return fmt.Sprintf("SELECT %s.organization_id::text FROM %s t0%s WHERE t0.id = ? AND t0.deleted_at IS NULL", alias, table, joins.String())
}

func (r *ownershipRepository) OrganizationOf(ctx context.Context, resource, id string) (string, error) {
	query, ok := r.queries[resource]
	if !ok {
		return "", fmt.Errorf("unknown resource %q", resource)
	}
	var org string
	err := r.db.WithContext(ctx).Raw(query, id).Scan(&org).Error
	return org, err
}
