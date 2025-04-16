package main

import (
	"database/sql"
	"log"
	"net"

	rpc "inventory_service/internal/delivery/grpc"
	"inventory_service/internal/repository/postgres"
	"inventory_service/internal/usecase"
	"inventory_service/proto/inventorypb"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {

	// Set up the Postgres connection
	db, err := sql.Open("postgres", "postgres://postgres:redmi@localhost:5433/postgres?sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	repo := postgres.NewProductRepository(db) // Assume initialized with DB connection
	uc := usecase.NewProductUsecase(repo)
	server := rpc.NewInventoryServer(uc)

	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	inventorypb.RegisterInventoryServiceServer(s, server)

	log.Println("gRPC Inventory Service started on port :8081")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
