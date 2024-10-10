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

// initRigsRoutes initializes the routes for the rigs API.
func (h *Handler) initRigsRoutes(api *gin.RouterGroup) {
	rigs := api.Group("/rigs", h.authMiddleware.UserIdentity)
	{
		rigs.GET("/", h.getRigs)
		rigs.POST("/", h.createRig)
		rigs.GET("/:id", h.getRigByID)
		rigs.PUT("/:id", h.updateRig)
		rigs.DELETE("/:id", h.deleteRig)
	}
}

// getRigs retrieves all rigs associated with a case ID.
// @Summary Get Rigs
// @Tags rigs
// @Description Retrieves all rigs for a case
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param caseId query string true "Case ID"
// @Success 200 {array} entities.Rig
// @Failure 500 {object} helpers.Response
// @Router /api/v1/rigs [get]
func (h *Handler) getRigs(c *gin.Context) {
	var inp requests.GetRigsRequest
	var err error
	var rigs []*entities.Rig

	if inp.CaseID, err = h.validateQueryIDParam(c, values.CaseIdQueryParam); err != nil {
		return
	}

	if rigs, err = h.services.Rigs.GetRigs(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, rigs)
}

// createRig creates a new rig.
// @Summary Create Rig
// @Tags rigs
// @Description Creates a new rig for a specific case
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param caseId query string true "Case ID"
// @Param input body requests.CreateRigRequest true "Rig input"
// @Success 201 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/rigs [post]
func (h *Handler) createRig(c *gin.Context) {
	var inp requests.CreateRigRequest
	var err error

	if err = c.BindJSON(&inp.Body); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, types.ErrInvalidInputBody.Error())
		return
	}
	if inp.CaseID, err = h.validateQueryIDParam(c, values.CaseIdQueryParam); err != nil {
		return
	}
	if err = h.services.Rigs.CreateRig(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, helpers.NewResponse("rig created"))
}

// getRigByID retrieves a rig by its ID.
// @Summary Get Rig by ID
// @Tags rigs
// @Description Retrieves a rig by its ID
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Rig ID"
// @Success 200 {object} entities.Rig
// @Failure 500 {object} helpers.Response
// @Router /api/v1/rigs/{id} [get]
func (h *Handler) getRigByID(c *gin.Context) {
	var inp requests.GetRigByIDRequest
	var err error
	var rig *entities.Rig

	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}
	if rig, err = h.services.Rigs.GetRigByID(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, rig)
}

// updateRig updates an existing rig.
// @Summary Update Rig
// @Tags rigs
// @Description Updates an existing rig
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Rig ID"
// @Param input body requests.UpdateRigRequest true "Rig input"
// @Success 200 {object} entities.Rig
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/rigs/{id} [put]
func (h *Handler) updateRig(c *gin.Context) {
	var inp requests.UpdateRigRequest
	var err error

	if err := c.BindJSON(&inp.Body); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, types.ErrInvalidInputBody.Error())
		return
	}
	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}

	rig, err := h.services.Rigs.UpdateRig(c.Request.Context(), &inp)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, rig)
}

// deleteRig deletes a rig by its ID.
// @Summary Delete Rig
// @Tags rigs
// @Description Deletes a rig by its ID
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Rig ID"
// @Success 200 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/rigs/{id} [delete]
func (h *Handler) deleteRig(c *gin.Context) {
	var inp requests.DeleteRigRequest
	var err error

	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}
	if err = h.services.Rigs.DeleteRig(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, helpers.NewResponse("rig deleted"))
}
