package requests

// CreateRigRequestBody represents the request body for creating a rig.
type CreateRigRequestBody struct {
	BlockRating                     *float64 `json:"block_rating,omitempty"`
	TorqueRating                    *float64 `json:"torque_rating,omitempty"`
	RatedWorkingPressure            float64  `json:"rated_working_pressure"`
	BopPressureRating               float64  `json:"bop_pressure_rating"`
	SurfacePressureLoss             float64  `json:"surface_pressure_loss"`
	StandpipeLength                 *float64 `json:"standpipe_length,omitempty"`
	StandpipeInternalDiameter       *float64 `json:"standpipe_internal_diameter,omitempty"`
	HoseLength                      *float64 `json:"hose_length,omitempty"`
	HoseInternalDiameter            *float64 `json:"hose_internal_diameter,omitempty"`
	SwivelLength                    *float64 `json:"swivel_length,omitempty"`
	SwivelInternalDiameter          *float64 `json:"swivel_internal_diameter,omitempty"`
	KellyLength                     *float64 `json:"kelly_length,omitempty"`
	KellyInternalDiameter           *float64 `json:"kelly_internal_diameter,omitempty"`
	PumpDischargeLineLength         *float64 `json:"pump_discharge_line_length,omitempty"`
	PumpDischargeLineInternalDiameter *float64 `json:"pump_discharge_line_internal_diameter,omitempty"`
	TopDriveStackupLength           *float64 `json:"top_drive_stackup_length,omitempty"`
	TopDriveStackupInternalDiameter *float64 `json:"top_drive_stackup_internal_diameter,omitempty"`
}

// UpdateRigRequestBody represents the request body for updating a rig.
type UpdateRigRequestBody struct {
	ID                              string   `json:"id"`
	BlockRating                     *float64 `json:"block_rating,omitempty"`
	TorqueRating                    *float64 `json:"torque_rating,omitempty"`
	RatedWorkingPressure            float64  `json:"rated_working_pressure"`
	BopPressureRating               float64  `json:"bop_pressure_rating"`
	SurfacePressureLoss             float64  `json:"surface_pressure_loss"`
	StandpipeLength                 *float64 `json:"standpipe_length,omitempty"`
	StandpipeInternalDiameter       *float64 `json:"standpipe_internal_diameter,omitempty"`
	HoseLength                      *float64 `json:"hose_length,omitempty"`
	HoseInternalDiameter            *float64 `json:"hose_internal_diameter,omitempty"`
	SwivelLength                    *float64 `json:"swivel_length,omitempty"`
	SwivelInternalDiameter          *float64 `json:"swivel_internal_diameter,omitempty"`
	KellyLength                     *float64 `json:"kelly_length,omitempty"`
	KellyInternalDiameter           *float64 `json:"kelly_internal_diameter,omitempty"`
	PumpDischargeLineLength         *float64 `json:"pump_discharge_line_length,omitempty"`
	PumpDischargeLineInternalDiameter *float64 `json:"pump_discharge_line_internal_diameter,omitempty"`
	TopDriveStackupLength           *float64 `json:"top_drive_stackup_length,omitempty"`
	TopDriveStackupInternalDiameter *float64 `json:"top_drive_stackup_internal_diameter,omitempty"`
}

// CreateRigRequest represents the full request for creating a rig.
type CreateRigRequest struct {
	Body   CreateRigRequestBody
	CaseID string
}

// UpdateRigRequest represents the full request for updating a rig.
type UpdateRigRequest struct {
	ID    string
	Body  UpdateRigRequestBody
}

// GetRigsRequest represents the request for retrieving all rigs for a case.
type GetRigsRequest struct {
	CaseID string
}

// GetRigByIDRequest represents the request for retrieving a specific rig by ID.
type GetRigByIDRequest struct {
	ID string
}

// DeleteRigRequest represents the request for deleting a specific rig by ID.
type DeleteRigRequest struct {
	ID string
}
