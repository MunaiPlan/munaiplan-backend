package requests

// CreateWellboreRequestBody represents the request body for creating a wellbore
type CreateWellboreRequestBody struct {
	Name                           string  `json:"name"`
	BottomHoleLocation             string  `json:"bottom_hole_location"`
	WellboreDepth                  float64 `json:"wellbore_depth"`
	AverageHookLoad                float64 `json:"average_hook_load"`
	RiserPressure                  float64 `json:"riser_pressure"`
	AverageInletFlow               float64 `json:"average_inlet_flow"`
	AverageColumnRotationFrequency float64 `json:"average_column_rotation_frequency"`
	MaximumColumnRotationFrequency float64 `json:"maximum_column_rotation_frequency"`
	AverageWeightOnBit             float64 `json:"average_weight_on_bit"`
	MaximumWeightOnBit             float64 `json:"maximum_weight_on_bit"`
	AverageTorque                  float64 `json:"average_torque"`
	MaximumTorque                  float64 `json:"maximum_torque"`
	DownStaticFriction             float64 `json:"down_static_friction"`
	DepthInterval                  float64 `json:"depth_interval"`
}

// CreateWellboreRequest represents the request for creating a wellbore
type CreateWellboreRequest struct {
	Body   CreateWellboreRequestBody
	WellID string
}

// UpdateWellboreRequestBody represents the request body for updating a wellbore
type UpdateWellboreRequestBody struct {
	Name                           string  `json:"name"`
	BottomHoleLocation             string  `json:"bottom_hole_location"`
	WellboreDepth                  float64 `json:"wellbore_depth"`
	AverageHookLoad                float64 `json:"average_hook_load"`
	RiserPressure                  float64 `json:"riser_pressure"`
	AverageInletFlow               float64 `json:"average_inlet_flow"`
	AverageColumnRotationFrequency float64 `json:"average_column_rotation_frequency"`
	MaximumColumnRotationFrequency float64 `json:"maximum_column_rotation_frequency"`
	AverageWeightOnBit             float64 `json:"average_weight_on_bit"`
	MaximumWeightOnBit             float64 `json:"maximum_weight_on_bit"`
	AverageTorque                  float64 `json:"average_torque"`
	MaximumTorque                  float64 `json:"maximum_torque"`
	DownStaticFriction             float64 `json:"down_static_friction"`
	DepthInterval                  float64 `json:"depth_interval"`
}

// UpdateWellboreRequest represents the request for updating a wellbore
type UpdateWellboreRequest struct {
	ID   string
	Body UpdateWellboreRequestBody
}

// GetWellboresRequest represents the request for getting wellbores
type GetWellboresRequest struct {
	WellID string
}

// GetWellboreByIDRequest represents the request for getting a wellbore by ID
type GetWellboreByIDRequest struct {
	ID string
}

// DeleteWellboreRequest represents the request for deleting a wellbore
type DeleteWellboreRequest struct {
	ID string
}
