package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/rsbh/ecom/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectToDB(config config.DatabaseConfig) *gorm.DB {
	log.Println("Connecting to database...")
	dsn := url.URL{
		User:   url.UserPassword(config.Username, config.Password),
		Scheme: "postgres",
		Host:   fmt.Sprintf("%s:%d", config.Host, config.Port),
		Path:   config.Database,
	}
	db, err := gorm.Open(postgres.Open(dsn.String()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	log.Println("Connected to database...")

	return db
}

func main() {
	c := config.LoadConfig()
	connectToDB(c.Database)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})
	log.Fatal(r.Run(":8080"))
}
