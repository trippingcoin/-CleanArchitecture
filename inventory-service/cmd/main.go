package main

import (
	"database/sql"
	"log"
	"net"
	"os"

	grpcH "inventory-service/internal/delivery/grpc"
	"inventory-service/internal/repository"
	"inventory-service/internal/service"
	pb "inventory-service/inventory-service/proto"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer db.Close()

	repo := repository.NewPostgresRepo(db)
	svc := service.NewInventoryService(repo)
	handler := grpcH.NewInventoryHandler(svc)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterInventoryServiceServer(grpcServer, handler)

	log.Println("InventoryService running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
