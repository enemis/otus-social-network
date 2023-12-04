package seeder

import (
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
)

func (s *Seeder) PostSeed(user uuid.UUID, postCount uint) {
	faker.Paragraph()
}
