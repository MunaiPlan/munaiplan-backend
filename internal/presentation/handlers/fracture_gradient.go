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

// initFractureGradientRoutes initializes the routes for the fracture gradients API.
func (h *Handler) initFractureGradientRoutes(api *gin.RouterGroup) {
	fractureGradients := api.Group("/fracture_gradients", h.authMiddleware.UserIdentity)
	{
		fractureGradients.GET("/", h.getFractureGradients)
		fractureGradients.POST("/", h.createFractureGradient)
		fractureGradients.GET("/:id", h.getFractureGradientByID)
		fractureGradients.PUT("/:id", h.updateFractureGradient)
		fractureGradients.DELETE("/:id", h.deleteFractureGradient)
	}
}

// getFractureGradients retrieves all fracture gradients.
// @Summary Get Fracture Gradients
// @Tags fracture_gradients
// @Description Retrieves all fracture gradients
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param caseId query string true "Case ID"
// @Success 200 {array} entities.FractureGradient
// @Failure 500 {object} helpers.Response
// @Router /api/v1/fracture_gradients [get]
func (h *Handler) getFractureGradients(c *gin.Context) {
	var inp requests.GetFractureGradientsRequest
	var err error
	var fractureGradients []*entities.FractureGradient

	if inp.CaseID, err = h.validateQueryIDParam(c, values.CaseIdQueryParam); err != nil {
		return
	}

	if fractureGradients, err = h.services.FractureGradients.GetFractureGradients(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, fractureGradients)
}

// createFractureGradient creates a new fracture gradient.
// @Summary Create Fracture Gradient
// @Tags fracture_gradients
// @Description Creates a new fracture gradient
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param caseId query string true "Case ID"
// @Param input body requests.CreateFractureGradientRequest true "Fracture Gradient input"
// @Success 201 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/fracture_gradients [post]
func (h *Handler) createFractureGradient(c *gin.Context) {
	var inp requests.CreateFractureGradientRequest
	var err error

	if err = c.BindJSON(&inp.Body); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, types.ErrInvalidInputBody.Error())
		return
	}
	if inp.CaseID, err = h.validateQueryIDParam(c, values.CaseIdQueryParam); err != nil {
		return
	}
	if err = h.services.FractureGradients.CreateFractureGradient(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, helpers.NewResponse("fracture gradient created"))
}

// updateFractureGradient updates an existing fracture gradient.
// @Summary Update Fracture Gradient
// @Tags fracture_gradients
// @Description Updates an existing fracture gradient
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Fracture Gradient ID"
// @Param input body requests.UpdateFractureGradientRequest true "Fracture Gradient input"
// @Success 200 {object} entities.FractureGradient
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/fracture_gradients/{id} [put]
func (h *Handler) updateFractureGradient(c *gin.Context) {
	var inp requests.UpdateFractureGradientRequest
	var err error

	if err := c.BindJSON(&inp.Body); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, types.ErrInvalidInputBody.Error())
		return
	}
	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}

	fractureGradient, err := h.services.FractureGradients.UpdateFractureGradient(c.Request.Context(), &inp)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, fractureGradient)
}

// deleteFractureGradient deletes an existing fracture gradient.
// @Summary Delete Fracture Gradient
// @Tags fracture_gradients
// @Description Deletes an existing fracture gradient
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Fracture Gradient ID"
// @Success 200 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/fracture_gradients/{id} [delete]
func (h *Handler) deleteFractureGradient(c *gin.Context) {
	var inp requests.DeleteFractureGradientRequest
	var err error

	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}
	if err = h.services.FractureGradients.DeleteFractureGradient(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, helpers.NewResponse("fracture gradient deleted"))
}

// getFractureGradientByID retrieves a fracture gradient by its ID.
// @Summary Get Fracture Gradient by ID
// @Tags fracture_gradients
// @Description Retrieves a fracture gradient by its ID
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Fracture Gradient ID"
// @Success 200 {object} entities.FractureGradient
// @Failure 500 {object} helpers.Response
// @Router /api/v1/fracture_gradients/{id} [get]
func (h *Handler) getFractureGradientByID(c *gin.Context) {
	var inp requests.GetFractureGradientByIDRequest
	var err error
	var fractureGradient *entities.FractureGradient

	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}
	if fractureGradient, err = h.services.FractureGradients.GetFractureGradientByID(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, fractureGradient)
}
