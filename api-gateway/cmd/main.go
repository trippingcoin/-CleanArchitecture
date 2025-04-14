package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	invpb "api-gateway/proto/inventorypb"
	ordpb "api-gateway/proto/orderpb"

	"api-gateway/internal/handler"
)

func main() {
	r := gin.Default()

	// Connect to Inventory gRPC
	invConn, err := grpc.Dial("inventory-service:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Inventory connection failed: %v", err)
	}
	invClient := invpb.NewInventoryServiceClient(invConn)

	// Connect to Order gRPC
	ordConn, err := grpc.Dial("order-service:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Order connection failed: %v", err)
	}
	ordClient := ordpb.NewOrderServiceClient(ordConn)

	// Register Routes
	handler.RegisterInventoryRoutes(r, invClient)
	handler.RegisterOrderRoutes(r, ordClient)

	log.Println("API Gateway running on :8080")
	r.Run(":8080")
}
