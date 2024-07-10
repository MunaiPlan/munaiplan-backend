package domain

import (
	"time"
)

// Скважина (под кустом)
// TODO() Correct all tables in this file
type Well struct {
	ID                      string      `json:"id"`
	Name                    string      `json:"name"`
	Description             string      `json:"description"`
	Location                string      `json:"location"`
	UniversalWellIdentifier string      `json:"universal_well_identifier"`
	Type                    string      `json:"type"`
	WellNumber              string      `json:"well_number"`
	WorkingGroup            string      `json:"working_group"`
	ActiveWellUnit          string      `json:"active_well_unit"`
	Wellbores               []*Wellbore `json:"wellbores"`
	CreatedAt               time.Time   `json:"created_at"`
}
