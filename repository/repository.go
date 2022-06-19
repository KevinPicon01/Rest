package repository

import (
	"context"
	"kevinPicon/go/rest-ws/models"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id string) (*models.UserPayload, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	InsertPost(ctx context.Context, post *models.Post) error
	GetPostById(ctx context.Context, id string) (*models.Post, error)
	Close() error
}

var impl UserRepository

func SetRepository(repo UserRepository) {
	impl = repo
}

func InsertUser(ctx context.Context, user *models.User) error {
	return impl.InsertUser(ctx, user)
}

func GetUserById(ctx context.Context, id string) (*models.UserPayload, error) {
	return impl.GetUserById(ctx, id)
}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return impl.GetUserByEmail(ctx, email)
}

func InsertPost(ctx context.Context, post *models.Post) error {
	return impl.InsertPost(ctx, post)
}

func GetPostById(ctx context.Context, id string) (*models.Post, error) {
	return impl.GetPostById(ctx, id)
}

func Close() error {
	return impl.Close()
}
