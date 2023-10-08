package command

import (
	"otus-social-network/internal/config"
	"otus-social-network/internal/seeder"
	"otus-social-network/internal/service"
)

type SeedApp struct {
	config    *config.Config
	container *service.Container
	seeder    *seeder.Seeder
}

func NewSeeder() (*SeedApp, error) {
	config, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	container, err := service.NewContainer(config)

	if err != nil {
		return nil, err
	}

	seeder := seeder.NewSeeder(container.AuthService())
	app := SeedApp{container: container, config: config, seeder: seeder}

	return &app, nil
}

func (app *SeedApp) RunImport() {
	app.seeder.Seed()

}
