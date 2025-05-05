package ser

import (
	"context"
	"fmt"
	"log"
	"statistics_service/internal/usecase"
	pb "statistics_service/proto/statisticspb"
)

type StatisticsHandler struct {
	pb.UnimplementedStatisticsServiceServer
	uc *usecase.StatisticsUsecase
}

func NewStatisticsHandler(uc *usecase.StatisticsUsecase) *StatisticsHandler {
	return &StatisticsHandler{uc: uc}
}

// GetUserOrdersStatistics fetches user order statistics.
func (h *StatisticsHandler) GetUserOrdersStatistics(ctx context.Context, req *pb.UserOrderStatisticsRequest) (*pb.UserOrderStatisticsResponse, error) {
	stats, err := h.uc.GetUserStats(req.UserId)
	if err != nil {
		log.Printf("Error fetching user order stats for user %s: %v", req.UserId, err)
		return nil, err
	}

	return &pb.UserOrderStatisticsResponse{
		UserId:         stats.UserID,
		TotalOrders:    int32(stats.TotalOrders),
		MostActiveHour: fmt.Sprintf("%d", stats.MostActiveHour), // assuming this is the most active hour of the day
	}, nil
}

// GetUserStatistics fetches comprehensive user statistics.
func (h *StatisticsHandler) GetUserStatistics(ctx context.Context, req *pb.UserStatisticsRequest) (*pb.UserStatisticsResponse, error) {
	stats, err := h.uc.GetUserStats(req.UserId)
	if err != nil {
		log.Printf("Error fetching user statistics for user %s: %v", req.UserId, err)
		return nil, err
	}

	return &pb.UserStatisticsResponse{
		UserId:            stats.UserID,
		RegistrationDate:  stats.RegistrationDate.Format("2006-01-02"), // formatting date to string "YYYY-MM-DD"
		TotalOrders:       int32(stats.TotalOrders),
		AverageOrderValue: float32(stats.AvgOrderValue),           // assuming AvgOrderValue is float64
		TotalSpent:        float32(stats.TotalSpent),              // assuming TotalSpent is float64
		PeakHour:          fmt.Sprintf("%d", stats.PeakOrderHour), // formatting peak hour as a string
	}, nil
}
