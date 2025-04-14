package main

import (
	"database/sql"
	"log"
	"net"
	"os"

	grpcH "order-service/internal/delivery/grpc"
	"order-service/internal/repository"
	"order-service/internal/service"
	pb "order-service/order-service/proto"

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
	svc := service.NewOrderService(repo)
	handler := grpcH.NewOrderHandler(svc)

	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterOrderServiceServer(server, handler)

	log.Println("OrderService is running on port 50052")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
