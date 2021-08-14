package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rsbh/ecom/config"
	"github.com/rsbh/ecom/database"
	"github.com/rsbh/ecom/router"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	c := config.LoadConfig(logger)
	db := database.Connect(c.Database)
	database.AutoMigrate(db)
	r := router.InitRouter(logger)
	logger.Fatal(r.Run(fmt.Sprintf("%s:%d", c.Server.Address, c.Server.Port)))
}
