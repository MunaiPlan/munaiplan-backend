package requests

// CreateWellRequestBody represents the request body for creating a well
type CreateWellRequestBody struct {
	Name                    string `json:"name"`
	Description             string `json:"description"`
	Location                string `json:"location"`
	UniversalWellIdentifier string `json:"universal_well_identifier"`
	Type                    string `json:"type"`
	WellNumber              string `json:"well_number"`
	WorkingGroup            string `json:"working_group"`
	ActiveWellUnit          string `json:"active_well_unit"`
}

// CreateWellRequest represents the request for creating a well
type CreateWellRequest struct {
	Body   CreateWellRequestBody
	SiteID string
}

// UpdateWellRequestBody represents the request body for updating a well
type UpdateWellRequestBody struct {
	Name                    string `json:"name"`
	Description             string `json:"description"`
	Location                string `json:"location"`
	UniversalWellIdentifier string `json:"universal_well_identifier"`
	Type                    string `json:"type"`
	WellNumber              string `json:"well_number"`
	WorkingGroup            string `json:"working_group"`
	ActiveWellUnit          string `json:"active_well_unit"`
}

// UpdateWellRequest represents the request for updating a well
type UpdateWellRequest struct {
	ID   string
	Body UpdateWellRequestBody
}

// GetWellsRequest represents the request for getting wells
type GetWellsRequest struct {
	SiteID string
}

// GetWellByIDRequest represents the request for getting a well by ID
type GetWellByIDRequest struct {
	ID string
}

// DeleteWellRequest represents the request for deleting a well
type DeleteWellRequest struct {
	ID string
}
