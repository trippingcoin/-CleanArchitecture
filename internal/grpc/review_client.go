package grpc

import (
	"github.com/trippingcoin/-CleanArchitecture/review_service/proto/reviewpb"
	"google.golang.org/grpc"
)

type ReviewGRPCClient interface {
	reviewpb.ReviewServiceClient
}

func NewReviewGRPCClient(conn *grpc.ClientConn) ReviewGRPCClient {
	return reviewpb.NewReviewServiceClient(conn)
}
