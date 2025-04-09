package main

import (
	"inventory-service/config"
	"inventory-service/internal/interfaces/routes"
)

func main() {
	db := config.InitDB()
	router := routes.SetupRouter(db)
	router.Run(":8001")
}
