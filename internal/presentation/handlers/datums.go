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

// initDatumRoutes initializes the routes for the datums API.
func (h *Handler) initDatumRoutes(api *gin.RouterGroup) {
	datums := api.Group("/datums", h.authMiddleware.UserIdentity)
	{
		datums.GET("/", h.getDatumsByCaseID)
		datums.POST("/", h.createDatum)
		datums.GET("/:id", h.getDatumByID)
		datums.PUT("/:id", h.updateDatum)
		datums.DELETE("/:id", h.deleteDatum)
	}
}

// getDatumsByCaseID retrieves all datums associated with a specific case ID.
// @Summary Get Datums
// @Tags datums
// @Description Retrieves all datums associated with a specific case ID
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param caseId query string true "Case ID"
// @Success 200 {array} entities.Datum
// @Failure 500 {object} helpers.Response
// @Router /api/v1/datums [get]
func (h *Handler) getDatumsByCaseID(c *gin.Context) {
	var inp requests.GetDatumsByCaseIDRequest
	var err error
	var datums []*entities.Datum
	if inp.CaseID, err = h.validateQueryIDParam(c, values.CaseIdQueryParam); err != nil {
		return
	}

	if datums, err = h.services.Datums.GetDatumsByCaseID(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, datums)
}

// createDatum creates a new datum.
// @Summary Create Datum
// @Tags datums
// @Description Creates a new datum
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param caseId query string true "Case ID"
// @Param input body requests.CreateDatumRequest true "Datum input"
// @Success 201 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/datums [post]
func (h *Handler) createDatum(c *gin.Context) {
	var inp requests.CreateDatumRequest
	var err error

	if err = c.BindJSON(&inp.Body); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, types.ErrInvalidInputBody.Error())
		return
	}
	if inp.CaseID, err = h.validateQueryIDParam(c, values.CaseIdQueryParam); err != nil {
		return
	}
	if err = h.services.Datums.CreateDatum(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, helpers.NewResponse("datum created"))
}

// updateDatum updates an existing datum.
// @Summary Update Datum
// @Tags datums
// @Description Updates an existing datum
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Datum ID"
// @Param input body requests.UpdateDatumRequest true "Datum input"
// @Success 200 {object} entities.Datum
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/datums/{id} [put]
func (h *Handler) updateDatum(c *gin.Context) {
	var inp requests.UpdateDatumRequest
	var err error
	if err := c.BindJSON(&inp.Body); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, types.ErrInvalidInputBody.Error())
		return
	}
	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}

	datum, err := h.services.Datums.UpdateDatum(c.Request.Context(), &inp)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, datum)
}

// deleteDatum deletes an existing datum.
// @Summary Delete Datum
// @Tags datums
// @Description Deletes an existing datum
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Datum ID"
// @Success 200 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/datums/{id} [delete]
func (h *Handler) deleteDatum(c *gin.Context) {
	var inp requests.DeleteDatumRequest
	var err error

	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}
	if err = h.services.Datums.DeleteDatum(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, helpers.NewResponse("datum deleted"))
}

// getDatumByID retrieves a datum by its ID.
// @Summary Get Datum by ID
// @Tags datums
// @Description Retrieves a datum by its ID
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Datum ID"
// @Success 200 {object} entities.Datum
// @Failure 500 {object} helpers.Response
// @Router /api/v1/datums/{id} [get]
func (h *Handler) getDatumByID(c *gin.Context) {
	var inp requests.GetDatumByIDRequest
	var err error
	var datum *entities.Datum
	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}
	if datum, err = h.services.Datums.GetDatumByID(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, datum)
}
