package service

import (
	"otus-social-network/internal/app_error"
	"otus-social-network/internal/dto"
	"otus-social-network/internal/model"
	"otus-social-network/internal/repository"

	"github.com/google/uuid"
)

type UserService interface {
	GetUserById(userId uuid.UUID) (*model.User, *app_error.HttpError)
	FindUsers(search dto.FindUser) ([]*model.User, *app_error.HttpError)
}

type UserServiceInstance struct {
	repositoryManager *repository.RepositoryManager
}

func NewUserService(repositoryManager *repository.RepositoryManager) *UserServiceInstance {
	return &UserServiceInstance{
		repositoryManager: repositoryManager,
	}
}

func (s *UserServiceInstance) GetUserById(userId uuid.UUID) (*model.User, *app_error.HttpError) {
	return s.repositoryManager.GetUserById(userId)
}

func (s *UserServiceInstance) FindUsers(search dto.FindUser) ([]*model.User, *app_error.HttpError) {
	return s.repositoryManager.FindUsers(search.Name, search.Surname)
}
