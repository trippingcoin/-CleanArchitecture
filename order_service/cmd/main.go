package main

import (
	"database/sql"
	"log"
	"net"

	rpc "order_service/internal/delivery/grpc"
	"order_service/internal/repository/postgres"
	"order_service/internal/usecase"
	"order_service/proto/orderpb"

	_ "github.com/lib/pq"
	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:redmi@localhost:5433/postgres?sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	// Connect to NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("failed to connect to NATS: %v", err)
	}
	defer nc.Close()
	log.Println("Order Service Nats server started")

	// Initialize the repository and usecase
	repo := postgres.NewOrderRepository(db)
	uc := usecase.NewOrderUsecase(repo, nc)

	// Set up the gRPC server
	srv := rpc.NewOrderServiceServer(uc)

	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	orderpb.RegisterOrderServiceServer(grpcServer, srv)

	log.Println("Order Service gRPC server started on port 8082")

	// Serve the gRPC server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
