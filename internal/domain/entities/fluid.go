package entities

import "time"

// Fluid entity
type Fluid struct {
	ID              string    `json:"id"`
	CaseID          string    `json:"case_id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Density         float64   `json:"density"`
	FluidBaseTypeID string    `json:"fluid_base_type_id"`
	BaseFluidID     string    `json:"base_fluid_id"`
	CreatedAt       time.Time `json:"created_at"`
}

// FluidType entity
type FluidType struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
