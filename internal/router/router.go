package router

import (
	"CLEANARCHITECTURE/internal/delivery/rest"
	grpcclient "CLEANARCHITECTURE/internal/grpc"
	"CLEANARCHITECTURE/internal/middleware"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func SetupRoutes(
	invConn *grpc.ClientConn,
	orderConn *grpc.ClientConn,
	userConn *grpc.ClientConn,
	reviewConn *grpc.ClientConn,
	statisticsConn *grpc.ClientConn,
	jwtSecret string,
) *gin.Engine {
	r := gin.Default()

	// Instantiate gRPC clients
	invClient := grpcclient.NewInventoryGRPCClient(invConn)
	orderClient := grpcclient.NewOrderGRPCClient(orderConn)
	userClient := grpcclient.NewUserGRPCClient(userConn)
	reviewClient := grpcclient.NewReviewGRPCClient(reviewConn)
	statisticsClient := grpcclient.NewStatisticsClient(statisticsConn)

	// Instantiate REST handlers
	invH := rest.NewProductHandler(invClient)
	orderH := rest.NewOrderHandler(orderClient)
	userH := rest.NewUserHandler(userClient)
	reviewH := rest.NewReviewHandler(reviewClient)
	statisticsH := rest.NewStatisticsHandler(statisticsClient)

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

	// User profile
	protected.GET("/users/:id", userH.GetUserProfile)

	// Statistics
	protected.GET("/users/:id/statistics", statisticsH.GetUserStatistics)
	protected.GET("/users/:id/order_statistics", statisticsH.GetUserOrdersStatistics)

	// Reviews
	protected.POST("/reviews", reviewH.CreateReview)
	protected.GET("/reviews", reviewH.ListReviews)
	protected.GET("/reviews/:id", reviewH.GetReview)
	protected.PATCH("/reviews/:id", reviewH.UpdateReview)

	for _, route := range r.Routes() {
		log.Println(route.Method, route.Path)
	}

	return r
}
