package repository

import (
	"fmt"
	"net/http"
	"otus-social-network/internal/app_error"
	"otus-social-network/internal/dto"
	"otus-social-network/internal/model"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type UserRepository interface {
	CreateUser(*dto.SignUpInput) (uuid.UUID, *app_error.HttpError)
	GetUser(email, password string) (*model.User, *app_error.HttpError)
	GetUserById(userId uuid.UUID) (*model.User, *app_error.HttpError)
}

type UserRepositoryInstance struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepositoryInstance {
	return &UserRepositoryInstance{db: db}
}

func (r *UserRepositoryInstance) CreateUser(user *dto.SignUpInput) (uuid.UUID, *app_error.HttpError) {
	rows, err := r.db.Query("SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)", user.Email)
	if err != nil {
		return uuid.Nil, app_error.NewInternalServerError(err)
	}

	var exists bool

	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&exists); err != nil {
			return uuid.Nil, app_error.NewInternalServerError(err)
		}

		if exists {
			return uuid.Nil, app_error.NewHttpError(err, fmt.Sprintf("User with email %s already registered", user.Email), "email", http.StatusBadRequest)
		}
	}

	query := "INSERT INTO users (name, surname, email, birthday, biography, city, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id"

	now := time.Now()
	var userId uuid.UUID
	err = r.db.QueryRow(query, user.Name, user.Surname, user.Email, user.Birthday, user.Biography, user.City, user.Password, now, now).Scan(&userId)

	if err != nil {
		return uuid.Nil, app_error.NewInternalServerError(err)
	}

	return userId, nil
}

func (r *UserRepositoryInstance) GetUser(email, password string) (*model.User, *app_error.HttpError) {
	rows, err := r.db.Queryx("SELECT * FROM users WHERE email=$1 and password=$2", email, password)
	if err != nil {
		return new(model.User), app_error.NewInternalServerError(err)
	}
	defer rows.Close()

	var user model.User
	for rows.Next() {
		err = rows.StructScan(&user)
		if err != nil {
			return new(model.User), app_error.NewInternalServerError(err)
		}
	}

	return &user, nil
}

func (r *UserRepositoryInstance) GetUserById(userId uuid.UUID) (*model.User, *app_error.HttpError) {
	rows, err := r.db.Queryx("SELECT * FROM users WHERE id=$1", userId)
	if err != nil {
		return new(model.User), app_error.NewInternalServerError(err)
	}
	defer rows.Close()

	var user model.User
	for rows.Next() {
		err = rows.StructScan(&user)
		if err != nil {
			return new(model.User), app_error.NewInternalServerError(err)
		}
	}

	return &user, nil
}
