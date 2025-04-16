package usecase

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"user_service/internal/domain"
	"user_service/internal/repository/postgres"
)

type UserUsecase struct {
	Repo postgres.UserRepository
}

func NewUserUsecase(repo postgres.UserRepository) *UserUsecase {
	return &UserUsecase{Repo: repo}
}

func (u *UserUsecase) Register(ctx context.Context, user *domain.User) error {
	existing, _ := u.Repo.GetUserByEmail(ctx, user.Email)
	if existing != nil {
		return errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return u.Repo.CreateUser(ctx, user)
}

func (u *UserUsecase) Authenticate(ctx context.Context, email, password string) (*domain.User, error) {
	user, err := u.Repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (u *UserUsecase) GetProfile(ctx context.Context, id int64) (*domain.User, error) {
	return u.Repo.GetUserByID(ctx, id)
}
