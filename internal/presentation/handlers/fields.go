package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/munaiplan/munaiplan-backend/internal/application/dto/requests"
	"github.com/munaiplan/munaiplan-backend/internal/helpers"
	"github.com/munaiplan/munaiplan-backend/internal/presentation/types"
	"github.com/munaiplan/munaiplan-backend/pkg/values"
)

// initFieldsRoutes initializes the routes for the fields API.
func (h *Handler) initFieldsRoutes(api *gin.RouterGroup) {
	fields := api.Group("/fields")
	{
		fields.GET("/", h.getFields)
		fields.POST("/", h.createField)
		fields.GET("/:id", h.getFieldByID)
		fields.PUT("/:id", h.updateField)
		fields.DELETE("/:id", h.deleteField)
	}
}

// getFields retrieves all fields.
// @Summary Get Fields
// @Tags fields
// @Description Retrieves all fields
// @Accept json
// @Produce json
// @Param companyId query string true "Company ID"
// @Success 200 {array} entities.Field
// @Failure 500 {object} helpers.Response
// @Router /api/v1/fields [get]
func (h *Handler) getFields(c *gin.Context) {
	var inp requests.GetFieldsRequest
	var err error
	if inp.CompanyID, err = h.validateQueryParam(c, values.CompanyIdQueryParam); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, types.ErrInvalidCompanyIDQueryParameter.Error())
		return
	}
	if err := uuid.Validate(inp.CompanyID); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, types.ErrInvalidUUID.Error())
		return
	}

	fields, err := h.services.Fields.GetFields(c.Request.Context(), &inp)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, fields)
}

// createField creates a new field.
// @Summary Create Field
// @Tags fields
// @Description Creates a new field
// @Accept json
// @Produce json
// @Param companyId query string true "Company ID"
// @Param input body requests.CreateFieldRequest true "Field input"
// @Success 201 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/fields [post]
func (h *Handler) createField(c *gin.Context) {
	var inp requests.CreateFieldRequest
	var err error

	if err = c.BindJSON(&inp.Body); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, types.ErrInvalidInputBody.Error())
		return
	}

	if inp.CompanyID, err = h.validateQueryParam(c, values.CompanyIdQueryParam); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, types.ErrInvalidCompanyIDQueryParameter.Error())
		return
	}
	if err := uuid.Validate(inp.CompanyID); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, types.ErrInvalidUUID.Error())
		return
	}

	if err = h.services.Fields.CreateField(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, helpers.NewResponse("field created"))
}

// updateField updates an existing field.
// @Summary Update Field
// @Tags fields
// @Description Updates an existing field
// @Accept json
// @Produce json
// @Param id path string true "Field ID"
// @Param companyId query string true "Company ID"
// @Param input body requests.UpdateFieldRequest true "Field input"
// @Success 200 {object} entities.Field
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/fields/{id} [put]
func (h *Handler) updateField(c *gin.Context) {
	var inp requests.UpdateFieldRequest
	var err error
	if err := c.BindJSON(&inp.Body); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, types.ErrInvalidInputBody.Error())
		return
	}

	if inp.CompanyID, err = h.validateQueryParam(c, values.CompanyIdQueryParam); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, types.ErrInvalidCompanyIDQueryParameter.Error())
		return
	}
	if inp.Body.ID, err = h.validateRequestParam(c, values.IdQueryParam); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, types.ErrInvalidIDQueryParameter.Error())
		return
	}
	if uuid.Validate(inp.CompanyID) != nil || uuid.Validate(inp.Body.ID) != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, types.ErrInvalidUUID.Error())
		return
	}

	field, err := h.services.Fields.UpdateField(c.Request.Context(), &inp)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, field)
}

// deleteField deletes an existing field.
// @Summary Delete Field
// @Tags fields
// @Description Deletes an existing field
// @Accept json
// @Produce json
// @Param companyId query string true "Company ID"
// @Param id path string true "Field ID"
// @Success 200 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/fields/{id} [delete]
func (h *Handler) deleteField(c *gin.Context) {
	var inp requests.DeleteFieldRequest
	var err error

	if inp.ID, err = h.validateRequestParam(c, values.IdQueryParam); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, types.ErrInvalidIDQueryParameter.Error())
		return
	}

	if inp.CompanyID, err = h.validateQueryParam(c, values.CompanyIdQueryParam); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, types.ErrInvalidCompanyIDQueryParameter.Error())
		return
	}
	if uuid.Validate(inp.CompanyID) != nil || uuid.Validate(inp.ID) != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, types.ErrInvalidUUID.Error())
		return
	}

	err = h.services.Fields.DeleteField(c.Request.Context(), &inp)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, helpers.NewResponse("field deleted"))
}

// getFieldByID retrieves a field by its ID.
// @Summary Get Field by ID
// @Tags fields
// @Description Retrieves a field by its ID
// @Accept json
// @Produce json
// @Param companyId query string true "Company ID"
// @Param id path string true "Field ID"
// @Success 200 {object} entities.Field
// @Failure 500 {object} helpers.Response
// @Router /api/v1/fields/{id} [get]
func (h *Handler) getFieldByID(c *gin.Context) {
	var inp requests.GetFieldByIDRequest
	var err error
	if inp.ID, err = h.validateRequestParam(c, values.IdQueryParam); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, types.ErrInvalidIDQueryParameter.Error())
		return
	}

	if inp.CompanyID, err = h.validateQueryParam(c, values.CompanyIdQueryParam); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, types.ErrInvalidCompanyIDQueryParameter.Error())
		return
	}
	if uuid.Validate(inp.CompanyID) != nil || uuid.Validate(inp.ID) != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, types.ErrInvalidUUID.Error())
		return
	}

	field, err := h.services.Fields.GetFieldByID(c.Request.Context(), &inp)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, field)
}
