package service

import (
	"fmt"
	"otus-social-network/internal/config"
	"otus-social-network/internal/repository"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Container struct {
	db                *sqlx.DB
	repositoryManager *repository.RepositoryManager
	authService       AuthService
	userService       UserService
}

func NewContainer(config *config.Config) (*Container, error) {
	connectString := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=%s", config.DBHost, config.DBUsername, config.DBName, config.DBPassword, config.DBSSLMode)
	logrus.Debug("db connect string: ")
	logrus.Debugln(connectString)
	db, err := sqlx.Connect("postgres", connectString)

	if err != nil {
		return nil, err
	}

	var container Container

	repositoryManager := repository.NewRepositoryManager(db)

	container.db = db
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
