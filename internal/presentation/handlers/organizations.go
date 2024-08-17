package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/munaiplan/munaiplan-backend/internal/application/dto/requests"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/helpers"
	"github.com/munaiplan/munaiplan-backend/internal/presentation/types"
	"github.com/munaiplan/munaiplan-backend/pkg/values"
)

// initOrganizationsRoutes initializes the routes for the organizations API.
func (h *Handler) initOrganizationsRoutes(api *gin.RouterGroup) {
	organizations := api.Group("/organizations")
	{
		organizations.GET("/", h.getOrganizations)
		organizations.POST("/", h.createOrganization)
		organizations.GET("/:name", h.getOrganizationByName)
		organizations.PUT("/:id", h.updateOrganization)
		organizations.DELETE("/:id", h.deleteOrganization)
	}
}

// @Summary Get Organizations
// @Tags organizations
// @Description Retrieve all organizations
// @Accept json
// @Produce json
// @Success 200 {array} entities.Organization
// @Failure 500 {object} helpers.Response
// @Router /api/v1/organizations [get]
func (h *Handler) getOrganizations(c *gin.Context) {
	organizations, err := h.services.Organizations.GetOrganizations(c.Request.Context())
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, types.ErrOrganizationsNotFound.Error())
		return
	}

	c.JSON(http.StatusOK, organizations)
}

// @Summary Create Organization
// @Tags organizations
// @Description Create a new organization
// @Accept json
// @Produce json
// @Param input body requests.CreateOrganizationRequest true "Organization details"
// @Success 201 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/organizations [post]
func (h *Handler) createOrganization(c *gin.Context) {
	var inp requests.CreateOrganizationRequest
	var err error

	if err = c.BindJSON(&inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	if err = h.services.Organizations.CreateOrganization(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, helpers.NewResponse("organization created"))
}

// @Summary Update Organization
// @Tags organizations
// @Description Update an existing organization
// @Accept json
// @Produce json
// @Param id path string true "Organization ID"
// @Param input body requests.UpdateOrganizationRequest true "Updated organization details"
// @Success 200 {object} entities.Organization
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/organizations/{id} [put]
func (h *Handler) updateOrganization(c *gin.Context) {
	var inp requests.UpdateOrganizationRequest
	var err error
	var organization *entities.Organization
	if err = c.BindJSON(&inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}
	if organization, err = h.services.Organizations.UpdateOrganization(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, organization)
}

// @Summary Delete Organization
// @Tags organizations
// @Description Delete an organization
// @Accept json
// @Produce json
// @Param id path string true "Organization ID"
// @Success 200 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/organizations/{id} [delete]
func (h *Handler) deleteOrganization(c *gin.Context) {
	var inp requests.DeleteOrganizationRequest
	var err error

	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}
	if err = h.services.Organizations.DeleteOrganization(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, helpers.NewResponse("organization deleted"))
}

// @Summary Get Organization by Name
// @Tags organizations
// @Description Retrieve an organization by its name
// @Accept json
// @Produce json
// @Param name path string true "Organization name"
// @Success 200 {object} entities.Organization
// @Failure 500 {object} helpers.Response
// @Router /api/v1/organizations/{name} [get]
func (h *Handler) getOrganizationByName(c *gin.Context) {
	var inp requests.GetOrganizationByNameRequest
	var err error
	var organization *entities.Organization
	if inp.Name, err = h.validateRequestParam(c, values.NameQueryParam); err != nil {
		return
	}
	if organization, err = h.services.Organizations.GetOrganizationByName(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, organization)
}
