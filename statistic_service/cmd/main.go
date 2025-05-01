package main

import (
	"log"
	"statistics_service/internal/app"
	"statistics_service/internal/repository"
	"statistics_service/internal/server"
)

func main() {
	repo := repository.NewInMemoryStatisticsRepository()
	statisticsApp := app.New(repo)

	s := server.NewServer(statisticsApp)

	if err := s.Start(":50051"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
