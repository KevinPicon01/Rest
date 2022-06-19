package database

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
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
	_, err := repo.db.ExecContext(ctx, "INSERT INTO users ( id, name, email, password) VALUES ($1, $2, $3, $4)",
		user.Id, user.Name, user.Email, user.Password)
	return err
}

func (repo *PostgresRepository) InsertPost(ctx context.Context, post *models.Post) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO posts ( id, title, post_content, user_id) VALUES ($1, $2, $3, $4)",
		post.Id, post.Title, post.Content, post.AuthorId)
	return err
}

func (repo *PostgresRepository) GetUserById(ctx context.Context, id string) (*models.UserPayload, error) {
	var user models.UserPayload
	err := repo.db.QueryRowContext(ctx, "SELECT id, name, email FROM users WHERE id = $1", id).Scan(&user.Id, &user.Name, &user.Email)
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
func (repo *PostgresRepository) GetPostById(ctx context.Context, postId string) (*models.Post, error) {
	var post models.Post
	err := repo.db.QueryRowContext(ctx, "SELECT id, title, post_content FROM posts WHERE id = $1",
		postId).Scan(&post.Id, &post.Title, &post.Content)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = repo.db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	return &post, nil
}

func (repo *PostgresRepository) UpdatePost(ctx context.Context, post *models.Post) error {
	repo.db.Ping()
	_, err := repo.db.ExecContext(ctx, "UPDATE posts SET title = $1, post_content = $2 WHERE id = $3 and user_id = $4",
		post.Title, post.Content, post.Id, post.AuthorId)
	return err
}

func (repo *PostgresRepository) DeletePost(ctx context.Context, post *models.Post) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM posts WHERE id = $1 and user_id = $2", post.Id, post.AuthorId)
	return err
}

func (repo *PostgresRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	rows, err := repo.db.QueryContext(ctx, "SELECT id, password, name, email FROM users WHERE email = $1",
		email)

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Password, &user.Name, &user.Email); err != nil {
			return nil, err
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *PostgresRepository) Close() error {
	return repo.db.Close()
}
