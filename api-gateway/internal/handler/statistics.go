package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/Neroframe/ecommerce-platform/api-gateway/internal/client"
	statpb "github.com/Neroframe/ecommerce-platform/api-gateway/proto/statistics"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

func GetUserOrdersStatistics(c *gin.Context) {
	userId := c.Param("userId")

	resp, err := client.Statistics.GetUserOrdersStatistics(context.Background(), &statpb.UserOrderStatisticsRequest{
		UserId: userId,
	})
	if err != nil {
		st, _ := status.FromError(err)
		log.Printf("Error fetching user order stats: %v", st.Message())
		c.JSON(http.StatusInternalServerError, gin.H{"error": st.Message()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func GetUserStatistics(c *gin.Context) {
	resp, err := client.Statistics.GetUserStatistics(context.Background(), &statpb.UserStatisticsRequest{})
	if err != nil {
		st, _ := status.FromError(err)
		log.Printf("Error fetching user stats: %v", st.Message())
		c.JSON(http.StatusInternalServerError, gin.H{"error": st.Message()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
