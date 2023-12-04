package app

import (
	"otus-social-network/internal/config"
	"otus-social-network/internal/server"
	"otus-social-network/internal/service"
)

type SocialNetworkApp struct {
	config    *config.Config
	container *service.Container
}

func NewApp() (*SocialNetworkApp, error) {
	config, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	container, err := service.NewContainer(config)

	if err != nil {
		return nil, err
	}

	app := SocialNetworkApp{container: container, config: config}
	return &app, nil
}

func (app *SocialNetworkApp) Run() error {
	server := server.NewServer(app.config, app.container.AuthService(), app.container.UserService(), app.container.FriendService())
	err := server.Run()

	return err
}
