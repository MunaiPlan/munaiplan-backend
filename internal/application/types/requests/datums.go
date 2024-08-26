package requests

// CreateDatumRequestBody represents the request body for creating a Datum
type CreateDatumRequestBody struct {
	Name              string  `json:"name"`
	SystemDescription string  `json:"system_description"`
	SystemElevation   float64 `json:"system_elevation"`
	DatumDescription  string  `json:"datum_description"`
	WellheadElevation float64 `json:"wellhead_elevation,omitempty"`
	DatumElevation    float64 `json:"datum_elevation,omitempty"`
	AirGap            float64 `json:"air_gap,omitempty"`
	GroundElevation   float64 `json:"ground_elevation,omitempty"`
	Type              string  `json:"type,omitempty"`
}

// CreateDatumRequest represents the request for creating a Datum
type CreateDatumRequest struct {
	Body   CreateDatumRequestBody
	CaseID string
}

// UpdateDatumRequestBody represents the request body for updating a Datum
type UpdateDatumRequestBody struct {
	Name              string  `json:"name"`
	SystemDescription string  `json:"system_description"`
	SystemElevation   float64 `json:"system_elevation"`
	DatumDescription  string  `json:"datum_description"`
	WellheadElevation float64 `json:"wellhead_elevation,omitempty"`
	DatumElevation    float64 `json:"datum_elevation,omitempty"`
	AirGap            float64 `json:"air_gap,omitempty"`
	GroundElevation   float64 `json:"ground_elevation,omitempty"`
	Type              string  `json:"type,omitempty"`
}

// UpdateDatumRequest represents the request for updating a Datum
type UpdateDatumRequest struct {
	ID   string
	Body UpdateDatumRequestBody
}

// GetDatumByIDRequest represents the request for getting a Datum by ID
type GetDatumByIDRequest struct {
	ID string
}

// GetDatumsByCaseIDRequest represents the request for getting Datums by Case ID
type GetDatumsByCaseIDRequest struct {
	CaseID string
}

// DeleteDatumRequest represents the request for deleting a Datum
type DeleteDatumRequest struct {
	ID string
}
