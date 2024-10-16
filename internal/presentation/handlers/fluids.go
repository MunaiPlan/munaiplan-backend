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

// initFluidsRoutes initializes the routes for the fluids API.
func (h *Handler) initFluidsRoutes(api *gin.RouterGroup) {
	fluids := api.Group("/fluids", h.authMiddleware.UserIdentity)
	{
		fluids.GET("/", h.getFluids)
		fluids.GET("/types", h.getFluidTypes)
		fluids.POST("/", h.createFluid)
		fluids.GET("/:id", h.getFluidByID)
		fluids.PUT("/:id", h.updateFluid)
		fluids.DELETE("/:id", h.deleteFluid)
	}
}

// getFluids retrieves all fluids associated with a case.
// @Summary Get Fluids
// @Tags fluids
// @Description Retrieves all fluids associated with a case
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param caseId query string true "Case ID"
// @Success 200 {array} entities.Fluid
// @Failure 500 {object} helpers.Response
// @Router /api/v1/fluids [get]
func (h *Handler) getFluids(c *gin.Context) {
	var inp requests.GetFluidsRequest
	var err error
	var fluids []*entities.Fluid

	if inp.CaseID, err = h.validateQueryIDParam(c, values.CaseIdQueryParam); err != nil {
		return
	}

	if fluids, err = h.services.Fluids.GetFluids(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, fluids)
}

// getFluidTypes retrieves all fluid types from the database.
// @Summary Get Fluid Types
// @Tags fluids
// @Description Retrieves all fluid types from the database
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {array} entities.FluidType
// @Failure 500 {object} helpers.Response
// @Router /api/v1/fluids/types [get]
func (h *Handler) getFluidTypes(c *gin.Context) {
	var err error
	var fluidTypes []*entities.FluidType

	if fluidTypes, err = h.services.Fluids.GetFluidTypes(c.Request.Context()); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, fluidTypes)
}

// createFluid creates a new fluid.
// @Summary Create Fluid
// @Tags fluids
// @Description Creates a new fluid
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param caseId query string true "Case ID"
// @Param input body requests.CreateFluidRequest true "Fluid input"
// @Success 201 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/fluids [post]
func (h *Handler) createFluid(c *gin.Context) {
	var inp requests.CreateFluidRequest
	var err error

	if err = c.BindJSON(&inp.Body); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, types.ErrInvalidInputBody.Error())
		return
	}
	if inp.CaseID, err = h.validateQueryIDParam(c, values.CaseIdQueryParam); err != nil {
		return
	}
	if err = h.services.Fluids.CreateFluid(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, helpers.NewResponse("fluid created"))
}

// updateFluid updates an existing fluid.
// @Summary Update Fluid
// @Tags fluids
// @Description Updates an existing fluid
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Fluid ID"
// @Param input body requests.UpdateFluidRequest true "Fluid input"
// @Success 200 {object} entities.Fluid
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/fluids/{id} [put]
func (h *Handler) updateFluid(c *gin.Context) {
	var inp requests.UpdateFluidRequest
	var err error

	if err := c.BindJSON(&inp.Body); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, types.ErrInvalidInputBody.Error())
		return
	}
	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}

	fluid, err := h.services.Fluids.UpdateFluid(c.Request.Context(), &inp)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, fluid)
}

// deleteFluid deletes an existing fluid.
// @Summary Delete Fluid
// @Tags fluids
// @Description Deletes an existing fluid
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Fluid ID"
// @Success 200 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/fluids/{id} [delete]
func (h *Handler) deleteFluid(c *gin.Context) {
	var inp requests.DeleteFluidRequest
	var err error

	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}
	if err = h.services.Fluids.DeleteFluid(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, helpers.NewResponse("fluid deleted"))
}

// getFluidByID retrieves a fluid by its ID.
// @Summary Get Fluid by ID
// @Tags fluids
// @Description Retrieves a fluid by its ID
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Fluid ID"
// @Success 200 {object} entities.Fluid
// @Failure 500 {object} helpers.Response
// @Router /api/v1/fluids/{id} [get]
func (h *Handler) getFluidByID(c *gin.Context) {
	var inp requests.GetFluidByIDRequest
	var err error
	var fluid *entities.Fluid

	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}
	if fluid, err = h.services.Fluids.GetFluidByID(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, fluid)
}
