package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rsbh/ecom/config"
	"github.com/rsbh/ecom/database"
)

func main() {
	c := config.LoadConfig()
	database.Connect(c.Database)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})
	log.Fatal(r.Run(":8080"))
}
