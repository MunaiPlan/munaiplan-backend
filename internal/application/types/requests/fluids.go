package requests

// CreateFluidRequestBody represents the request body for creating a fluid.
type CreateFluidRequestBody struct {
	Name            string  `json:"name" binding:"required"`
	Description     string  `json:"description"`
	Density         float64 `json:"density" binding:"required"`
	FluidBaseTypeID string  `json:"fluid_base_type_id" binding:"required"`
	BaseFluidID     string  `json:"base_fluid_id" binding:"required"`
}

// UpdateFluidRequestBody represents the request body for updating a fluid.
type UpdateFluidRequestBody struct {
	ID              string  `json:"id" binding:"required"`
	Name            string  `json:"name" binding:"required"`
	Description     string  `json:"description"`
	Density         float64 `json:"density" binding:"required"`
	FluidBaseTypeID string  `json:"fluid_base_type_id" binding:"required"`
	BaseFluidID     string  `json:"base_fluid_id" binding:"required"`
}

// CreateFluidRequest represents the request for creating a fluid.
type CreateFluidRequest struct {
	Body   CreateFluidRequestBody `json:"body" binding:"required"`
	CaseID string                 `json:"case_id" binding:"required"`
}

// UpdateFluidRequest represents the request for updating a fluid.
type UpdateFluidRequest struct {
	ID   string               `json:"id" binding:"required"`
	Body UpdateFluidRequestBody `json:"body" binding:"required"`
}

// GetFluidsRequest represents the request for retrieving all fluids by case ID.
type GetFluidsRequest struct {
	CaseID string `json:"case_id" binding:"required"`
}

// GetFluidByIDRequest represents the request for retrieving a single fluid by ID.
type GetFluidByIDRequest struct {
	ID string `json:"id" binding:"required"`
}

// DeleteFluidRequest represents the request for deleting a fluid.
type DeleteFluidRequest struct {
	ID string `json:"id" binding:"required"`
}
