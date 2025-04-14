package handler

import (
	"context"
	"net/http"

	invpb "api-gateway/proto/inventorypb"

	"github.com/gin-gonic/gin"
)

func RegisterInventoryRoutes(r *gin.Engine, client invpb.InventoryServiceClient) {
	r.POST("/products", func(c *gin.Context) {
		var req invpb.CreateProductRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, err := client.CreateProduct(context.Background(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res)
	})

	r.GET("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		res, err := client.GetProductByID(context.Background(), &invpb.GetProductRequest{Id: id})
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res)
	})
}
