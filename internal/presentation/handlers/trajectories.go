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

// initTrajectoriesRoutes initializes the routes for the trajectories API.
func (h *Handler) initTrajectoriesRoutes(api *gin.RouterGroup) {
	trajectories := api.Group("/trajectories", h.authMiddleware.UserIdentity)
	{
		trajectories.GET("/", h.getTrajectories)
		trajectories.POST("/", h.createTrajectory)
		trajectories.GET("/:id", h.getTrajectoryByID)
		trajectories.PUT("/:id", h.updateTrajectory)
		trajectories.DELETE("/:id", h.deleteTrajectory)
	}
}

// getTrajectories retrieves all trajectories.
// @Summary Get Trajectories
// @Tags trajectories
// @Description Retrieves all trajectories
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param designId query string true "Design ID"
// @Success 200 {array} entities.Trajectory
// @Failure 500 {object} helpers.Response
// @Router /api/v1/trajectories [get]
func (h *Handler) getTrajectories(c *gin.Context) {
	var inp requests.GetTrajectoriesRequest
	var err error
	var trajectories []*entities.Trajectory

	if inp.DesignID, err = h.validateQueryIDParam(c, values.DesignIdQueryParam); err != nil {
		return
	}

	if trajectories, err = h.services.Trajectories.GetTrajectories(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, trajectories)
}

// createTrajectory creates a new trajectory.
// @Summary Create Trajectory
// @Tags trajectories
// @Description Creates a new trajectory
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param designId query string true "Design ID"
// @Param input body requests.CreateTrajectoryRequest true "Trajectory input"
// @Success 201 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/trajectories [post]
func (h *Handler) createTrajectory(c *gin.Context) {
	var inp requests.CreateTrajectoryRequest
	var err error

	if err = c.BindJSON(&inp.Body); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, types.ErrInvalidInputBody.Error())
		return
	}
	if inp.DesignID, err = h.validateQueryIDParam(c, values.DesignIdQueryParam); err != nil {
		return
	}
	if err = h.services.Trajectories.CreateTrajectory(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, helpers.NewResponse("trajectory created"))
}

// updateTrajectory updates an existing trajectory.
// @Summary Update Trajectory
// @Tags trajectories
// @Description Updates an existing trajectory
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Trajectory ID"
// @Param input body requests.UpdateTrajectoryRequest true "Trajectory input"
// @Success 200 {object} entities.Trajectory
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/trajectories/{id} [put]
func (h *Handler) updateTrajectory(c *gin.Context) {
	var inp requests.UpdateTrajectoryRequest
	var err error

	if err := c.BindJSON(&inp.Body); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, types.ErrInvalidInputBody.Error())
		return
	}
	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}

	trajectory, err := h.services.Trajectories.UpdateTrajectory(c.Request.Context(), &inp)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, trajectory)
}

// deleteTrajectory deletes an existing trajectory.
// @Summary Delete Trajectory
// @Tags trajectories
// @Description Deletes an existing trajectory
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Trajectory ID"
// @Success 200 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /api/v1/trajectories/{id} [delete]
func (h *Handler) deleteTrajectory(c *gin.Context) {
	var inp requests.DeleteTrajectoryRequest
	var err error

	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}
	if err = h.services.Trajectories.DeleteTrajectory(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, helpers.NewResponse("trajectory deleted"))
}

// getTrajectoryByID retrieves a trajectory by its ID.
// @Summary Get Trajectory by ID
// @Tags trajectories
// @Description Retrieves a trajectory by its ID
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Trajectory ID"
// @Success 200 {object} entities.Trajectory
// @Failure 500 {object} helpers.Response
// @Router /api/v1/trajectories/{id} [get]
func (h *Handler) getTrajectoryByID(c *gin.Context) {
	var inp requests.GetTrajectoryByIDRequest
	var err error
	var trajectory *entities.Trajectory

	if inp.ID, err = h.validateRequestIDParam(c, values.IdQueryParam); err != nil {
		return
	}
	if trajectory, err = h.services.Trajectories.GetTrajectoryByID(c.Request.Context(), &inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, trajectory)
}
