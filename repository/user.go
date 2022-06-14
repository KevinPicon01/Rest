package repository

import (
	"context"
	"kevinPicon/go/rest-ws/models"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id string) (*models.User, error)
	Close() error
}

var impl UserRepository

func SetRepository(repo UserRepository) {
	impl = repo
}

func InsertUser(ctx context.Context, user *models.User) error {
	return impl.InsertUser(ctx, user)
}

func GetUserById(ctx context.Context, id string) (*models.User, error) {
	return impl.GetUserById(ctx, id)
}

func Close() error {
	return impl.Close()
}
