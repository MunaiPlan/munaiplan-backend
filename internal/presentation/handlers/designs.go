package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/munaiplan/munaiplan-backend/internal/application/dto/requests"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/helpers"
	"github.com/munaiplan/munaiplan-backend/internal/presentation/types"
	"github.com/munaiplan/munaiplan-backend/pkg/values"
)

// initDesignsRoutes initializes the routes for the designs API.
func (h *Handler) initDesignsRoutes(api *gin.RouterGroup) {
	designs := api.Group("/designs", h.authMiddleware.UserIdentity)
	{
		designs.GET("/", h.getDesigns)
		designs.POST("/", h.createDesign)
		designs.GET("/:id", h.getDesignByID)
		designs.PUT("/:id", h.updateDesign)
		designs.DELETE("/:id", h.deleteDesign)
	}
}

// getDesigns retrieves all designs.
// @Summary Get Designs
// @Tags designs
// @Description Retrieves all designs
// @Accept json
// @Produce json
// @Param wellboreId query string true "Wellbore ID"
// @Success 200 {array} entities.Design
// @Failure 500 {object} helpers.Response
// @Router /api/v1/designs [get]
func (h *Handler) getDesigns(c *gin.Context) {
	var inp requests.GetDesignsRequest
	var err error
	if inp.WellboreID, err = h.validateQueryIDParam(c, values.WellboreIdQueryParam); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, types.ErrInvalidWellboreIDQueryParameter.Error())
		return
	}
	if err := uuid.Validate(inp.WellboreID); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, types.ErrInvalidUUID.Error())
		return
	}

	designs, err := h.services.Designs.GetDesigns(c.Request.Context(), &inp)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, designs)
}

// createDesign creates a new design.
// @Summary Create Design
// @Tags designs
// @Description Creates a new design
// @Accept json
// @Produce json
// @Param wellboreId query string true "Wellbore ID"
// @Param input body requests.CreateDesignRequest true "Design input"
// @Success 201 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/designs [post]
func (h *Handler) createDesign(c *gin.Context) {
	var inp requests.CreateDesignRequest
	var err error

	if err = c.BindJSON(&inp.Body); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, types.ErrInvalidInputBody.Error())
		return
	}
	if inp.WellboreID, err = h.validateQueryIDParam(c, values.WellboreIdQueryParam); err != nil {
		return
	}
	if err = h.services.Designs.CreateDesign(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, helpers.NewResponse("design created"))
}

// updateDesign updates an existing design.
// @Summary Update Design
// @Tags designs
// @Description Updates an existing design
// @Accept json
// @Produce json
// @Param id path string true "Design ID"
// @Param input body requests.UpdateDesignRequest true "Design input"
// @Success 200 {object} entities.Design
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/designs/{id} [put]
func (h *Handler) updateDesign(c *gin.Context) {
	var inp requests.UpdateDesignRequest
	var err error
	if err := c.BindJSON(&inp.Body); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, types.ErrInvalidInputBody.Error())
		return
	}
	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}

	design, err := h.services.Designs.UpdateDesign(c.Request.Context(), &inp)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, design)
}

// deleteDesign deletes an existing design.
// @Summary Delete Design
// @Tags designs
// @Description Deletes an existing design
// @Accept json
// @Produce json
// @Param id path string true "Design ID"
// @Success 200 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/designs/{id} [delete]
func (h *Handler) deleteDesign(c *gin.Context) {
	var inp requests.DeleteDesignRequest
	var err error

	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}
	if err = h.services.Designs.DeleteDesign(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, helpers.NewResponse("design deleted"))
}

// getDesignByID retrieves a design by its ID.
// @Summary Get Design by ID
// @Tags designs
// @Description Retrieves a design by its ID
// @Accept json
// @Produce json
// @Param id path string true "Design ID"
// @Success 200 {object} entities.Design
// @Failure 500 {object} helpers.Response
// @Router /api/v1/designs/{id} [get]
func (h *Handler) getDesignByID(c *gin.Context) {
	var inp requests.GetDesignByIDRequest
	var err error
	var design *entities.Design
	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}
	if design, err = h.services.Designs.GetDesignByID(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, design)
}