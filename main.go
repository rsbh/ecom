package main

import (
	"fmt"
	"log"

	"github.com/rsbh/ecom/config"
	"github.com/rsbh/ecom/database"
	"github.com/rsbh/ecom/router"
)

func main() {
	c := config.LoadConfig()
	db := database.Connect(c.Database)
	database.AutoMigrate(db)
	r := router.InitRouter()
	log.Fatal(r.Run(fmt.Sprintf("%s:%d", c.Server.Address, c.Server.Port)))
}
