package http

import "github.com/gin-gonic/gin"

type Handler struct {

}

func (h *Handler) Init() *gin.Engine {
	// Init Gin Handler
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	return router
}