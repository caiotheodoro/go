package postgres

import (
	"context"

	"github.com/caiotheodoro/go/auth-grpc-microservice/internal/app/user"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) FindAccountByID(ctx context.Context, id string) (*user.Account, error) {
	return nil, nil
}

func (r *UserRepository) AddAccount(ctx context.Context, a *user.Account) error {
	return nil
}
