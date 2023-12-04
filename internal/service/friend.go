package service

import (
	"otus-social-network/internal/app_error"
	"otus-social-network/internal/model"
	"otus-social-network/internal/repository"
)

type FriendService interface {
	AddFriend(user *model.User, friend *model.User) *app_error.HttpError
	RemoveFriend(user *model.User, friend *model.User) *app_error.HttpError
}

type FriendServiceInstance struct {
	repositoryManager *repository.RepositoryManager
}

func NewFriendService(repositoryManager *repository.RepositoryManager) *FriendServiceInstance {
	return &FriendServiceInstance{
		repositoryManager: repositoryManager,
	}
}

func (s *FriendServiceInstance) AddFriend(user *model.User, friend *model.User) *app_error.HttpError {

	return s.repositoryManager.AddFriend(user, friend)
}

func (s *FriendServiceInstance) RemoveFriend(user *model.User, friend *model.User) *app_error.HttpError {

	return s.repositoryManager.RemoveFriend(user, friend)
}
