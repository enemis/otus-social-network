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
	friendService     FriendService
	postService       PostService
}

func NewContainer(config *config.Config) (*Container, error) {
	dbStack := db.NewDatabaseStack(config)

	var container Container

	repositoryManager := repository.NewRepositoryManager(dbStack)

	container.db = dbStack
	container.repositoryManager = repositoryManager

	container.authService = NewAuthService(repositoryManager, config)
	container.userService = NewUserService(repositoryManager)
	container.friendService = NewFriendService(repositoryManager)
	container.postService = NewPostService(repositoryManager)

	return &container, nil
}

func (c *Container) AuthService() AuthService {
	return c.authService
}

func (c *Container) UserService() UserService {
	return c.userService
}

func (c *Container) FriendService() FriendService {
	return c.friendService
}

func (c *Container) PostService() PostService {
	return c.postService
}
