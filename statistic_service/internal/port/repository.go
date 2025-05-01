package port

import "statistics_service/internal/domain"

type StatisticsRepository interface {
	UpdateUserStatistics(userID string, amount float64, hour int) error
	GetUserStatistics(userID string) (*domain.UserStatistics, error)
}
