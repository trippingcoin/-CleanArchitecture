package domain

import "time"

type UserStatistics struct {
	UserID           string    `json:"user_id"`
	TotalOrders      int       `json:"total_orders"`
	TotalSpent       float64   `json:"total_spent"`
	PeakOrderHour    int       `json:"peak_order_hour"` // Track the peak order hour
	RegistrationDate time.Time `json:"registration_date"`
	AvgOrderValue    float64   `json:"avg_order_value"`
	MostActiveHour   int       `json:"most_active_hour"` // Add the field for most active hour
}
