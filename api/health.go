package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingResponse struct {
	Message string `json:"message"`
}

// Ping godoc
// @Summary Health route
// @Accept  json
// @Success 200 {object} PingResponse
// @Router /api/ping [get]
func (h *Handler) Ping(c *gin.Context) {
	resp := PingResponse{"pong"}
	c.JSON(http.StatusOK, resp)
}
