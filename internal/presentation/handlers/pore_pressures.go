package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/munaiplan/munaiplan-backend/internal/application/types/requests"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/helpers"
	"github.com/munaiplan/munaiplan-backend/pkg/values"
)

// initPorePressureRoutes initializes the routes for the pore pressure API.
func (h *Handler) initPorePressureRoutes(api *gin.RouterGroup) {
	porePressures := api.Group("/pore-pressures", h.authMiddleware.UserIdentity)
	{
		porePressures.GET("/", h.getPorePressures)
		porePressures.POST("/", h.createPorePressure)
		porePressures.GET("/:id", h.getPorePressureByID)
		porePressures.PUT("/:id", h.updatePorePressure)
		porePressures.DELETE("/:id", h.deletePorePressure)
	}
}

// getPorePressures retrieves all pore pressures for a case.
// @Summary Get Pore Pressures
// @Tags pore-pressures
// @Description Retrieves all pore pressures for a case
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param caseId query string true "Case ID"
// @Success 200 {array} entities.PorePressure
// @Failure 500 {object} helpers.Response
// @Router /api/v1/pore-pressures [get]
func (h *Handler) getPorePressures(c *gin.Context) {
	var inp requests.GetPorePressuresRequest
	var err error
	var porePressures []*entities.PorePressure

	if inp.CaseID, err = h.validateQueryIDParam(c, values.CaseIdQueryParam); err != nil {
		return
	}

	if porePressures, err = h.services.PorePressures.GetPorePressures(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, porePressures)
}

// createPorePressure creates a new pore pressure record.
// @Summary Create Pore Pressure
// @Tags pore-pressures
// @Description Creates a new pore pressure record
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param input body requests.CreatePorePressureRequest true "Pore Pressure input"
// @Success 201 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/pore-pressures [post]
func (h *Handler) createPorePressure(c *gin.Context) {
	var inp requests.CreatePorePressureRequest
	var err error
	if err = c.BindJSON(&inp.Body); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	if inp.CaseID, err = h.validateQueryIDParam(c, values.CaseIdQueryParam); err != nil {
		return
	}

	if err = h.services.PorePressures.CreatePorePressure(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, helpers.NewResponse("Pore pressure created successfully"))
}

// updatePorePressure updates an existing pore pressure record.
// @Summary Update Pore Pressure
// @Tags pore-pressures
// @Description Updates an existing pore pressure record
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Pore Pressure ID"
// @Param input body requests.UpdatePorePressureRequest true "Pore Pressure input"
// @Success 200 {object} entities.PorePressure
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/pore-pressures/{id} [put]
func (h *Handler) updatePorePressure(c *gin.Context) {
	var inp requests.UpdatePorePressureRequest
	var err error

	if err = c.BindJSON(&inp.Body); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, "Invalid request body")
		return
	}
	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}

	updatedPorePressure, err := h.services.PorePressures.UpdatePorePressure(c.Request.Context(), &inp)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, updatedPorePressure)
}

// deletePorePressure deletes a pore pressure record.
// @Summary Delete Pore Pressure
// @Tags pore-pressures
// @Description Deletes a pore pressure record
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Pore Pressure ID"
// @Success 200 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/pore-pressures/{id} [delete]
func (h *Handler) deletePorePressure(c *gin.Context) {
	var inp requests.DeletePorePressureRequest
	var err error

	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}
	if err := h.services.PorePressures.DeletePorePressure(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, helpers.NewResponse("Pore pressure deleted successfully"))
}

// getPorePressureByID retrieves a pore pressure record by its ID.
// @Summary Get Pore Pressure by ID
// @Tags pore-pressures
// @Description Retrieves a pore pressure record by its ID
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Pore Pressure ID"
// @Success 200 {object} entities.PorePressure
// @Failure 500 {object} helpers.Response
// @Router /api/v1/pore-pressures/{id} [get]
func (h *Handler) getPorePressureByID(c *gin.Context) {
	var inp requests.GetPorePressureByIDRequest
	var err error

	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}

	porePressure, err := h.services.PorePressures.GetPorePressureByID(c.Request.Context(), &inp)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, porePressure)
}
