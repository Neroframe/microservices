package client

import (
	inventorypb "github.com/Neroframe/ecommerce-platform/api-gateway/proto"

	"google.golang.org/grpc"
)

var Inventory inventorypb.InventoryServiceClient

func InitInventoryClient(conn *grpc.ClientConn) {
	Inventory = inventorypb.NewInventoryServiceClient(conn)
}
