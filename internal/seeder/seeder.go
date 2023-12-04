package seeder

import "otus-social-network/internal/service"

type Seeder struct {
	authService service.AuthService
}

func NewSeeder(authService service.AuthService) *Seeder {
	return &Seeder{authService: authService}
}

func (s *Seeder) Seed() {
	s.UserSeed(10000)
}
