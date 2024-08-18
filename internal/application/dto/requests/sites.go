package requests

// CreateSiteRequestBody represents the request body for creating a site
type CreateSiteRequestBody struct {
	Name    string  `json:"name"`
	Area    float64 `json:"area"`
	Block   string  `json:"block"`
	Azimuth float64 `json:"azimuth"`
	Country string  `json:"country"`
	State   string  `json:"state"`
	Region  string  `json:"region"`
}

// CreateSiteRequest represents the request for creating a site
type CreateSiteRequest struct {
	Body    CreateSiteRequestBody
	FieldID string
}

// UpdateSiteRequestBody represents the request body for updating a site
type UpdateSiteRequestBody struct {
	Name    string  `json:"name"`
	Area    float64 `json:"area"`
	Block   string  `json:"block"`
	Azimuth float64 `json:"azimuth"`
	Country string  `json:"country"`
	State   string  `json:"state"`
	Region  string  `json:"region"`
}

// UpdateSiteRequest represents the request for updating a site
type UpdateSiteRequest struct {
	ID   string
	Body UpdateSiteRequestBody
}

// GetSitesRequest represents the request for getting sites
type GetSitesRequest struct {
	FieldID string
}

// GetSiteByIDRequest represents the request for getting a site by ID
type GetSiteByIDRequest struct {
	ID string
}

// DeleteSiteRequest represents the request for deleting a site
type DeleteSiteRequest struct {
	ID string
}
