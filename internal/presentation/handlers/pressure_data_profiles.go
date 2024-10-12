package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/munaiplan/munaiplan-backend/internal/application/types/requests"
	"github.com/munaiplan/munaiplan-backend/internal/helpers"
	"github.com/munaiplan/munaiplan-backend/pkg/values"
)

func (h *Handler) initPressureDataProfilesRoutes(api *gin.RouterGroup) {
	pressureDataProfiles := api.Group("/pressure-data-profiles", h.authMiddleware.UserIdentity)
	{
		pressureDataProfiles.GET("/", h.getPressureDataProfiles)
		pressureDataProfiles.POST("/", h.createPressureDataProfile)
		pressureDataProfiles.GET("/:id", h.getPressureDataProfileByID)
		pressureDataProfiles.PUT("/:id", h.updatePressureDataProfile)
		pressureDataProfiles.DELETE("/:id", h.deletePressureDataProfile)
	}
}

func (h *Handler) getPressureDataProfiles(c *gin.Context) {
	var inp requests.GetPressureDataProfilesRequest
	inp.CaseID = c.Query(values.CaseIdQueryParam)
	profiles, err := h.services.PressureDataProfiles.GetPressureDataProfiles(c.Request.Context(), &inp)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, profiles)
}

func (h *Handler) createPressureDataProfile(c *gin.Context) {
	var inp requests.CreatePressureDataProfileRequest
	var err error

	if err = c.BindJSON(&inp.Body); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, "Invalid request body")
		return
	}
	if inp.CaseID, err = h.validateQueryIDParam(c, values.CaseIdQueryParam); err != nil {
		return
	}
	if err = h.services.PressureDataProfiles.CreatePressureDataProfile(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, helpers.NewResponse("Pressure data profile created successfully"))
}

func (h *Handler) getPressureDataProfileByID(c *gin.Context) {
	var inp requests.GetPressureDataProfileByIDRequest
	inp.ID = c.Param(values.IdQueryParam)
	profile, err := h.services.PressureDataProfiles.GetPressureDataProfileByID(c.Request.Context(), &inp)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, profile)
}

func (h *Handler) updatePressureDataProfile(c *gin.Context) {
	var inp requests.UpdatePressureDataProfileRequest
	inp.ID = c.Param(values.IdQueryParam)
	if err := c.BindJSON(&inp.Body); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, "Invalid request body")
		return
	}
	if err := h.services.PressureDataProfiles.UpdatePressureDataProfile(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, helpers.NewResponse("Pressure data profile updated successfully"))
}

func (h *Handler) deletePressureDataProfile(c *gin.Context) {
	var inp requests.DeletePressureDataProfileRequest
	inp.ID = c.Param(values.IdQueryParam)
	if err := h.services.PressureDataProfiles.DeletePressureDataProfile(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, helpers.NewResponse("Pressure data profile deleted successfully"))
}
