package repository

import (
	"fmt"
	"net/http"
	"otus-social-network/internal/app_error"
	"otus-social-network/internal/db"
	"otus-social-network/internal/model"
	"time"
)

type FriendRepository interface {
	AddFriend(user *model.User, friend *model.User) *app_error.HttpError
	RemoveFriend(user *model.User, friend *model.User) *app_error.HttpError
}

type FriendRepositoryInstance struct {
	db *db.DatabaseStack
}

func NewFriendRepository(db *db.DatabaseStack) *UserRepositoryInstance {
	return &UserRepositoryInstance{db: db}
}

func (r *UserRepositoryInstance) AddFriend(user *model.User, friend *model.User) *app_error.HttpError {
	rows, err := r.db.Slave().Query("SELECT EXISTS(SELECT 1 FROM friends WHERE user_id=$1 AND friend_id=$2)", user.Id, friend.Id)
	if err != nil {
		return app_error.NewInternalServerError(err)
	}

	var exists bool

	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&exists); err != nil {
			return app_error.NewInternalServerError(err)
		}

		if exists {
			return app_error.NewHttpError(err, fmt.Sprintf("User %s %s already added as friend for user %s %s", friend.Surname, friend.Name, user.Surname, friend.Name), "friend_id", http.StatusBadRequest)
		}
	}

	query := "INSERT INTO friends (user_id, friend_id, created_at) VALUES ($1, $2, $3)"

	now := time.Now()

	_, err = r.db.Master().Exec(query, user.Id, friend.Id, now)

	if err != nil {
		return app_error.NewInternalServerError(err)
	}

	return nil
}

func (r *UserRepositoryInstance) RemoveFriend(user *model.User, friend *model.User) *app_error.HttpError {
	rows, err := r.db.Slave().Query("SELECT EXISTS(SELECT 1 FROM friends WHERE user_id=$1 AND friend_id=$2)", user.Id, friend.Id)
	if err != nil {
		return app_error.NewInternalServerError(err)
	}

	var exists bool

	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&exists); err != nil {
			return app_error.NewInternalServerError(err)
		}

		if !exists {
			return nil
		}
	}

	query := "DELETE FROM friends WHERE user_id=$1 AND friend_id=$2"

	_, err = r.db.Master().Exec(query, user.Id, friend.Id)

	if err != nil {
		return app_error.NewInternalServerError(err)
	}

	return nil
}
