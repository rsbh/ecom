package database

import (
	"fmt"
	"log"
	"net/url"

	"github.com/rsbh/ecom/config"
	"github.com/rsbh/ecom/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&entity.Customer{})
}

func Connect(config config.DatabaseConfig) *gorm.DB {
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
