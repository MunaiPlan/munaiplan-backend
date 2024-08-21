package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/munaiplan/munaiplan-backend/internal/application/types/requests"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/helpers"
	"github.com/munaiplan/munaiplan-backend/internal/presentation/types"
	"github.com/munaiplan/munaiplan-backend/pkg/values"
)

// initSitesRoutes initializes the routes for the sites API.
func (h *Handler) initSitesRoutes(api *gin.RouterGroup) {
	sites := api.Group("/sites", h.authMiddleware.UserIdentity)
	{
		sites.GET("/", h.getSites)
		sites.POST("/", h.createSite)
		sites.GET("/:id", h.getSiteByID)
		sites.PUT("/:id", h.updateSite)
		sites.DELETE("/:id", h.deleteSite)
	}
}

// getSites retrieves all sites.
// @Summary Get Sites
// @Tags sites
// @Description Retrieves all sites
// @Accept json
// @Produce json
// @Param fieldId query string true "Field ID"
// @Success 200 {array} entities.Site
// @Failure 500 {object} helpers.Response
// @Router /api/v1/sites [get]
func (h *Handler) getSites(c *gin.Context) {
	var inp requests.GetSitesRequest
	var err error
	if inp.FieldID, err = h.validateQueryIDParam(c, values.FieldIdQueryParam); err != nil {
		return
	}

	sites, err := h.services.Sites.GetSites(c.Request.Context(), &inp)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, sites)
}

// createSite creates a new site.
// @Summary Create Site
// @Tags sites
// @Description Creates a new site
// @Accept json
// @Produce json
// @Param fieldId query string true "Field ID"
// @Param input body requests.CreateSiteRequest true "Site input"
// @Success 201 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/sites [post]
func (h *Handler) createSite(c *gin.Context) {
	var inp requests.CreateSiteRequest
	var err error

	if err = c.BindJSON(&inp.Body); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, types.ErrInvalidInputBody.Error())
		return
	}
	if inp.FieldID, err = h.validateQueryIDParam(c, values.FieldIdQueryParam); err != nil {
		return
	}
	if err = h.services.Sites.CreateSite(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, helpers.NewResponse("site created"))
}

// updateSite updates an existing site.
// @Summary Update Site
// @Tags sites
// @Description Updates an existing site
// @Accept json
// @Produce json
// @Param id path string true "Site ID"
// @Param input body requests.UpdateSiteRequest true "Site input"
// @Success 200 {object} entities.Site
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/sites/{id} [put]
func (h *Handler) updateSite(c *gin.Context) {
	var inp requests.UpdateSiteRequest
	var err error
	if err := c.BindJSON(&inp.Body); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, types.ErrInvalidInputBody.Error())
		return
	}
	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}

	site, err := h.services.Sites.UpdateSite(c.Request.Context(), &inp)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, site)
}

// deleteSite deletes an existing site.
// @Summary Delete Site
// @Tags sites
// @Description Deletes an existing site
// @Accept json
// @Produce json
// @Param id path string true "Site ID"
// @Success 200 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/sites/{id} [delete]
func (h *Handler) deleteSite(c *gin.Context) {
	var inp requests.DeleteSiteRequest
	var err error

	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}
	if err = h.services.Sites.DeleteSite(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, helpers.NewResponse("site deleted"))
}

// getSiteByID retrieves a site by its ID.
// @Summary Get Site by ID
// @Tags sites
// @Description Retrieves a site by its ID
// @Accept json
// @Produce json
// @Param id path string true "Site ID"
// @Success 200 {object} entities.Site
// @Failure 500 {object} helpers.Response
// @Router /api/v1/sites/{id} [get]
func (h *Handler) getSiteByID(c *gin.Context) {
	var inp requests.GetSiteByIDRequest
	var err error
	var site *entities.Site
	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}
	if site, err = h.services.Sites.GetSiteByID(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, site)
}
