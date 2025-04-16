package postgres

import (
	"database/sql"
	"order_service/internal/domain"
	"time"
)

type orderRepo struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) domain.OrderRepository {
	return &orderRepo{db: db}
}

func (r *orderRepo) Create(order *domain.Order) error {
	// Start a transaction
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Insert order data (including total price)
	query := `INSERT INTO orders (user_id, total_price, status) 
	          VALUES ($1, $2, $3) RETURNING order_id`
	err = tx.QueryRow(query, order.UserID, order.TotalPrice, order.Status).Scan(&order.ID)
	if err != nil {
		return err
	}

	// Insert each order item
	for _, item := range order.Items {
		_, err := tx.Exec(`INSERT INTO order_items (order_id, product_id, quantity, price_per_item)
			VALUES ($1, $2, $3, $4)`, order.ID, item.ProductID, item.Quantity, item.PricePerItem)
		if err != nil {
			return err
		}
	}

	// Commit the transaction
	return tx.Commit()
}

func (r *orderRepo) GetByID(id string) (*domain.Order, error) {
	row := r.db.QueryRow(`SELECT order_id, user_id, total_price, status 
	                       FROM orders WHERE order_id = $1`, id)

	var order domain.Order
	err := row.Scan(&order.ID, &order.UserID, &order.TotalPrice, &order.Status)
	if err != nil {
		return nil, err
	}

	// Fetch order items
	itemsRows, err := r.db.Query(`SELECT product_id, quantity, price_per_item 
	                               FROM order_items WHERE order_id = $1`, id)
	if err != nil {
		return nil, err
	}
	defer itemsRows.Close()

	for itemsRows.Next() {
		var item domain.OrderItem
		err := itemsRows.Scan(&item.ProductID, &item.Quantity, &item.PricePerItem)
		if err != nil {
			return nil, err
		}
		order.Items = append(order.Items, item)
	}

	return &order, nil
}

func (r *orderRepo) UpdateStatus(id string, status string) error {
	// Update order status
	_, err := r.db.Exec(`UPDATE orders SET status = $1, updated_at = $2 WHERE order_id = $3`,
		status, time.Now(), id)
	return err
}

func (r *orderRepo) ListByUser(userID string) ([]domain.Order, error) {
	rows, err := r.db.Query(`SELECT order_id, total_price, status 
	                         FROM orders WHERE user_id = $1`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []domain.Order
	for rows.Next() {
		var order domain.Order
		err := rows.Scan(&order.ID, &order.TotalPrice, &order.Status)
		if err != nil {
			return nil, err
		}

		// Fetch the items for each order
		itemsRows, err := r.db.Query(`SELECT product_id, quantity, price_per_item 
		                               FROM order_items WHERE order_id = $1`, order.ID)
		if err != nil {
			return nil, err
		}
		defer itemsRows.Close()

		for itemsRows.Next() {
			var item domain.OrderItem
			err := itemsRows.Scan(&item.ProductID, &item.Quantity, &item.PricePerItem)
			if err != nil {
				return nil, err
			}
			order.Items = append(order.Items, item)
		}

		orders = append(orders, order)
	}

	return orders, nil
}
