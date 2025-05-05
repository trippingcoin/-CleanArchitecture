package grpc

import (
	"CLEANARCHITECTURE/pkg/proto/statisticspb"
	"context"
	"log"

	"google.golang.org/grpc"
)

type StatisticsClient struct {
	client statisticspb.StatisticsServiceClient
}

// NewStatisticsClient initializes a new gRPC client for the StatisticsService.
func NewStatisticsClient(conn *grpc.ClientConn) statisticspb.StatisticsServiceClient {
	return statisticspb.NewStatisticsServiceClient(conn)
}

// GetUserStatistics fetches comprehensive user statistics.
func (c *StatisticsClient) GetUserStatistics(userID string) (*statisticspb.UserStatisticsResponse, error) {
	resp, err := c.client.GetUserStatistics(context.Background(), &statisticspb.UserStatisticsRequest{
		UserId: userID,
	})
	if err != nil {
		log.Printf("Error calling GetUserStatistics: %v", err)
		return nil, err
	}
	return resp, nil
}

// GetUserOrdersStatistics fetches user order statistics.
func (c *StatisticsClient) GetUserOrdersStatistics(userID string) (*statisticspb.UserOrderStatisticsResponse, error) {
	resp, err := c.client.GetUserOrdersStatistics(context.Background(), &statisticspb.UserOrderStatisticsRequest{
		UserId: userID,
	})
	if err != nil {
		log.Printf("Error calling GetUserOrdersStatistics: %v", err)
		return nil, err
	}
	return resp, nil
}
