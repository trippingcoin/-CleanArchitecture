package usecase

import (
	"encoding/json"
	"fmt"
	"log"
	"statistics_service/internal/domain"
	"time"

	"github.com/nats-io/nats.go"
)

// StatisticsRepo defines the methods required for interacting with the data layer.
type StatisticsRepo interface {
	UpdateUserStatistics(userID string, amount float64, hour int) error
	GetUserStatistics(userID string) (*domain.UserStatistics, error)
}

// StatisticsUsecase defines the business logic related to user statistics.
type StatisticsUsecase struct {
	nc   *nats.Conn
	repo StatisticsRepo
}

// NewStatisticsUsecase creates and returns a new instance of StatisticsUsecase.
func NewStatisticsUsecase(repo StatisticsRepo, nats *nats.Conn) *StatisticsUsecase {
	return &StatisticsUsecase{repo: repo, nc: nats}
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

func (u *StatisticsUsecase) PublishHourlyStats() {
	ticker := time.NewTicker(15 * time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				stats := map[string]interface{}{
					"type":   "hourly_update",
					"items":  []string{"item1", "item2"},
					"orders": []int{101, 102},
					"time":   time.Now().Format(time.RFC3339),
				}

				payload, err := json.Marshal(stats)
				if err != nil {
					log.Println("Failed to marshal statistics:", err)
					continue
				}

				err = u.nc.Publish("ap2.statistics.event.updated", payload)
				if err != nil {
					log.Println("Failed to publish to NATS:", err)
				} else {
					log.Println("Published hourly statistics to NATS.")
				}
			}
		}
	}()
}
