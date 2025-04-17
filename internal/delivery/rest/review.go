package rest

import (
	"context"
	"net/http"

	"CLEANARCHITECTURE/internal/grpc"

	"github.com/gin-gonic/gin"
	"github.com/trippingcoin/-CleanArchitecture/review_service/proto/reviewpb"
)

type ReviewHandler struct {
	GrpcClient grpc.ReviewGRPCClient
}

func NewReviewHandler(client grpc.ReviewGRPCClient) *ReviewHandler {
	return &ReviewHandler{GrpcClient: client}
}

func (h *ReviewHandler) CreateReview(c *gin.Context) {
	var req reviewpb.CreateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	review, err := h.GrpcClient.CreateReview(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, review)
}

func (h *ReviewHandler) GetReview(c *gin.Context) {
	idStr := c.Param("id")

	review, err := h.GrpcClient.GetReview(context.Background(), &reviewpb.GetReviewRequest{
		Id: idStr,
	})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, review)
}

func (h *ReviewHandler) ListReviews(c *gin.Context) {
	res, err := h.GrpcClient.ListReviews(context.Background(), &reviewpb.ListReviewsRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res.Reviews)
}

func (h *ReviewHandler) UpdateReview(c *gin.Context) {
	idStr := c.Param("id")

	var req reviewpb.UpdateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Id = idStr

	review, err := h.GrpcClient.UpdateReview(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, review)
}
