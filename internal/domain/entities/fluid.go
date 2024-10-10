package entities

import "time"

// Fluid entity
type Fluid struct {
	ID            string     `json:"id"`
	Name          string     `json:"name"`
	Description   string     `json:"description"`
	Density       float64    `json:"density"`
	FluidBaseType *FluidType `json:"fluid_base_type"`
	BaseFluid     *FluidType `json:"base_fluid"`
	CreatedAt     time.Time  `json:"created_at"`
}

// FluidType entity
type FluidType struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
