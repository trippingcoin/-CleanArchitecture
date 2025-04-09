package main

import (
	"order-service/config"
	"order-service/internal/interfaces/routes"
)

func main() {
	db := config.InitDB()
	router := routes.SetupRouter(db)
	router.Run(":8002")
}
