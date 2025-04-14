package repository

import (
	"database/sql"
	"errors"
	"inventory-service/internal/models"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type ProductRepository interface {
	Create(product models.Product) (models.Product, error)
	GetByID(id string) (models.Product, error)
	Update(product models.Product) (models.Product, error)
	Delete(id string) error
	List() ([]models.Product, error)
}

type postgresRepo struct {
	db *sql.DB
}

func NewPostgresRepo(db *sql.DB) ProductRepository {
	return &postgresRepo{db: db}
}

func (r *postgresRepo) Create(p models.Product) (models.Product, error) {
	p.ID = uuid.NewString()
	_, err := r.db.Exec(`
		INSERT INTO products (id, name, category, price, stock)
		VALUES ($1, $2, $3, $4, $5)
	`, p.ID, p.Name, p.Category, p.Price, p.Stock)
	return p, err
}

func (r *postgresRepo) GetByID(id string) (models.Product, error) {
	row := r.db.QueryRow(`SELECT id, name, category, price, stock FROM products WHERE id=$1`, id)
	var p models.Product
	err := row.Scan(&p.ID, &p.Name, &p.Category, &p.Price, &p.Stock)
	if err != nil {
		return p, errors.New("product not found")
	}
	return p, nil
}

func (r *postgresRepo) Update(p models.Product) (models.Product, error) {
	_, err := r.db.Exec(`
		UPDATE products SET name=$1, category=$2, price=$3, stock=$4 WHERE id=$5
	`, p.Name, p.Category, p.Price, p.Stock, p.ID)
	return p, err
}

func (r *postgresRepo) Delete(id string) error {
	_, err := r.db.Exec(`DELETE FROM products WHERE id=$1`, id)
	return err
}

func (r *postgresRepo) List() ([]models.Product, error) {
	rows, err := r.db.Query(`SELECT id, name, category, price, stock FROM products`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Category, &p.Price, &p.Stock); err == nil {
			products = append(products, p)
		}
	}
	return products, nil
}
