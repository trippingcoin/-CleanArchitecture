package domain

type UserStatistics struct {
	UserID        string
	TotalOrders   int
	PeakOrderHour int
	TotalSpent    float64
}
