package service

import (
	"otus-social-network/internal/app_error"
	"otus-social-network/internal/dto"
	"otus-social-network/internal/model"
	"otus-social-network/internal/repository"
)

type PostService interface {
	AddPost(user *model.User, post dto.PostInput) *app_error.HttpError
	RemovePost(user *model.User, friend *model.User) *app_error.HttpError
}

type PostServiceInstance struct {
	repositoryManager *repository.RepositoryManager
}

func (s *PostServiceInstance) AddPost(user *model.User, post dto.PostInput) *app_error.HttpError {

	return s.repositoryManager.PostRepository.AddPost(user, post)
}

func NewPostService(repositoryManager *repository.RepositoryManager) *PostServiceInstance {
	return &PostServiceInstance{
		repositoryManager: repositoryManager,
	}
}

func (s *PostServiceInstance) RemovePost(user *model.User, post dto.PostInput) *app_error.HttpError {
	return nil

	// return s.repositoryManager.CreatePost(user, post)
}
