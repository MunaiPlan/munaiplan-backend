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
		torqueAndDrag.POST("/weight-on-bit", h.calculateWeightOnBitFromMLModel)
		torqueAndDrag.POST("/surface-torque", h.calculateMomentFromMLModel)
		torqueAndDrag.POST("/min-weight", h.calculateMinWeightFromMLModel)
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

// calculateWeightOnBitFromMLModel handles the calculation of Weight on Bit (WOB) based on input data.
// @Summary Calculate Weight on Bit
// @Tags torque-and-drag
// @Description Calculates Weight on Bit based on input data.
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param input body requests.WeightOnBitRequest true "Weight on Bit Input Data"
// @Success 200 {object} entities.WeightOnBitResponse
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/torque-and-drag/weight-on-bit [post]
func (h *Handler) calculateWeightOnBitFromMLModel(c *gin.Context) {
	caseID, err := h.validateQueryIDParam(c, values.CaseIdQueryParam)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Call the service to calculate weight on bit
	result, err := h.services.TorqueAndDrag.CalculateWeightOnBitFromMlModel(c.Request.Context(), caseID)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Respond with the result
	c.JSON(http.StatusOK, result)
}

// calculateMomentFromMLModel handles the calculation of moment from the ML model.
// @Summary Calculate Moment
// @Tags torque-and-drag
// @Description Calculates moment based on input data and returns the results from the ML model.
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param input body requests.MomentRequest true "Moment Input Data"
// @Success 200 {object} responses.MomentFromMLModelResponse
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/torque-and-drag/moment [post]
func (h *Handler) calculateMomentFromMLModel(c *gin.Context) {
	caseID, err := h.validateQueryIDParam(c, values.CaseIdQueryParam)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Call the service to calculate the moment from the ML model
	result, err := h.services.TorqueAndDrag.CalculateSurfaceTorqueFromMlModel(c.Request.Context(), caseID)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Respond with the result
	c.JSON(http.StatusOK, result)
}

// calculateMinWeightFromMlModel handles the calculation of minimum weight using the ML model.
// @Summary Calculate Minimum Weight
// @Tags torque-and-drag
// @Description Calculates minimum weight based on input data using the ML model.
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param caseId query string true "Case ID"
// @Success 200 {object} entities.MinWeightFromMLModelResponse
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/torque-and-drag/min-weight [post]
func (h *Handler) calculateMinWeightFromMLModel(c *gin.Context) {
	// Validate the Case ID query parameter
	caseID, err := h.validateQueryIDParam(c, values.CaseIdQueryParam)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Call the service to calculate minimum weight
	result, err := h.services.TorqueAndDrag.CalculateMinWeightFromMLModel(c.Request.Context(), caseID)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Respond with the result
	c.JSON(http.StatusOK, result)
}
