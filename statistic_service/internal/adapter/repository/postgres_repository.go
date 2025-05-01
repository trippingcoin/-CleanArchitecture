package repository

import (
	"database/sql"
	"statistics_service/internal/domain"
	"statistics_service/internal/port"
)

type PostgresRepo struct {
	db *sql.DB
}

func NewPostgresRepo(db *sql.DB) port.StatisticsRepository {
	return &PostgresRepo{db: db}
}

func (r *PostgresRepo) UpdateUserStatistics(userID string, amount float64, hour int) error {
	_, err := r.db.Exec(`INSERT INTO user_statistics (user_id, total_orders, total_spent, peak_hour) 
                         VALUES ($1, 1, $2, $3) 
                         ON CONFLICT (user_id) DO UPDATE 
                         SET total_orders = user_statistics.total_orders + 1, 
                             total_spent = user_statistics.total_spent + $2`, userID, amount, hour)
	return err
}

func (r *PostgresRepo) GetUserStatistics(userID string) (*domain.UserStatistics, error) {
	row := r.db.QueryRow(`SELECT user_id, total_orders, peak_hour, total_spent FROM user_statistics WHERE user_id = $1`, userID)
	stat := &domain.UserStatistics{}
	err := row.Scan(&stat.UserID, &stat.TotalOrders, &stat.PeakOrderHour, &stat.TotalSpent)
	return stat, err
}
