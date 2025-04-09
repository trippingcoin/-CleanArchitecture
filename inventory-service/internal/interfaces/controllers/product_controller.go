package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	db *sql.DB
}

func NewProductController(db *sql.DB) *ProductController {
	return &ProductController{db: db}
}

func (p *ProductController) List(c *gin.Context) {
	rows, err := p.db.Query("SELECT id, name, price FROM products")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot fetch products"})
		return
	}
	defer rows.Close()

	var products []map[string]interface{}
	for rows.Next() {
		var id int
		var name string
		var price float64
		rows.Scan(&id, &name, &price)
		products = append(products, gin.H{"id": id, "name": name, "price": price})
	}
	c.JSON(http.StatusOK, products)
}

func (p *ProductController) Create(c *gin.Context) {
	var body struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	_, err := p.db.Exec("INSERT INTO products (name, price) VALUES ($1, $2)", body.Name, body.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert product"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Product created"})
}
