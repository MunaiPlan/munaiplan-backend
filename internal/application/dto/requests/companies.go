package requests

type CreateCompanyRequestBody struct {
	//OrganzationID string `json:"organization_id"`
	Name           string `json:"name"`
	Division       string `json:"division"`
	Group          string `json:"group"`
	Representative string `json:"representative"`
	Address        string `json:"address"`
	Phone          string `json:"phone"`
}

type CreateCompanyRequest struct {
	OrganizationID string
	Body           CreateCompanyRequestBody
}

type UpdateCompanyRequest struct {
	OrganizationID string                   `json:"organization_id"`
	Body           UpdateCompanyRequestBody `json:"body"`
}

type UpdateCompanyRequestBody struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Division       string `json:"division"`
	Group          string `json:"group"`
	Representative string `json:"representative"`
	Address        string `json:"address"`
	Phone          string `json:"phone"`
}

type GetCompaniesRequest struct {
	OrganizationID string `json:"organization_id"`
}

type GetCompanyByNameRequest struct {
	OrganizationID string `json:"organization_id"`
	Name           string `json:"name"`
}

type GetCompanyByID struct {
	OrganizationID string `json:"organization_id"`
	ID             string `json:"id"`
}

type DeleteCompanyRequest struct {
	OrganizationID string `json:"organization_id"`
	ID             string `json:"id"`
}
