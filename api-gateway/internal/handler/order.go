package handler

import (
	"context"
	"net/http"

	ordpb "api-gateway/proto/orderpb"

	"github.com/gin-gonic/gin"
)

func RegisterOrderRoutes(r *gin.Engine, client ordpb.OrderServiceClient) {
	r.POST("/orders", func(c *gin.Context) {
		var req ordpb.CreateOrderRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, err := client.CreateOrder(context.Background(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res)
	})

	r.GET("/orders/:id", func(c *gin.Context) {
		id := c.Param("id")
		res, err := client.GetOrderByID(context.Background(), &ordpb.GetOrderRequest{Id: id})
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res)
	})
}
