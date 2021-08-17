package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rsbh/ecom/entity"
)

type GetCustomerListResponse struct {
	Data []entity.Customer `json:"data"`
}

// GetCustomerList godoc
// @Summary Return List of customers
// @Accept  json
// @Success 200 {object} GetCustomerListResponse
// @Router /api/customer [get]
func (h *Handler) GetCustomerList(c *gin.Context) {
	var customers []entity.Customer
	h.db.Find(&customers)
	c.JSON(http.StatusOK, GetCustomerListResponse{customers})
}
