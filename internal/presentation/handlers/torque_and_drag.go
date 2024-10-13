package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/munaiplan/munaiplan-backend/internal/helpers"
	"github.com/munaiplan/munaiplan-backend/pkg/values"
)

// initTorqueAndDragRoutes initializes routes for the Torque and Drag module.
func (h *Handler) initTorqueAndDragRoutes(api *gin.RouterGroup) {
	torqueAndDrag := api.Group("/torque-and-drag", h.authMiddleware.UserIdentity)
	{
		torqueAndDrag.POST("/effective-tension", h.calculateEffectiveTensionFromMLModel)
	}
}

// calculateEffectiveTension handles the calculation of effective tension.
// @Summary Calculate Effective Tension
// @Tags torque-and-drag
// @Description Calculates effective tension based on input data.
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param input body requests.EffectiveTensionRequest true "Effective Tension Input Data"
// @Success 200 {object} entities.EffectiveTensionResponse
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/torque-and-drag/effective-tension [post]
func (h *Handler) calculateEffectiveTensionFromMLModel(c *gin.Context) {
	caseID, err := h.validateQueryIDParam(c, values.CaseIdQueryParam)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Call the service to calculate effective tension
	result, err := h.services.TorqueAndDrag.CalculateEffectiveTensionFromMLModel(c.Request.Context(), caseID)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Respond with the result
	c.JSON(http.StatusOK, result)
}
