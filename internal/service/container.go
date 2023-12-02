package service

import (
	"otus-social-network/internal/config"
	"otus-social-network/internal/db"
	"otus-social-network/internal/repository"
)

type Container struct {
	db                *db.DatabaseStack
	repositoryManager *repository.RepositoryManager
	authService       AuthService
	userService       UserService
}

func NewContainer(config *config.Config) (*Container, error) {
	dbStack := db.NewDatabaseStack(config)

	var container Container

	repositoryManager := repository.NewRepositoryManager(dbStack)

	container.db = dbStack
	container.repositoryManager = repositoryManager

	container.authService = NewAuthService(repositoryManager, config)
	container.userService = NewUserService(repositoryManager)

	return &container, nil
}

func (c *Container) AuthService() AuthService {
	return c.authService
}

func (c *Container) UserService() UserService {
	return c.userService
}
