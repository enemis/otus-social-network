package repository

import (
	"otus-social-network/internal/app_error"
	"otus-social-network/internal/db"
	"otus-social-network/internal/dto"
	"otus-social-network/internal/model"
	"time"

	"github.com/google/uuid"
)

type PostRepository interface {
	AddPost(user *model.User, post dto.PostInput) *app_error.HttpError
	RemovePost(postId uuid.UUID) *app_error.HttpError
}

type PostRepositoryInstance struct {
	db *db.DatabaseStack
}

func NewPostRepository(db *db.DatabaseStack) *PostRepositoryInstance {
	return &PostRepositoryInstance{db: db}
}

func (r *PostRepositoryInstance) AddPost(user *model.User, post dto.PostInput) *app_error.HttpError {
	query := "INSERT INTO posts (title, post, status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"

	now := time.Now()

	var postId uuid.UUID

	err := r.db.Master().QueryRow(query, post.Title, post.Post, post.Status, now, now).Scan(&postId)

	if err != nil {
		return app_error.NewInternalServerError(err)
	}

	return nil
}

func (r *PostRepositoryInstance) RemovePost(postId uuid.UUID) *app_error.HttpError {
	return nil
}
