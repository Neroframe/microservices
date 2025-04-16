package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/Neroframe/ecommerce-platform/api-gateway/internal/client"
	orderpb "github.com/Neroframe/ecommerce-platform/api-gateway/proto/order"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

func CreateOrder(c *gin.Context) {
	var req orderpb.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Error binding order data: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order data"})
		return
	}

	log.Printf("Creating order for user: %s", req.UserId)

	resp, err := client.Order.CreateOrder(context.Background(), &req)
	if err != nil {
		st, _ := status.FromError(err)
		log.Printf("Error creating order: %v", st.Message())
		c.JSON(http.StatusInternalServerError, gin.H{"error": st.Message()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func GetOrderByID(c *gin.Context) {
	id := c.Param("id")
	log.Printf("Fetching order by ID: %s", id)

	resp, err := client.Order.GetOrderByID(context.Background(), &orderpb.GetOrderRequest{Id: id})
	if err != nil {
		st, _ := status.FromError(err)
		log.Printf("Error fetching order: %v", st.Message())
		c.JSON(http.StatusInternalServerError, gin.H{"error": st.Message()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func UpdateOrderStatus(c *gin.Context) {
	var req orderpb.UpdateOrderStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Error binding update status: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status data"})
		return
	}

	log.Printf("Updating order status for ID: %s", req.Id)

	resp, err := client.Order.UpdateOrderStatus(context.Background(), &req)
	if err != nil {
		st, _ := status.FromError(err)
		log.Printf("Error updating order: %v", st.Message())
		c.JSON(http.StatusInternalServerError, gin.H{"error": st.Message()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func ListUserOrders(c *gin.Context) {
	userId := c.Query("user_id")
	log.Printf("Listing orders for user: %s", userId)

	resp, err := client.Order.ListUserOrders(context.Background(), &orderpb.ListOrdersRequest{UserId: userId})
	if err != nil {
		st, _ := status.FromError(err)
		log.Printf("Error listing orders: %v", st.Message())
		c.JSON(http.StatusInternalServerError, gin.H{"error": st.Message()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
