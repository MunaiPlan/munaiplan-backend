package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/munaiplan/munaiplan-backend/internal/application/types/requests"
	"github.com/munaiplan/munaiplan-backend/internal/helpers"
	"github.com/munaiplan/munaiplan-backend/internal/presentation/types"
	"github.com/munaiplan/munaiplan-backend/pkg/values"
)

// initCompaniesRoutes initializes the routes for the companies API.
func (h *Handler) initCompaniesRoutes(api *gin.RouterGroup) {
	companies := api.Group("/companies", h.authMiddleware.UserIdentity)
	{
		companies.GET("/", h.getCompanies)
		companies.POST("/", h.createCompany)
		companies.GET("/:id", h.getCompanyByID)
		companies.PUT("/:id", h.updateCompany)
		companies.DELETE("/:id", h.deleteCompany)
	}
}

// getCompanies retrieves all companies.
// @Summary Get Companies
// @Tags companies
// @Description Retrieves all companies
// @Accept json
// @Produce json
// @Param organizationId query string true "Organization ID"
// @Success 200 {array} entities.Company
// @Failure 500 {object} helpers.Response
// @Router /api/v1/companies [get]
func (h *Handler) getCompanies(c *gin.Context) {
	var inp requests.GetCompaniesRequest
	var err error
	inp.OrganizationID, err = h.validateContextIDKey(c, values.OrganizationIdCtx)
	if err != nil {
		return
	}
	companies, err := h.services.Companies.GetCompanies(c.Request.Context(), &inp)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, companies)
}

// createCompany creates a new company.
// @Summary Create Company
// @Tags companies
// @Description Creates a new company
// @Accept json
// @Produce json
// @Param organizationId query string true "Organization ID"
// @Param input body requests.CreateCompanyRequest true "Company input"
// @Success 201 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/companies [post]
func (h *Handler) createCompany(c *gin.Context) {
	var inp requests.CreateCompanyRequest
	var err error
	if err = c.BindJSON(&inp.Body); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, types.ErrInvalidInputBody.Error())
		return
	}
	if inp.OrganizationID, err = h.validateContextIDKey(c, values.OrganizationIdQueryParam); err != nil {
		return
	}
	if err = h.services.Companies.CreateCompany(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, helpers.NewResponse("company created"))
}

// updateCompany updates an existing company.
// @Summary Update Company
// @Tags companies
// @Description Updates an existing company
// @Accept json
// @Produce json
// @Param id path string true "Company ID"
// @Param organizationId query string true "Organization ID"
// @Param input body requests.UpdateCompanyRequest true "Company input"
// @Success 200 {object} entities.Company
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/companies/{id} [put]
func (h *Handler) updateCompany(c *gin.Context) {
	var inp requests.UpdateCompanyRequest
	var err error
	if err := c.BindJSON(&inp.Body); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, types.ErrInvalidInputBody.Error())
		return
	}
	if inp.OrganizationID, err = h.validateContextIDKey(c, values.OrganizationIdQueryParam); err != nil {
		return
	}
	if inp.Body.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}

	company, err := h.services.Companies.UpdateCompany(c.Request.Context(), &inp)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, company)
}

// deleteCompany deletes an existing company.
// @Summary Delete Company
// @Tags companies
// @Description Deletes an existing company
// @Accept json
// @Produce json
// @Param organizationId query string true "Organization ID"
// @Param id path string true "Company ID"
// @Success 200 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/companies/{id} [delete]
func (h *Handler) deleteCompany(c *gin.Context) {
	var inp requests.DeleteCompanyRequest
	var err error

	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}
	if inp.OrganizationID, err = h.validateContextIDKey(c, values.OrganizationIdQueryParam); err != nil {
		return
	}

	if err = h.services.Companies.DeleteCompany(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, helpers.NewResponse("company deleted"))
}

// getCompanyByName retrieves a company by its name.
// @Summary Get Company by Name
// @Tags companies
// @Description Retrieves a company by its name
// @Accept json
// @Produce json
// @Param organizationId query string true "Organization ID"
// @Param companyName path string true "Company Name"
// @Success 200 {object} entities.Company
// @Failure 500 {object} helpers.Response
// @Router /api/v1/companies/{name} [get]
func (h *Handler) getCompanyByID(c *gin.Context) {
	var inp requests.GetCompanyByIDRequest
	var err error
	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}
	fmt.Println("inp.ID ", inp.ID)
	if inp.OrganizationID, err = h.validateContextIDKey(c, values.OrganizationIdQueryParam); err != nil {
		return
	}

	company, err := h.services.Companies.GetCompanyByID(c.Request.Context(), &inp)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, company)
}
