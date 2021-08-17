package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	logger *log.Logger
	db     *gorm.DB
}

func NewHandler(logger *log.Logger, db *gorm.DB) *Handler {
	return &Handler{
		logger: logger,
		db:     db,
	}
}

func (h *Handler) SetupRoutes(r *gin.RouterGroup) {
	r.GET("/ping", h.Ping)
	r.GET("/customer", h.GetCustomerList)
}
