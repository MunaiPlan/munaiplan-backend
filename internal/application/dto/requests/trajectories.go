package requests

// CreateTrajectoryHeaderRequestBody represents the request body for updating a trajectory header
type CreateTrajectoryHeaderRequestBody struct {
	Customer         string  `json:"customer"`
	Project          string  `json:"project"`
	ProfileType      string  `json:"profile_type"`
	Field            string  `json:"field"`
	YourRef          string  `json:"your_ref"`
	Structure        string  `json:"structure"`
	JobNumber        string  `json:"job_number"`
	Wellhead         string  `json:"wellhead"`
	KellyBushingElev float64 `json:"kelly_bushing_elev"`
	Profile          string  `json:"profile"`
}

// CreateTrajectoryUnitRequestBody represents the request body for creating or updating a trajectory unit
type CreateTrajectoryUnitRequestBody struct {
	MD              float64 `json:"md"`
	Incl            float64 `json:"incl"`
	Azim            float64 `json:"azim"`
	SubSea          float64 `json:"sub_sea"`
	TVD             float64 `json:"tvd"`
	LocalNCoord     float64 `json:"local_n_coord"`
	LocalECoord     float64 `json:"local_e_coord"`
	GlobalNCoord    float64 `json:"global_n_coord"`
	GlobalECoord    float64 `json:"global_e_coord"`
	Dogleg          float64 `json:"dogleg"`
	VerticalSection float64 `json:"vertical_section"`
}

// UpdateTrajectoryHeaderRequestBody represents the request body for updating a trajectory header
type UpdateTrajectoryHeaderRequestBody struct {
	ID               string  `json:"id"`
	Customer         string  `json:"customer"`
	Project          string  `json:"project"`
	ProfileType      string  `json:"profile_type"`
	Field            string  `json:"field"`
	YourRef          string  `json:"your_ref"`
	Structure        string  `json:"structure"`
	JobNumber        string  `json:"job_number"`
	Wellhead         string  `json:"wellhead"`
	KellyBushingElev float64 `json:"kelly_bushing_elev"`
	Profile          string  `json:"profile"`
}

// UpdateTrajectoryUnitRequestBody represents the request body for updating a trajectory unit
type UpdateTrajectoryUnitRequestBody struct {
	ID              string  `json:"id"`
	MD              float64 `json:"md"`
	Incl            float64 `json:"incl"`
	Azim            float64 `json:"azim"`
	SubSea          float64 `json:"sub_sea"`
	TVD             float64 `json:"tvd"`
	LocalNCoord     float64 `json:"local_n_coord"`
	LocalECoord     float64 `json:"local_e_coord"`
	GlobalNCoord    float64 `json:"global_n_coord"`
	GlobalECoord    float64 `json:"global_e_coord"`
	Dogleg          float64 `json:"dogleg"`
	VerticalSection float64 `json:"vertical_section"`
}

// CreateTrajectoryRequestBody represents the request body for creating a trajectory
type CreateTrajectoryRequestBody struct {
	Name        string                              `json:"name"`
	Description string                              `json:"description"`
	Headers     []CreateTrajectoryHeaderRequestBody `json:"headers"`
	Units       []CreateTrajectoryUnitRequestBody   `json:"units"`
}

// UpdateTrajectoryRequestBody represents the request body for creating a trajectory
type UpdateTrajectoryRequestBody struct {
	Name        string                              `json:"name"`
	Description string                              `json:"description"`
	Headers     []UpdateTrajectoryHeaderRequestBody `json:"headers"`
	Units       []UpdateTrajectoryUnitRequestBody   `json:"units"`
}

// CreateTrajectoryRequest represents the request for creating a trajectory
type CreateTrajectoryRequest struct {
	Body     CreateTrajectoryRequestBody
	DesignID string
}

// UpdateTrajectoryRequest represents the request for updating a trajectory
type UpdateTrajectoryRequest struct {
	ID   string
	Body UpdateTrajectoryRequestBody // The structure is identical to the create request
}

// GetTrajectoriesRequest represents the request for getting trajectories
type GetTrajectoriesRequest struct {
	DesignID string
}

// GetTrajectoryByIDRequest represents the request for getting a trajectory by ID
type GetTrajectoryByIDRequest struct {
	ID string
}

// DeleteTrajectoryRequest represents the request for deleting a trajectory
type DeleteTrajectoryRequest struct {
	ID string
}
