package grpc

import (
	"github.com/trippingcoin/-CleanArchitecture/review_service/proto/orderpb"
	"google.golang.org/grpc"
)

type ReviewGRPCClient interface {
	reviewpb.OrderServiceClient
}

func NewOrderGRPCClient(conn *grpc.ClientConn) OrderGRPCClient {
	return orderpb.NewOrderServiceClient(conn)
}
