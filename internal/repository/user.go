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
	FindUsers(name, surname string) ([]*model.User, *app_error.HttpError)
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
	var user model.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE id=$1 LIMIT 1", userId)

	if err != nil {
		return new(model.User), app_error.NewHttpError(err, "user not found", "user_id", http.StatusBadRequest)
	}

	return &user, nil
}

func (r *UserRepositoryInstance) FindUsers(name, surname string) ([]*model.User, *app_error.HttpError) {
	var users []*model.User
	query := "SELECT * FROM users WHERE "
	paramName := name + "%"
	paramSurname := surname + "%"
	limitPart := " LIMIT 30;"

	var err error

	if len(name) > 1 && len(surname) > 1 {
		err = r.db.Select(&users, query+"(name LIKE $1 and surname LIKE $2) OR (surname LIKE $3 and name LIKE $4)"+limitPart, paramName, paramSurname, paramName, paramSurname)
	} else if len(name) > 0 {
		err = r.db.Select(&users, query+"name LIKE $1"+limitPart, paramName)
	} else if len(surname) > 0 {
		err = r.db.Select(&users, query+"surname LIKE $1"+limitPart, paramName)
	} else {
		err = r.db.Select(&users, query+limitPart)
	}

	if err != nil {
		return users, app_error.NewHttpError(err, "user not found", "users", http.StatusBadRequest)
	}

	return users, nil
}
