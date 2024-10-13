package entities

import "time"

// String represents the domain entity for a String.
type String struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Depth     float64    `json:"depth"`
	CreatedAt time.Time  `json:"created_at"`
	Sections  []*Section `json:"sections"`
}

// Section represents the domain entity for a Section associated with a String.
type Section struct {
	ID                  string    `json:"id"`
	Description         *string   `json:"description,omitempty"`
	Manufacturer        *string   `json:"manufacturer,omitempty"`
	Type                string    `json:"type"`
	BodyMD              float64   `json:"body_md"`
	BodyLength          float64   `json:"body_length"`
	BodyOD              float64   `json:"body_od"`
	BodyID              float64   `json:"body_id"`
	AvgJointLength      *float64  `json:"avg_joint_length,omitempty"`
	StabilizerLength    *float64  `json:"stabilizer_length,omitempty"`
	StabilizerOD        *float64  `json:"stabilizer_od,omitempty"`
	StabilizerID        *float64  `json:"stabilizer_id,omitempty"`
	Weight              *float64  `json:"weight,omitempty"`
	Material            *string   `json:"material,omitempty"`
	Grade               *string   `json:"grade,omitempty"`
	Class               *int      `json:"class,omitempty"`
	FrictionCoefficient *float64  `json:"friction_coefficient,omitempty"`
	MinYieldStrength    *float64  `json:"min_yield_strength,omitempty"`
	CreatedAt           time.Time `json:"created_at"`
}
