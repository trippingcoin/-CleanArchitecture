package repository

import "statistics_service/internal/domain"

type InMemoryStatisticsRepository struct {
	data map[string]*domain.UserStatistics
}

func NewInMemoryStatisticsRepository() *InMemoryStatisticsRepository {
	return &InMemoryStatisticsRepository{
		data: make(map[string]*domain.UserStatistics),
	}
}

func (r *InMemoryStatisticsRepository) UpdateUserStatistics(userID string, amount float64, hour int) error {
	stats, exists := r.data[userID]
	if !exists {
		stats = &domain.UserStatistics{UserID: userID}
		r.data[userID] = stats
	}
	stats.TotalOrders++
	stats.TotalSpent += amount
	stats.PeakOrderHour = hour
	return nil
}

func (r *InMemoryStatisticsRepository) GetUserStatistics(userID string) (*domain.UserStatistics, error) {
	stats, exists := r.data[userID]
	if !exists {
		return nil, nil
	}
	return stats, nil
}
