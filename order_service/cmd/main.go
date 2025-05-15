package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"time"

	"order_service/internal/cache"
	rpc "order_service/internal/delivery/grpc"
	"order_service/internal/repository/postgres"
	u "order_service/internal/usecase"
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
	uc := u.NewOrderUsecase(repo, nc)

	cache.InitCache()
	uc.InitCache()

	ctx := context.Background()
	orders, err := uc.ListOrders(ctx, "1")
	if err != nil {
		log.Printf("Failed to preload orders into cache: %v", err)
	} else {
		cache.OrderCache.Set("order_list", orders, cache.DefaultExpiration)
		for _, o := range orders {
			cache.OrderCache.Set(
				fmt.Sprintf("order_%s", o.ID),
				o,
				cache.DefaultExpiration,
			)
		}
		log.Println("Cache preloaded with order data")
	}

	go func() {
		ticker := time.NewTicker(12 * time.Hour)
		for range ticker.C {
			log.Println("Refreshing order cache...")
			orders, err := uc.ListOrders(context.Background(), "")
			if err != nil {
				log.Printf("Cache refresh failed: %v", err)
				continue
			}
			cache.OrderCache.Set("order_list", orders, cache.DefaultExpiration)
			for _, o := range orders {
				cache.OrderCache.Set(
					fmt.Sprintf("order_%s", o.ID),
					o,
					cache.DefaultExpiration,
				)
			}
			log.Println("Order cache successfully refreshed")
		}
	}()

	srv := rpc.NewOrderServiceServer(uc)

	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	orderpb.RegisterOrderServiceServer(grpcServer, srv)

	log.Println("Order Service gRPC server started on port 8082")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
