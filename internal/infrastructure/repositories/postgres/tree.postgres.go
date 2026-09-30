package postgres

import (
	"context"

	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"gorm.io/gorm"
)

type treeRepository struct {
	db *gorm.DB
}

func NewTreeRepository(db *gorm.DB) *treeRepository {
	return &treeRepository{db: db}
}

// treeLevels orders the hierarchy; each row carries the id and name of every level (NULL below
// the deepest existing one), so one LEFT JOIN query returns the whole tree.
var treeLevels = []string{"company", "field", "site", "well", "wellbore", "design", "trajectory", "case"}

const treeQuery = `
SELECT c.id::text AS company_id, c.name AS company_name,
       f.id::text AS field_id, f.name AS field_name,
       s.id::text AS site_id, s.name AS site_name,
       w.id::text AS well_id, w.name AS well_name,
       wb.id::text AS wellbore_id, wb.name AS wellbore_name,
       d.id::text AS design_id, d.plan_name AS design_name,
       t.id::text AS trajectory_id, t.name AS trajectory_name,
       k.id::text AS case_id, k.case_name AS case_name
FROM companies c
LEFT JOIN fields f ON f.company_id = c.id AND f.deleted_at IS NULL
LEFT JOIN sites s ON s.field_id = f.id AND s.deleted_at IS NULL
LEFT JOIN wells w ON w.site_id = s.id AND w.deleted_at IS NULL
LEFT JOIN wellbores wb ON wb.well_id = w.id AND wb.deleted_at IS NULL
LEFT JOIN designs d ON d.wellbore_id = wb.id AND d.deleted_at IS NULL
LEFT JOIN trajectories t ON t.design_id = d.id AND t.deleted_at IS NULL
LEFT JOIN cases k ON k.trajectory_id = t.id AND k.deleted_at IS NULL
WHERE c.organization_id = ? AND c.deleted_at IS NULL
ORDER BY c.name, f.name, s.name, w.name, wb.name, d.plan_name, t.created_at, k.created_at`

type treeRow struct {
	CompanyID, CompanyName, FieldID, FieldName, SiteID, SiteName, WellID, WellName                 *string
	WellboreID, WellboreName, DesignID, DesignName, TrajectoryID, TrajectoryName, CaseID, CaseName *string
}

func (r treeRow) level(i int) (id, name *string) {
	pairs := [][2]*string{{r.CompanyID, r.CompanyName}, {r.FieldID, r.FieldName}, {r.SiteID, r.SiteName}, {r.WellID, r.WellName},
		{r.WellboreID, r.WellboreName}, {r.DesignID, r.DesignName}, {r.TrajectoryID, r.TrajectoryName}, {r.CaseID, r.CaseName}}
	return pairs[i][0], pairs[i][1]
}

// Tree returns the organization's hierarchy (names and ids only) in display order.
func (r *treeRepository) Tree(ctx context.Context, organizationID string) ([]*entities.TreeNode, error) {
	var rows []treeRow
	if err := r.db.WithContext(ctx).Raw(treeQuery, organizationID).Scan(&rows).Error; err != nil {
		return nil, err
	}
	return buildTree(rows), nil
}

// buildTree nests flat join rows; the first occurrence of an id fixes its position.
func buildTree(rows []treeRow) []*entities.TreeNode {
	roots := []*entities.TreeNode{}
	index := map[string]*entities.TreeNode{}
	for _, row := range rows {
		siblings := &roots
		for level := range treeLevels {
			id, name := row.level(level)
			if id == nil {
				break
			}
			node, ok := index[*id]
			if !ok {
				label := ""
				if name != nil {
					label = *name
				}
				node = &entities.TreeNode{ID: *id, Kind: treeLevels[level], Name: label, Children: []*entities.TreeNode{}}
				index[*id] = node
				*siblings = append(*siblings, node)
			}
			siblings = &node.Children
		}
	}
	return roots
}
