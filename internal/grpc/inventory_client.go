package grpc

import (
	"github.com/trippingcoin/-CleanArchitecture/inventory_service/proto/inventorypb"
	"google.golang.org/grpc"
)

type InventoryGRPCClient interface {
	inventorypb.InventoryServiceClient
}

func NewInventoryGRPCClient(conn *grpc.ClientConn) InventoryGRPCClient {
	return inventorypb.NewInventoryServiceClient(conn)
}
