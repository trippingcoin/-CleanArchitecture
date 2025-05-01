package server

import (
	"log"
	"net"
	grpcadapter "statistics_service/internal/adapter/grpc"
	"statistics_service/internal/app"

	"google.golang.org/grpc"
)

type StatisticsServer struct {
	grpcServer *grpc.Server
}

func NewServer(app *app.StatisticsApp) *StatisticsServer {
	grpcServer := grpc.NewServer()

	grpcadapter.RegisterHandlers(grpcServer, app)

	return &StatisticsServer{
		grpcServer: grpcServer,
	}
}

func (s *StatisticsServer) Start(address string) error {
	// Listen on the given address
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", address, err)
		return err
	}

	// Start the gRPC server
	log.Printf("Server is starting on %s...", address)
	if err := s.grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to start the gRPC server: %v", err)
		return err
	}

	return nil
}

func (s *StatisticsServer) Stop() {
	// Gracefully stop the server
	s.grpcServer.GracefulStop()
}
