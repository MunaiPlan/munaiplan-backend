package entities

import "time"

// ImportResult identifies what an import created or reused in the well hierarchy.
type ImportResult struct {
	ImportID     string          `json:"import_id"`
	CompanyID    string          `json:"company_id"`
	FieldID      string          `json:"field_id"`
	SiteID       string          `json:"site_id"`
	WellID       string          `json:"well_id"`
	WellboreID   string          `json:"wellbore_id"`
	DesignID     string          `json:"design_id"`
	TrajectoryID string          `json:"trajectory_id"`
	CaseID       string          `json:"case_id"`
	Created      map[string]bool `json:"created"` // level -> true when newly created, false when reused
	CreatedAt    time.Time       `json:"created_at"`
}
