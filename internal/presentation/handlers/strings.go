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

// initStringsRoutes initializes the routes for the strings API.
func (h *Handler) initStringsRoutes(api *gin.RouterGroup) {
	strings := api.Group("/strings", h.authMiddleware.UserIdentity)
	{
		strings.GET("/", h.getStrings)
		strings.POST("/", h.createString)
		strings.GET("/:id", h.getStringByID)
		strings.PUT("/:id", h.updateString)
		strings.DELETE("/:id", h.deleteString)
	}
}

// getStrings retrieves all strings associated with a case ID.
// @Summary Get Strings
// @Tags strings
// @Description Retrieves all strings
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param caseId query string true "Case ID"
// @Success 200 {array} entities.String
// @Failure 500 {object} helpers.Response
// @Router /api/v1/strings [get]
func (h *Handler) getStrings(c *gin.Context) {
	var inp requests.GetStringsRequest
	var err error
	var strings []*entities.String
	if inp.CaseID, err = h.validateQueryIDParam(c, values.CaseIdQueryParam); err != nil {
		return
	}

	if strings, err = h.services.Strings.GetStrings(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, strings)
}

// createString creates a new string along with its sections.
// @Summary Create String
// @Tags strings
// @Description Creates a new string
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param caseId query string true "Case ID"
// @Param input body requests.CreateStringRequestBody true "String input"
// @Success 201 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/strings [post]
func (h *Handler) createString(c *gin.Context) {
	var inp requests.CreateStringRequest
	var err error

	if err = c.BindJSON(&inp.Body); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, types.ErrInvalidInputBody.Error())
		return
	}
	if inp.CaseID, err = h.validateQueryIDParam(c, values.CaseIdQueryParam); err != nil {
		return
	}
	if err = h.services.Strings.CreateString(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, helpers.NewResponse("string created"))
}

// updateString updates an existing string and its sections.
// @Summary Update String
// @Tags strings
// @Description Updates an existing string
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "String ID"
// @Param input body requests.UpdateStringRequestBody true "String input"
// @Success 200 {object} entities.String
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/strings/{id} [put]
func (h *Handler) updateString(c *gin.Context) {
	var inp requests.UpdateStringRequest
	var err error
	if err := c.BindJSON(&inp.Body); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, types.ErrInvalidInputBody.Error())
		return
	}
	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}

	stringEntity, err := h.services.Strings.UpdateString(c.Request.Context(), &inp)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, stringEntity)
}

// deleteString deletes an existing string.
// @Summary Delete String
// @Tags strings
// @Description Deletes an existing string
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "String ID"
// @Success 200 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/strings/{id} [delete]
func (h *Handler) deleteString(c *gin.Context) {
	var inp requests.DeleteStringRequest
	var err error

	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}
	if err = h.services.Strings.DeleteString(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, helpers.NewResponse("string deleted"))
}

// getStringByID retrieves a string by its ID.
// @Summary Get String by ID
// @Tags strings
// @Description Retrieves a string by its ID
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "String ID"
// @Success 200 {object} entities.String
// @Failure 500 {object} helpers.Response
// @Router /api/v1/strings/{id} [get]
func (h *Handler) getStringByID(c *gin.Context) {
	var inp requests.GetStringByIDRequest
	var err error
	var stringEntity *entities.String
	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}
	if stringEntity, err = h.services.Strings.GetStringByID(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, stringEntity)
}
