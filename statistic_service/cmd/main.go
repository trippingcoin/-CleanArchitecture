package main

import (
	"database/sql"
	"log"
	"net"
	ser "statistics_service/internal/delivery/grpc"
	"statistics_service/internal/repository/postgres"
	"statistics_service/internal/subscriber"
	"statistics_service/internal/usecase"
	pb "statistics_service/proto/statisticspb"

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

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	defer nc.Close()

	repo := postgres.NewPostgresRepo(db)
	uc := usecase.NewStatisticsUsecase(repo, nc)
	uc.PublishHourlyStats()

	go subscriber.StartNATSSubscriber(uc)

	lis, err := net.Listen("tcp", ":8085")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterStatisticsServiceServer(grpcServer, ser.NewStatisticsHandler(uc))

	log.Println("Statistics Service is running on :8085")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
