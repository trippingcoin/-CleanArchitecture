package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	_ "os"
	rpc "user_service/internal/delivery/grpc"
	"user_service/internal/repository/postgres"
	"user_service/internal/usecase"
	"user_service/proto/userpb"

	_ "github.com/lib/pq" // Postgres driver
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {

	// Set up the Postgres connection
	db, err := sql.Open("postgres", "postgres://postgres:redmi@localhost:5433/postgres?sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize repositories and use cases
	userRepo := postgres.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(*userRepo)

	// Initialize the gRPC server and register the User Service
	grpcServer := grpc.NewServer()
	userServiceServer := rpc.NewUserServiceServer(userUsecase)
	userpb.RegisterUserServiceServer(grpcServer, userServiceServer)

	// Register reflection service on gRPC server (optional, for testing with gRPC CLI)
	reflection.Register(grpcServer)

	// Start the gRPC server
	port := ":8083" // You can change the port if necessary
	fmt.Printf("Starting gRPC server on port %s...\n", port)

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start listener: %v", err)
	}

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}
