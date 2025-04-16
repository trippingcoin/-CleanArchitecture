package rest

import (
	"/order_service/proto/orderpb"
	"Gym-Management-System/internal/grpc"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	client grpc.OrderGRPCClient
}

func NewOrderHandler(client grpc.OrderGRPCClient) *OrderHandler {
	return &OrderHandler{client: client}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req orderpb.OrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	res, err := h.client.CreateOrder(c, &req)
	if err != nil {
		log.Printf("Error creating order: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"order_id": res.OrderId})
}

func (h *OrderHandler) GetOrder(c *gin.Context) {
	orderID := c.Param("id")

	order, err := h.client.GetOrder(c, &orderpb.GetOrderRequest{OrderId: orderID})
	if err != nil {
		log.Printf("Error getting order: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) UpdateOrderStatus(c *gin.Context) {
	orderID := c.Param("id")

	var req orderpb.UpdateOrderStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	req.OrderId = orderID

	_, err := h.client.UpdateOrderStatus(c, &req)
	if err != nil {
		log.Printf("Error updating order status: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Order status updated successfully"})
}

func (h *OrderHandler) ListOrders(c *gin.Context) {
	userID := c.DefaultQuery("user_id", "")

	orders, err := h.client.ListOrders(c, &orderpb.OrderListRequest{UserId: userID})
	if err != nil {
		log.Printf("Error listing orders: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}
