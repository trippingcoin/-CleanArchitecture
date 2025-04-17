package usecase

import (
	"context"
	"review_service/internal/domain"
	"time"
)

type ReviewUsecase struct {
	ReviewRepo domain.ReviewRepository
}

func NewReviewUsecase(ReviewRepo domain.ReviewRepository) *ReviewUsecase {
	return &ReviewUsecase{
		ReviewRepo: ReviewRepo,
	}
}

// Create a review
func (u *ReviewUsecase) CreateReview(ctx context.Context, review *domain.Review) (*domain.Review, error) {
	// Call repository to save the review and return the saved review with ID
	err := u.ReviewRepo.Create(review)
	if err != nil {
		return nil, err
	}
	return review, nil
}

// Update a review
func (u *ReviewUsecase) UpdateReview(ctx context.Context, id string, rating float32, comment string) (*domain.Review, error) {
	review, err := u.ReviewRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Update the review fields
	review.Rating = rating
	review.Comment = comment
	review.UpdatedAt = time.Now()

	// Save the updated review
	err = u.ReviewRepo.Update(review)
	if err != nil {
		return nil, err
	}

	return review, nil
}
