package router

import (
	"Gym-Management-System/internal/delivery/rest"
	grpcclient "Gym-Management-System/internal/grpc"
	"Gym-Management-System/internal/middleware"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// Now takes three gRPC connections + JWT secret
func SetupRoutes(
	invConn *grpc.ClientConn,
	orderConn *grpc.ClientConn,
	userConn *grpc.ClientConn,
	jwtSecret string,
) *gin.Engine {
	r := gin.Default()

	// Instantiate gRPC clients
	invClient := grpcclient.NewInventoryGRPCClient(invConn)
	orderClient := grpcclient.NewOrderGRPCClient(orderConn)
	userClient := grpcclient.NewUserGRPCClient(userConn)

	// Instantiate REST handlers
	invH := rest.NewProductHandler(invClient)
	orderH := rest.NewOrderHandler(orderClient)
	userH := rest.NewUserHandler(userClient)

	// Public routes (no auth)
	r.POST("/users/register", userH.RegisterUser)
	r.POST("/users/authenticate", userH.AuthenticateUser)

	// Protected routes
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware(jwtSecret))

	// Inventory
	protected.POST("/products", invH.CreateProduct)
	protected.GET("/products", invH.ListProducts)
	protected.GET("/products/:id", invH.GetProduct)
	protected.PATCH("/products/:id", invH.UpdateProduct)
	protected.DELETE("/products/:id", invH.DeleteProduct)

	// Orders
	protected.POST("/orders", orderH.CreateOrder)
	protected.GET("/orders", orderH.ListOrders)
	protected.GET("/orders/:id", orderH.GetOrder)
	protected.PATCH("/orders/:id", orderH.UpdateOrderStatus)

	// User profile (protected)
	protected.GET("/users/:id", userH.GetUserProfile)

	return r
}
