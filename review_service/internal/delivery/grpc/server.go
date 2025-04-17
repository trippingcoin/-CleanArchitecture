package grpc

import (
	"context"
	"review_service/internal/domain"
	"review_service/internal/usecase"
	"review_service/proto/reviewpb"
	"time"
)

type ReviewServiceServer struct {
	reviewpb.ReviewServiceServer
	UC *usecase.ReviewUsecase
}

func NewReviewServiceServer(uc *usecase.ReviewUsecase) *ReviewServiceServer {
	return &ReviewServiceServer{UC: uc}
}

// CreateReview handles the creation of a new review.
func (s *ReviewServiceServer) CreateReview(ctx context.Context, req *reviewpb.CreateReviewRequest) (*reviewpb.ReviewResponse, error) {
	// Create the Review struct and populate it with request data
	review := &domain.Review{
		ProductID: req.GetProductId(),
		UserID:    req.GetUserId(),
		Rating:    req.GetRating(),
		Comment:   req.GetComment(),
		CreatedAt: time.Now(), // Set created_at to the current time
		UpdatedAt: time.Now(), // Set updated_at to the current time
	}

	// Call the usecase to create the review
	createdReview, err := s.UC.CreateReview(ctx, review)
	if err != nil {
		return nil, err
	}

	// Return the created review response
	return &reviewpb.ReviewResponse{
		Id:        createdReview.ID,
		ProductId: createdReview.ProductID,
		UserId:    createdReview.UserID,
		Rating:    createdReview.Rating,
		Comment:   createdReview.Comment,
	}, nil
}

// UpdateReview handles the update of an existing review's details.
func (s *ReviewServiceServer) UpdateReview(ctx context.Context, req *reviewpb.UpdateReviewRequest) (*reviewpb.ReviewResponse, error) {
	// Call the usecase to update the review
	updatedReview, err := s.UC.UpdateReview(ctx, req.GetId(), req.GetRating(), req.GetComment())
	if err != nil {
		return nil, err
	}

	// Return the updated review response
	return &reviewpb.ReviewResponse{
		Id:        updatedReview.ID,
		ProductId: updatedReview.ProductID,
		UserId:    updatedReview.UserID,
		Rating:    updatedReview.Rating,
		Comment:   updatedReview.Comment,
	}, nil
}
