package main

import (
	"database/sql"
	"log"
	"net"

	rpc "review_service/internal/delivery/grpc"
	"review_service/internal/repository/postgres"
	"review_service/internal/usecase"
	"review_service/proto/reviewpb"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
	"google.golang.org/grpc"
)

func main() {
	// Set up the Postgres connection
	db, err := sql.Open("postgres", "postgres://postgres:redmi@localhost:5433/postgres?sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Verify that the database connection is established
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	// Initialize the repository and usecase
	repo := postgres.NewReviewRepository(db) // Use the actual db init
	uc := usecase.NewReviewUsecase(repo)

	// Set up the gRPC server
	srv := rpc.NewReviewServiceServer(uc)

	// Create a listener on port 50054
	lis, err := net.Listen("tcp", ":8084")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	reviewpb.RegisterReviewServiceServer(grpcServer, srv)

	// Log server startup
	log.Println("Order Service gRPC server started on port 8084")

	// Serve the gRPC server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
