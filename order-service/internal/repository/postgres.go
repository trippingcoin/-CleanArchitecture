package repository

import (
	"database/sql"
	"order-service/internal/model"

	"github.com/google/uuid"
)

type OrderRepository interface {
	Create(o model.Order) (model.Order, error)
	GetByID(id string) (model.Order, error)
	UpdateStatus(id, status string) (model.Order, error)
	List(userID string) ([]model.Order, error)
}

type postgresRepo struct {
	db *sql.DB
}

func NewPostgresRepo(db *sql.DB) OrderRepository {
	return &postgresRepo{db: db}
}

func (r *postgresRepo) Create(o model.Order) (model.Order, error) {
	o.ID = uuid.NewString()
	tx, err := r.db.Begin()
	if err != nil {
		return o, err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`INSERT INTO orders (id, user_id, status) VALUES ($1, $2, $3)`, o.ID, o.UserID, "pending")
	if err != nil {
		return o, err
	}

	for _, item := range o.Items {
		_, err = tx.Exec(`INSERT INTO order_items (id, product_id, quantity) VALUES ($1, $2, $3)`, o.ID, item.ProductID, item.Quantity)
		if err != nil {
			return o, err
		}
	}
	err = tx.Commit()
	return o, err
}

func (r *postgresRepo) GetByID(id string) (model.Order, error) {
	var o model.Order
	err := r.db.QueryRow(`SELECT id, user_id, status FROM orders WHERE id=$1`, id).
		Scan(&o.ID, &o.UserID, &o.Status)
	if err != nil {
		return o, err
	}

	rows, err := r.db.Query(`SELECT product_id, quantity FROM order_items WHERE id=$1`, id)
	if err != nil {
		return o, err
	}
	defer rows.Close()

	for rows.Next() {
		var item model.OrderItem
		rows.Scan(&item.ProductID, &item.Quantity)
		o.Items = append(o.Items, item)
	}
	return o, nil
}

func (r *postgresRepo) UpdateStatus(id, status string) (model.Order, error) {
	_, err := r.db.Exec(`UPDATE orders SET status=$1 WHERE id=$2`, status, id)
	if err != nil {
		return model.Order{}, err
	}
	return r.GetByID(id)
}

func (r *postgresRepo) List(userID string) ([]model.Order, error) {
	rows, err := r.db.Query(`SELECT id FROM orders WHERE user_id=$1`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []model.Order
	for rows.Next() {
		var id string
		rows.Scan(&id)
		o, err := r.GetByID(id)
		if err == nil {
			orders = append(orders, o)
		}
	}
	return orders, nil
}
