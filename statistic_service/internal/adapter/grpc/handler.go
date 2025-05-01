package grpcadapter

import (
	"context"
	"fmt"
	"statistics_service/internal/app"
	"statistics_service/internal/domain"
	pb "statistics_service/proto/statisticspb"

	"log"

	"google.golang.org/grpc"
)

type StatisticsHandler struct {
	pb.UnimplementedStatisticsServiceServer
	app *app.StatisticsApp
}

func RegisterHandlers(s *grpc.Server, app *app.StatisticsApp) {
	pb.RegisterStatisticsServiceServer(s, &StatisticsHandler{app: app})
}

func (h *StatisticsHandler) GetUserStatistics(ctx context.Context, req *pb.UserStatisticsRequest) (*pb.UserStatisticsResponse, error) {
	stats, err := h.app.GetStatistics(req.UserId)
	if err != nil {
		log.Printf("Error retrieving statistics for user %s: %v", req.UserId, err)
		return nil, err
	}

	s, ok := stats.(*domain.UserStatistics)
	if !ok {
		log.Printf("Unexpected type for statistics: %T", stats)
		return nil, fmt.Errorf("unexpected type for statistics")
	}

	return &pb.UserStatisticsResponse{
		UserId:      s.UserID,
		TotalOrders: int32(s.TotalOrders),
		TotalSpent:  float32(s.TotalSpent),
		PeakHour:    fmt.Sprintf("%d", s.PeakOrderHour),
	}, nil
}
