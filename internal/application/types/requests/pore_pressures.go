package requests

// CreatePorePressureRequest represents the request body for creating a pore pressure record.
type CreatePorePressureRequest struct {
	CaseID   string  `json:"case_id" binding:"required"`
	TVD      float64 `json:"tvd" binding:"required"`
	Pressure float64 `json:"pressure" binding:"required"`
	EMW      float64 `json:"emw" binding:"required"`
}

// UpdatePorePressureRequest represents the request body for updating a pore pressure record.
type UpdatePorePressureRequest struct {
	ID       string  `json:"id" binding:"required"`
	TVD      float64 `json:"tvd" binding:"required"`
	Pressure float64 `json:"pressure" binding:"required"`
	EMW      float64 `json:"emw" binding:"required"`
}

// GetPorePressuresRequest represents the request for retrieving pore pressures associated with a case.
type GetPorePressuresRequest struct {
	CaseID string `json:"case_id" binding:"required"`
}

// GetPorePressureByIDRequest represents the request for retrieving a pore pressure by ID.
type GetPorePressureByIDRequest struct {
	ID string `json:"id" binding:"required"`
}

// DeletePorePressureRequest represents the request for deleting a pore pressure by ID.
type DeletePorePressureRequest struct {
	ID string `json:"id" binding:"required"`
}
