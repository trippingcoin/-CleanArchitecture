package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"time"

	"inventory_service/internal/cache"
	rpc "inventory_service/internal/delivery/grpc"
	"inventory_service/internal/repository/postgres"
	"inventory_service/internal/usecase"
	"inventory_service/proto/inventorypb"

	_ "github.com/lib/pq"
	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
)

func main() {
	// Set up the Postgres connection
	db, err := sql.Open("postgres", "postgres://postgres:redmi@localhost:5433/postgres?sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Connect to NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	defer nc.Close()

	// Initialize repository and usecase
	repo := postgres.NewProductRepository(db)
	uc := usecase.NewProductUsecase(repo, nc)

	// 🟡 Initialize cache
	cache.InitCache()

	ctx := context.Background()
	products, err := uc.List(ctx)
	if err != nil {
		log.Printf("Failed to preload products into cache: %v", err)
	} else {
		cache.InventoryCache.Set("product_list", products, cache.DefaultExpiration)
		for _, p := range products {
			cache.InventoryCache.Set(
				fmt.Sprintf("product_%d", p.ID),
				p,
				cache.DefaultExpiration,
			)
		}
		log.Println("Cache preloaded with product data")
	}

	go func() {
		ticker := time.NewTicker(12 * time.Hour)
		for range ticker.C {
			log.Println("Refreshing cache...")
			products, err := uc.List(context.Background())
			if err != nil {
				log.Printf("Cache refresh failed: %v", err)
				continue
			}
			cache.InventoryCache.Set("product_list", products, cache.DefaultExpiration)
			for _, p := range products {
				cache.InventoryCache.Set(
					fmt.Sprintf("product_%d", p.ID),
					p,
					cache.DefaultExpiration,
				)
			}
			log.Println("Cache successfully refreshed")
		}
	}()

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
