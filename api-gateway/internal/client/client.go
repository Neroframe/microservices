package client

import (
	inventorypb "github.com/Neroframe/ecommerce-platform/api-gateway/proto/inventory"
	orderpb "github.com/Neroframe/ecommerce-platform/api-gateway/proto/order"
	statisticspb "github.com/Neroframe/ecommerce-platform/api-gateway/proto/statistics"

	"google.golang.org/grpc"
)

var (
	Inventory  inventorypb.InventoryServiceClient
	Order      orderpb.OrderServiceClient
	Payment    orderpb.PaymentServiceClient
	Statistics statisticspb.StatisticsServiceClient
)

func InitInventoryClient(conn *grpc.ClientConn) {
	Inventory = inventorypb.NewInventoryServiceClient(conn)
}

func InitOrderClient(conn *grpc.ClientConn) {
	Order = orderpb.NewOrderServiceClient(conn)
	Payment = orderpb.NewPaymentServiceClient(conn)
}

func InitStatisticsClient(conn *grpc.ClientConn) {
	Statistics = statisticspb.NewStatisticsServiceClient(conn)
}
