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

// initWellboresRoutes initializes the routes for the wellbores API.
func (h *Handler) initWellboresRoutes(api *gin.RouterGroup) {
	wellbores := api.Group("/wellbores", h.authMiddleware.UserIdentity)
	{
		wellbores.GET("/", h.getWellbores)
		wellbores.POST("/", h.createWellbore)
		wellbores.GET("/:id", h.getWellboreByID)
		wellbores.PUT("/:id", h.updateWellbore)
		wellbores.DELETE("/:id", h.deleteWellbore)
	}
}

// getWellbores retrieves all wellbores.
// @Summary Get Wellbores
// @Tags wellbores
// @Description Retrieves all wellbores
// @Accept json
// @Produce json
// @Param wellId query string true "Well ID"
// @Success 200 {array} entities.Wellbore
// @Failure 500 {object} helpers.Response
// @Router /api/v1/wellbores [get]
func (h *Handler) getWellbores(c *gin.Context) {
	var inp requests.GetWellboresRequest
	var err error
	if inp.WellID, err = h.validateQueryIDParam(c, values.WellIdQueryParam); err != nil {
		return
	}

	wellbores, err := h.services.Wellbores.GetWellbores(c.Request.Context(), &inp)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, wellbores)
}

// createWellbore creates a new wellbore.
// @Summary Create Wellbore
// @Tags wellbores
// @Description Creates a new wellbore
// @Accept json
// @Produce json
// @Param wellId query string true "Well ID"
// @Param input body requests.CreateWellboreRequest true "Wellbore input"
// @Success 201 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/wellbores [post]
func (h *Handler) createWellbore(c *gin.Context) {
	var inp requests.CreateWellboreRequest
	var err error

	if err = c.BindJSON(&inp.Body); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, types.ErrInvalidInputBody.Error())
		return
	}
	if inp.WellID, err = h.validateQueryIDParam(c, values.WellIdQueryParam); err != nil {
		return
	}
	if err = h.services.Wellbores.CreateWellbore(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, helpers.NewResponse("wellbore created"))
}

// updateWellbore updates an existing wellbore.
// @Summary Update Wellbore
// @Tags wellbores
// @Description Updates an existing wellbore
// @Accept json
// @Produce json
// @Param id path string true "Wellbore ID"
// @Param input body requests.UpdateWellboreRequest true "Wellbore input"
// @Success 200 {object} entities.Wellbore
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/wellbores/{id} [put]
func (h *Handler) updateWellbore(c *gin.Context) {
	var inp requests.UpdateWellboreRequest
	var err error
	if err := c.BindJSON(&inp.Body); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, types.ErrInvalidInputBody.Error())
		return
	}
	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}

	wellbore, err := h.services.Wellbores.UpdateWellbore(c.Request.Context(), &inp)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, wellbore)
}

// deleteWellbore deletes an existing wellbore.
// @Summary Delete Wellbore
// @Tags wellbores
// @Description Deletes an existing wellbore
// @Accept json
// @Produce json
// @Param id path string true "Wellbore ID"
// @Success 200 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/wellbores/{id} [delete]
func (h *Handler) deleteWellbore(c *gin.Context) {
	var inp requests.DeleteWellboreRequest
	var err error

	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}
	if err = h.services.Wellbores.DeleteWellbore(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, helpers.NewResponse("wellbore deleted"))
}

// getWellboreByID retrieves a wellbore by its ID.
// @Summary Get Wellbore by ID
// @Tags wellbores
// @Description Retrieves a wellbore by its ID
// @Accept json
// @Produce json
// @Param id path string true "Wellbore ID"
// @Success 200 {object} entities.Wellbore
// @Failure 500 {object} helpers.Response
// @Router /api/v1/wellbores/{id} [get]
func (h *Handler) getWellboreByID(c *gin.Context) {
	var inp requests.GetWellboreByIDRequest
	var err error
	var wellbore *entities.Wellbore
	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}
	if wellbore, err = h.services.Wellbores.GetWellboreByID(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, wellbore)
}
