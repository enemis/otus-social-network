package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type RepositoryManager struct {
	UserRepository
}

func NewRepositoryManager(db *sqlx.DB) *RepositoryManager {
	return &RepositoryManager{
		UserRepository: NewUserRepository(db),
	}
}
