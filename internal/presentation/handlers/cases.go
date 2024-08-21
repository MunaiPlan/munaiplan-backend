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

// initCasesRoutes initializes the routes for the cases API.
func (h *Handler) initCasesRoutes(api *gin.RouterGroup) {
	cases := api.Group("/cases", h.authMiddleware.UserIdentity)
	{
		cases.GET("/", h.getCases)
		cases.POST("/", h.createCase)
		cases.GET("/:id", h.getCaseByID)
		cases.PUT("/:id", h.updateCase)
		cases.DELETE("/:id", h.deleteCase)
	}
}

// getCases retrieves all cases.
// @Summary Get Cases
// @Tags cases
// @Description Retrieves all cases
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param trajectoryId query string true "Trajectory ID"
// @Success 200 {array} entities.Case
// @Failure 500 {object} helpers.Response
// @Router /api/v1/cases [get]
func (h *Handler) getCases(c *gin.Context) {
	var inp requests.GetCasesRequest
	var err error
	var cases []*entities.Case
	if inp.TrajectoryID, err = h.validateQueryIDParam(c, values.TrajectoryIdQueryParam); err != nil {
		return
	}

	if cases, err = h.services.Cases.GetCases(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, cases)
}

// createCase creates a new case.
// @Summary Create Case
// @Tags cases
// @Description Creates a new case
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param trajectoryId query string true "Trajectory ID"
// @Param input body requests.CreateCaseRequest true "Case input"
// @Success 201 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/cases [post]
func (h *Handler) createCase(c *gin.Context) {
	var inp requests.CreateCaseRequest
	var err error

	if err = c.BindJSON(&inp.Body); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, types.ErrInvalidInputBody.Error())
		return
	}
	if inp.TrajectoryID, err = h.validateQueryIDParam(c, values.TrajectoryIdQueryParam); err != nil {
		return
	}
	if err = h.services.Cases.CreateCase(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, helpers.NewResponse("case created"))
}

// updateCase updates an existing case.
// @Summary Update Case
// @Tags cases
// @Description Updates an existing case
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Case ID"
// @Param input body requests.UpdateCaseRequest true "Case input"
// @Success 200 {object} entities.Case
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/cases/{id} [put]
func (h *Handler) updateCase(c *gin.Context) {
	var inp requests.UpdateCaseRequest
	var err error
	if err := c.BindJSON(&inp.Body); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, types.ErrInvalidInputBody.Error())
		return
	}
	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}

	caseEntity, err := h.services.Cases.UpdateCase(c.Request.Context(), &inp)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, caseEntity)
}

// deleteCase deletes an existing case.
// @Summary Delete Case
// @Tags cases
// @Description Deletes an existing case
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Case ID"
// @Success 200 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/cases/{id} [delete]
func (h *Handler) deleteCase(c *gin.Context) {
	var inp requests.DeleteCaseRequest
	var err error

	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}
	if err = h.services.Cases.DeleteCase(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, helpers.NewResponse("case deleted"))
}

// getCaseByID retrieves a case by its ID.
// @Summary Get Case by ID
// @Tags cases
// @Description Retrieves a case by its ID
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Case ID"
// @Success 200 {object} entities.Case
// @Failure 500 {object} helpers.Response
// @Router /api/v1/cases/{id} [get]
func (h *Handler) getCaseByID(c *gin.Context) {
	var inp requests.GetCaseByIDRequest
	var err error
	var caseEntity *entities.Case
	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}
	if caseEntity, err = h.services.Cases.GetCaseByID(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, caseEntity)
}
