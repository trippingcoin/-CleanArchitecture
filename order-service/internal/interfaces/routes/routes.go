package routes

import (
	"database/sql"
	"order-service/internal/interfaces/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	orderCtrl := controllers.NewOrderController(db)

	r.POST("/orders", orderCtrl.CreateOrder)
	r.GET("/orders/:id", orderCtrl.GetOrder)
	r.PATCH("/orders/:id", orderCtrl.UpdateStatus)
	r.GET("/orders", orderCtrl.ListOrders)

	return r
}
