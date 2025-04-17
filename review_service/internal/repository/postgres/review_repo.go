package postgres

import (
	"database/sql"
	"fmt"
	"review_service/internal/domain"
	"time"
)

type reviewRepo struct {
	db *sql.DB
}

// NewReviewRepository creates a new instance of reviewRepo
func NewReviewRepository(db *sql.DB) domain.ReviewRepository {
	return &reviewRepo{db: db}
}

// Create creates a new review in the repository
func (r *reviewRepo) Create(review *domain.Review) error {
	// Start a transaction
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Insert review data
	query := `INSERT INTO reviews (product_id, user_id, rating, comment, created_at) 
	          VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err = tx.QueryRow(query, review.ProductID, review.UserID, review.Rating, review.Comment, time.Now()).Scan(&review.ID)
	if err != nil {
		return err
	}

	// Commit the transaction
	return tx.Commit()
}

// Update updates an existing review
func (r *reviewRepo) Update(review *domain.Review) error {
	// Update review status
	_, err := r.db.Exec(`UPDATE reviews SET rating = $1, comment = $2, updated_at = $3 WHERE id = $4`,
		review.Rating, review.Comment, time.Now(), review.ID)
	return err
}

func (r *reviewRepo) GetByID(reviewID string) (*domain.Review, error) {
	query := `SELECT id, product_id, user_id, rating, comment, created_at, updated_at 
	          FROM reviews WHERE id = $1`

	var review domain.Review
	err := r.db.QueryRow(query, reviewID).Scan(
		&review.ID,
		&review.ProductID,
		&review.UserID,
		&review.Rating,
		&review.Comment,
		&review.CreatedAt,
		&review.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("review not found")
		}
		return nil, err
	}

	return &review, nil
}
