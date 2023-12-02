package repository

import (
	"otus-social-network/internal/db"

	_ "github.com/lib/pq"
)

type RepositoryManager struct {
	UserRepository
}

func NewRepositoryManager(db *db.DatabaseStack) *RepositoryManager {
	return &RepositoryManager{
		UserRepository: NewUserRepository(db),
	}
}
