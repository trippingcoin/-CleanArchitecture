package grpc

import (
	"context"
	"log"

	pb "github.com/trippingcoin/-CleanArchitecture/review_service/proto/"

	"google.golang.org/grpc"
)

type StatisticsClient struct {
	client pb.StatisticsServiceClient
}

func NewStatisticsClient(conn *grpc.ClientConn) *StatisticsClient {
	return &StatisticsClient{
		client: pb.NewStatisticsServiceClient(conn),
	}
}

func (c *StatisticsClient) GetUserStatistics(userID string) (*pb.UserStatisticsResponse, error) {
	resp, err := c.client.GetUserStatistics(context.Background(), &pb.UserStatisticsRequest{
		UserId: userID,
	})
	if err != nil {
		log.Printf("Error calling GetUserStatistics: %v", err)
		return nil, err
	}
	return resp, nil
}
