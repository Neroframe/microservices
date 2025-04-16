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

func CreatePayment(c *gin.Context) {
	var req orderpb.CreatePaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Error binding payment data: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment data"})
		return
	}

	log.Printf("Creating payment for order ID: %s", req.OrderId)

	resp, err := client.Payment.CreatePayment(context.Background(), &req)
	if err != nil {
		st, _ := status.FromError(err)
		log.Printf("Error creating payment: %v", st.Message())
		c.JSON(http.StatusInternalServerError, gin.H{"error": st.Message()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func GetPaymentByID(c *gin.Context) {
	id := c.Param("id")
	log.Printf("Fetching payment by ID: %s", id)

	resp, err := client.Payment.GetPaymentByID(context.Background(), &orderpb.GetPaymentRequest{
		PaymentId: id,
	})
	if err != nil {
		st, _ := status.FromError(err)
		log.Printf("Error fetching payment: %v", st.Message())
		c.JSON(http.StatusInternalServerError, gin.H{"error": st.Message()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
