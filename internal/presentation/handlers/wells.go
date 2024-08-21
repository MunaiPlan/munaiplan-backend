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

// initWellsRoutes initializes the routes for the wells API.
func (h *Handler) initWellsRoutes(api *gin.RouterGroup) {
	wells := api.Group("/wells", h.authMiddleware.UserIdentity)
	{
		wells.GET("/", h.getWells)
		wells.POST("/", h.createWell)
		wells.GET("/:id", h.getWellByID)
		wells.PUT("/:id", h.updateWell)
		wells.DELETE("/:id", h.deleteWell)
	}
}

// getWells retrieves all wells.
// @Summary Get Wells
// @Tags wells
// @Description Retrieves all wells
// @Accept json
// @Produce json
// @Param siteId query string true "Site ID"
// @Success 200 {array} entities.Well
// @Failure 500 {object} helpers.Response
// @Router /api/v1/wells [get]
func (h *Handler) getWells(c *gin.Context) {
	var inp requests.GetWellsRequest
	var err error
	if inp.SiteID, err = h.validateQueryIDParam(c, values.SiteIdQueryParam); err != nil {
		return
	}

	wells, err := h.services.Wells.GetWells(c.Request.Context(), &inp)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, wells)
}

// createWell creates a new well.
// @Summary Create Well
// @Tags wells
// @Description Creates a new well
// @Accept json
// @Produce json
// @Param siteId query string true "Site ID"
// @Param input body requests.CreateWellRequest true "Well input"
// @Success 201 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/wells [post]
func (h *Handler) createWell(c *gin.Context) {
	var inp requests.CreateWellRequest
	var err error

	if err = c.BindJSON(&inp.Body); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, types.ErrInvalidInputBody.Error())
		return
	}
	if inp.SiteID, err = h.validateQueryIDParam(c, values.SiteIdQueryParam); err != nil {
		return
	}
	if err = h.services.Wells.CreateWell(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, helpers.NewResponse("well created"))
}

// updateWell updates an existing well.
// @Summary Update Well
// @Tags wells
// @Description Updates an existing well
// @Accept json
// @Produce json
// @Param id path string true "Well ID"
// @Param input body requests.UpdateWellRequest true "Well input"
// @Success 200 {object} entities.Well
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/wells/{id} [put]
func (h *Handler) updateWell(c *gin.Context) {
	var inp requests.UpdateWellRequest
	var err error
	if err := c.BindJSON(&inp.Body); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, types.ErrInvalidInputBody.Error())
		return
	}
	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}

	well, err := h.services.Wells.UpdateWell(c.Request.Context(), &inp)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, well)
}

// deleteWell deletes an existing well.
// @Summary Delete Well
// @Tags wells
// @Description Deletes an existing well
// @Accept json
// @Produce json
// @Param id path string true "Well ID"
// @Success 200 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/wells/{id} [delete]
func (h *Handler) deleteWell(c *gin.Context) {
	var inp requests.DeleteWellRequest
	var err error

	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}
	if err = h.services.Wells.DeleteWell(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, helpers.NewResponse("well deleted"))
}

// getWellByID retrieves a well by its ID.
// @Summary Get Well by ID
// @Tags wells
// @Description Retrieves a well by its ID
// @Accept json
// @Produce json
// @Param id path string true "Well ID"
// @Success 200 {object} entities.Well
// @Failure 500 {object} helpers.Response
// @Router /api/v1/wells/{id} [get]
func (h *Handler) getWellByID(c *gin.Context) {
	var inp requests.GetWellByIDRequest
	var err error
	var well *entities.Well
	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}
	if well, err = h.services.Wells.GetWellByID(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, well)
}
