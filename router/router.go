package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rsbh/ecom/router/api"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	apiRouter := r.Group("/api")
	apiRouter.GET("/ping", api.Ping)
	return r
}
