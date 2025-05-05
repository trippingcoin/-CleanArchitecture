package rest

import (
	"CLEANARCHITECTURE/pkg/proto/statisticspb"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StatisticsHandler struct {
	client statisticspb.StatisticsServiceClient
}

func NewStatisticsHandler(client statisticspb.StatisticsServiceClient) *StatisticsHandler {
	return &StatisticsHandler{client: client}
}

// GetUserStatistics fetches comprehensive user statistics.
func (h *StatisticsHandler) GetUserStatistics(c *gin.Context) {
	userID := c.Param("id")

	resp, err := h.client.GetUserStatistics(c.Request.Context(), &statisticspb.UserStatisticsRequest{
		UserId: userID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetUserOrdersStatistics fetches user order statistics.
func (h *StatisticsHandler) GetUserOrdersStatistics(c *gin.Context) {
	userID := c.Param("id")

	resp, err := h.client.GetUserOrdersStatistics(c.Request.Context(), &statisticspb.UserOrderStatisticsRequest{
		UserId: userID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
