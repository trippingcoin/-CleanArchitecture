package rest

import (
	"context"
	"net/http"
	"strconv"

	"CLEANARCHITECTURE/internal/grpc"

	"github.com/gin-gonic/gin"
	"github.com/trippingcoin/-CleanArchitecture/inventory_service/proto/inventorypb"
)

type ProductHandler struct {
	GrpcClient grpc.InventoryGRPCClient
}

func NewProductHandler(client grpc.InventoryGRPCClient) *ProductHandler {
	return &ProductHandler{GrpcClient: client}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req inventorypb.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := h.GrpcClient.CreateProduct(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	product, err := h.GrpcClient.GetProduct(context.Background(), &inventorypb.GetProductRequest{
		ProductId: int32(id),
	})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) ListProducts(c *gin.Context) {
	res, err := h.GrpcClient.ListProducts(context.Background(), &inventorypb.ListProductsRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res.Products)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	var req inventorypb.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.ProductId = int32(id)

	product, err := h.GrpcClient.UpdateProduct(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	_, err := h.GrpcClient.DeleteProduct(context.Background(), &inventorypb.DeleteProductRequest{
		ProductId: int32(id),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
