package grpc

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
	"user_service/internal/domain"
	"user_service/internal/usecase"
	"user_service/proto/userpb"
)

type UserServiceServer struct {
	userpb.UnimplementedUserServiceServer
	Usecase *usecase.UserUsecase
}

const jwtSecret = "superSecret"

func NewUserServiceServer(uc *usecase.UserUsecase) *UserServiceServer {
	return &UserServiceServer{Usecase: uc}
}

func (s *UserServiceServer) RegisterUser(ctx context.Context, req *userpb.UserRequest) (*userpb.UserResponse, error) {
	user := &domain.User{
		Name:     req.GetName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Phone:    req.GetPhone(),
		Role:     req.GetRole(),
	}
	err := s.Usecase.Register(ctx, user)
	if err != nil {
		return nil, status.Errorf(codes.AlreadyExists, "registration failed: %v", err)
	}

	return &userpb.UserResponse{UserId: int32(user.ID)}, nil
}

func (s *UserServiceServer) AuthenticateUser(ctx context.Context, req *userpb.AuthRequest) (*userpb.AuthResponse, error) {
	user, err := s.Usecase.Authenticate(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %v", err)
	}

	// Build JWT claims
	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign it
	signed, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not sign token: %v", err)
	}

	return &userpb.AuthResponse{Token: signed}, nil
}

func (s *UserServiceServer) GetUserProfile(ctx context.Context, req *userpb.UserID) (*userpb.UserProfile, error) {
	user, err := s.Usecase.GetProfile(ctx, int64(req.GetUserId()))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found: %v", err)
	}

	return &userpb.UserProfile{
		UserId: int32(user.ID),
		Name:   user.Name,
		Email:  user.Email,
		Phone:  user.Phone,
		Role:   user.Role,
	}, nil
}
