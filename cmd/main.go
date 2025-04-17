package main

import (
	"log"
	"os"

	"CLEANARCHITECTURE/internal/router"

	"google.golang.org/grpc"
)

func main() {
	invConn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to Inventory gRPC: %v", err)
		os.Exit(1)
	}
	defer invConn.Close()

	orderConn, err := grpc.Dial("localhost:8082", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to Order gRPC: %v", err)
		os.Exit(1)
	}
	defer orderConn.Close()

	userConn, err := grpc.Dial("localhost:8083", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to User gRPC: %v", err)
		os.Exit(1)
	}
	defer userConn.Close()

	reviewConn, err := grpc.Dial("localhost:8084", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to Order gRPC: %v", err)
		os.Exit(1)
	}
	defer reviewConn.Close()

	r := router.SetupRoutes(invConn, orderConn, userConn, reviewConn, "superSecret")

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
}
