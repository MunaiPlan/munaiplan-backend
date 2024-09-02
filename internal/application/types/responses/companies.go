package responses

type CompanyResponse struct {
	ID     string           `json:"id"`
	Name   string           `json:"name"`
	Fields []FieldResponse  `json:"fields"`
}

type FieldResponse struct {
	ID    string         `json:"id"`
	Name  string         `json:"name"`
	Sites []SiteResponse `json:"sites"`
}

type SiteResponse struct {
	ID    string         `json:"id"`
	Name  string         `json:"name"`
	Wells []WellResponse `json:"wells"`
}

type WellResponse struct {
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	Wellbores []WellboreResponse `json:"wellbores"`
}

type WellboreResponse struct {
	ID    string       `json:"id"`
	Name  string       `json:"name"`
	Desings []DesignResponse `json:"designs"`
}

type DesignResponse struct {
	ID    string       `json:"id"`
	Name  string       `json:"name"`
}
