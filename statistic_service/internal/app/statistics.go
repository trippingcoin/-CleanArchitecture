package app

import "statistics_service/internal/port"

type StatisticsApp struct {
	repo port.StatisticsRepository
}

func New(repo port.StatisticsRepository) *StatisticsApp {
	return &StatisticsApp{repo: repo}
}

func (a *StatisticsApp) ProcessOrder(userID string, amount float64, hour int) error {
	return a.repo.UpdateUserStatistics(userID, amount, hour)
}

func (a *StatisticsApp) GetStatistics(userID string) (interface{}, error) {
	return a.repo.GetUserStatistics(userID)
}
