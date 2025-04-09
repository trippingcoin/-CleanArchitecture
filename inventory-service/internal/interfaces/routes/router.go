package routes

import (
	"database/sql"
	"inventory-service/internal/interfaces/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	productCtrl := controllers.NewProductController(db)

	r.GET("/products", productCtrl.List)
	r.POST("/products", productCtrl.Create)
	// Add others: GET /:id, PATCH /:id, DELETE /:id

	return r
}
