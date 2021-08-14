package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rsbh/ecom/config"
)

func main() {
	config.LoadConfig()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})
	log.Fatal(r.Run(":8080"))
}
