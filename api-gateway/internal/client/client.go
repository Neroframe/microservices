package client

import (
	inventorypb "github.com/Neroframe/ecommerce-platform/api-gateway/proto/inventory"
	orderpb "github.com/Neroframe/ecommerce-platform/api-gateway/proto/order"

	"google.golang.org/grpc"
)

var (
	Inventory inventorypb.InventoryServiceClient
	Order     orderpb.OrderServiceClient
	Payment   orderpb.PaymentServiceClient
)

func InitInventoryClient(conn *grpc.ClientConn) {
	Inventory = inventorypb.NewInventoryServiceClient(conn)
}

func InitOrderClient(conn *grpc.ClientConn) {
	Order = orderpb.NewOrderServiceClient(conn)
	Payment = orderpb.NewPaymentServiceClient(conn)
}
