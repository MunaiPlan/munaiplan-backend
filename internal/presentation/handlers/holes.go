package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/munaiplan/munaiplan-backend/internal/application/types/requests"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/helpers"
	"github.com/munaiplan/munaiplan-backend/internal/presentation/types"
	"github.com/munaiplan/munaiplan-backend/pkg/values"
)

// initHolesRoutes initializes the routes for the holes API.
func (h *Handler) initHolesRoutes(api *gin.RouterGroup) {
	holes := api.Group("/holes", h.authMiddleware.UserIdentity)
	{
		holes.GET("/", h.getHoles)
		holes.POST("/", h.createHole)
		holes.GET("/:id", h.getHoleByID)
		holes.PUT("/:id", h.updateHole)
		holes.DELETE("/:id", h.deleteHole)
	}
}

// getHoles retrieves all holes.
// @Summary Get Holes
// @Tags holes
// @Description Retrieves all holes
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param caseId query string true "Case ID"
// @Success 200 {array} entities.Hole
// @Failure 500 {object} helpers.Response
// @Router /api/v1/holes [get]
func (h *Handler) getHoles(c *gin.Context) {
	var inp requests.GetHolesRequest
	var err error
	var holes []*entities.Hole
	if inp.CaseID, err = h.validateQueryIDParam(c, values.CaseIdQueryParam); err != nil {
		return
	}
	
	if holes, err = h.services.Holes.GetHoles(c.Request.Context(), &inp); err != nil {
		fmt.Println(holes)
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, holes)
}

// createHole creates a new hole.
// @Summary Create Hole
// @Tags holes
// @Description Creates a new hole
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param caseId query string true "Case ID"
// @Param input body requests.CreateHoleRequest true "Hole input"
// @Success 201 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/holes [post]
func (h *Handler) createHole(c *gin.Context) {
	var inp requests.CreateHoleRequest
	var err error

	if err = c.BindJSON(&inp.Body); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, types.ErrInvalidInputBody.Error())
		return
	}
	if inp.CaseID, err = h.validateQueryIDParam(c, values.CaseIdQueryParam); err != nil {
		return
	}
	if err = h.services.Holes.CreateHole(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, helpers.NewResponse("hole created"))
}

// updateHole updates an existing hole.
// @Summary Update Hole
// @Tags holes
// @Description Updates an existing hole
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Hole ID"
// @Param input body requests.UpdateHoleRequest true "Hole input"
// @Success 200 {object} entities.Hole
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/holes/{id} [put]
func (h *Handler) updateHole(c *gin.Context) {
	var inp requests.UpdateHoleRequest
	var err error
	if err := c.BindJSON(&inp.Body); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, types.ErrInvalidInputBody.Error())
		return
	}
	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}
	fmt.Println("id is " + inp.ID)

	hole, err := h.services.Holes.UpdateHole(c.Request.Context(), &inp)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, hole)
}

// deleteHole deletes an existing hole.
// @Summary Delete Hole
// @Tags holes
// @Description Deletes an existing hole
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Hole ID"
// @Success 200 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/holes/{id} [delete]
func (h *Handler) deleteHole(c *gin.Context) {
	var inp requests.DeleteHoleRequest
	var err error

	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}
	if err = h.services.Holes.DeleteHole(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, helpers.NewResponse("hole deleted"))
}

// getHoleByID retrieves a hole by its ID.
// @Summary Get Hole by ID
// @Tags holes
// @Description Retrieves a hole by its ID
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Hole ID"
// @Success 200 {object} entities.Hole
// @Failure 500 {object} helpers.Response
// @Router /api/v1/holes/{id} [get]
func (h *Handler) getHoleByID(c *gin.Context) {
	var inp requests.GetHoleByIDRequest
	var err error
	var hole *entities.Hole
	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}
	if hole, err = h.services.Holes.GetHoleByID(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, hole)
}
