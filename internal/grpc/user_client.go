package grpc

import (
	"/user_service/proto/userpb"

	"google.golang.org/grpc"
)

type UserGRPCClient interface {
	userpb.UserServiceClient
}

func NewUserGRPCClient(conn *grpc.ClientConn) UserGRPCClient {
	return userpb.NewUserServiceClient(conn)
}
