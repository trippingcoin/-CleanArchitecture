package domain

import "time"

type Review struct {
	ID        string    `json:"id"`
	ProductID string    `json:"product_id"`
	UserID    string    `json:"user_id"`
	Rating    float32   `json:"rating"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ReviewRepository interface {
	// Create a new review in the repository.
	Create(review *Review) error
	// Update updates an existing review.
	Update(review *Review) error
	GetByID(reviewID string) (*Review, error)
}
