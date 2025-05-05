package usecase

import (
	"fmt"
	"statistics_service/internal/domain"
)

// StatisticsRepo defines the methods required for interacting with the data layer.
type StatisticsRepo interface {
	UpdateUserStatistics(userID string, amount float64, hour int) error
	GetUserStatistics(userID string) (*domain.UserStatistics, error)
}

// StatisticsUsecase defines the business logic related to user statistics.
type StatisticsUsecase struct {
	repo StatisticsRepo
}

// NewStatisticsUsecase creates and returns a new instance of StatisticsUsecase.
func NewStatisticsUsecase(repo StatisticsRepo) *StatisticsUsecase {
	return &StatisticsUsecase{repo: repo}
}

// UpdateUserStatistics updates statistics for a given user.
// It stores the amount spent and the most active hour for the user.
func (u *StatisticsUsecase) UpdateUserStatistics(userID string, amount float64, hour int) error {
	// Call the repository method to update the user statistics.
	err := u.repo.UpdateUserStatistics(userID, amount, hour)
	if err != nil {
		// Adding context to the error for easier debugging.
		return fmt.Errorf("failed to update statistics for user %s: %v", userID, err)
	}
	return nil
}

// GetUserStats retrieves the statistics for a given user.
func (u *StatisticsUsecase) GetUserStats(userID string) (*domain.UserStatistics, error) {
	// Call the repository method to get the user statistics.
	stats, err := u.repo.GetUserStatistics(userID)
	if err != nil {
		// Adding context to the error for easier debugging.
		return nil, fmt.Errorf("failed to get statistics for user %s: %v", userID, err)
	}
	return stats, nil
}
