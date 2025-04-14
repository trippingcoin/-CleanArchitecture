package main

import (
	"context"
	"log"
	"net/http"

	pb "proto/inventorypb"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewInventoryServiceClient(conn)
	r := gin.Default()

	r.POST("/api/products", func(c *gin.Context) {
		var req struct {
			Name     string  `json:"name"`
			Category string  `json:"category"`
			Price    float32 `json:"price"`
			Stock    int32   `json:"stock"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res, err := client.CreateProduct(context.Background(), &pb.CreateProductRequest{
			Name:     req.Name,
			Category: req.Category,
			Price:    req.Price,
			Stock:    req.Stock,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res)
	})

	// другие маршруты можно добавить позже (GET, PUT, DELETE)
	r.Run(":8080")
}
