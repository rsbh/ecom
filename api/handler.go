package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	logger *log.Logger
}

func NewHandler(logger *log.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}

func (h *Handler) SetupRoutes(r *gin.RouterGroup) {
	r.GET("/ping", h.Ping)
}
