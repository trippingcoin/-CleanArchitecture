package rest

import (
	"fmt"
	"net/http"

	"github.com/trippingcoin/-CleanArchitecture/user_service/proto/userpb"

	"CLEANARCHITECTURE/internal/grpc"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	client grpc.UserGRPCClient
}

func NewUserHandler(client grpc.UserGRPCClient) *UserHandler {
	return &UserHandler{client: client}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var req userpb.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.client.RegisterUser(c, &req)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

func (h *UserHandler) AuthenticateUser(c *gin.Context) {
	var req userpb.AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.client.AuthenticateUser(c, &req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Return only the token
	c.JSON(http.StatusOK, gin.H{"token": res.GetToken()})
}

func (h *UserHandler) GetUserProfile(c *gin.Context) {
	id := c.Param("id")

	res, err := h.client.GetUserProfile(c, &userpb.UserID{UserId: parseID(id)})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func parseID(s string) int32 {
	var id int
	fmt.Sscan(s, &id)
	return int32(id)
}
