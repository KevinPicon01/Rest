package database

import (
	"context"
	"database/sql"
	"kevinPicon/go/rest-ws/models"
	"log"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db: db}, nil
}

func (repo *PostgresRepository) InsertUser(ctx context.Context, user *models.User) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", user.Name, user.Email, user.Password)
	return err
}

func (repo *PostgresRepository) GetUserById(ctx context.Context, id int64) (*models.User, error) {
	var user models.User
	err := repo.db.QueryRowContext(ctx, "SELECT name, email FROM users WHERE id = $1", id).Scan(&user.Name, &user.Email)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = repo.db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	return &user, nil
}

func (repo *PostgresRepository) Close() error {
	return repo.db.Close()
}
