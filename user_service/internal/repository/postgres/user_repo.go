package postgres

import (
	"context"
	"database/sql"
	"errors"
	internal "user_service/internal/domain"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *internal.User) error {
	query := `INSERT INTO users (name, email, password, phone, role) VALUES ($1, $2, $3, $4, $5) RETURNING user_id`
	err := r.DB.QueryRowContext(ctx, query, user.Name, user.Email, user.Password, user.Phone, user.Role).Scan(&user.ID)
	return err
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*internal.User, error) {
	query := `SELECT user_id, name, email, password, phone, role FROM users WHERE email=$1`
	row := r.DB.QueryRowContext(ctx, query, email)

	var user internal.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Phone, &user.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // User not found
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int64) (*internal.User, error) {
	query := `SELECT user_id, name, email, password, phone, role FROM users WHERE user_id=$1`
	row := r.DB.QueryRowContext(ctx, query, id)

	var user internal.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Phone, &user.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
