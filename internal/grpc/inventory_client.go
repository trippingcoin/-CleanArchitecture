package grpc

import (
	"./inventory_service/proto/inventorypb"
	"google.golang.org/grpc"
)

type InventoryGRPCClient interface {
	inventorypb.InventoryServiceClient
}

func NewInventoryGRPCClient(conn *grpc.ClientConn) InventoryGRPCClient {
	return inventorypb.NewInventoryServiceClient(conn)
}
