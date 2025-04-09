package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	db *sql.DB
}

func NewOrderController(db *sql.DB) *OrderController {
	return &OrderController{db: db}
}

func (oc *OrderController) CreateOrder(c *gin.Context) {
	var order struct {
		UserID int `json:"user_id"`
		Items  []struct {
			ProductID int     `json:"product_id"`
			Quantity  int     `json:"quantity"`
			Price     float64 `json:"price"`
		} `json:"items"`
	}
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	tx, _ := oc.db.Begin()
	defer tx.Rollback()

	var orderID int
	err := tx.QueryRow("INSERT INTO orders (user_id) VALUES ($1) RETURNING id", order.UserID).Scan(&orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "order creation failed"})
		return
	}

	for _, item := range order.Items {
		_, err := tx.Exec("INSERT INTO order_items (order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4)",
			orderID, item.ProductID, item.Quantity, item.Price)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to insert order items"})
			return
		}
	}

	tx.Commit()
	c.JSON(http.StatusCreated, gin.H{"order_id": orderID})
}

func (oc *OrderController) GetOrder(c *gin.Context) {
	id := c.Param("id")
	row := oc.db.QueryRow("SELECT id, user_id, status FROM orders WHERE id=$1", id)

	var orderID, userID int
	var status string
	err := row.Scan(&orderID, &userID, &status)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": orderID, "user_id": userID, "status": status})
}

func (oc *OrderController) UpdateStatus(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Status string `json:"status"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}
	_, err := oc.db.Exec("UPDATE orders SET status=$1 WHERE id=$2", body.Status, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "update failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "status updated"})
}

func (oc *OrderController) ListOrders(c *gin.Context) {
	rows, err := oc.db.Query("SELECT id, user_id, status FROM orders")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "query failed"})
		return
	}
	defer rows.Close()

	var result []map[string]interface{}
	for rows.Next() {
		var id, userID int
		var status string
		rows.Scan(&id, &userID, &status)
		result = append(result, gin.H{"id": id, "user_id": userID, "status": status})
	}

	c.JSON(http.StatusOK, result)
}
