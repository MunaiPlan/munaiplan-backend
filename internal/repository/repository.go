package repository

import (
	"github.com/munaiplan/munaiplan-backend/infrastructure/database/postgres/repository"
	"gorm.io/gorm"
)

type Repository struct {
    Users UserRepository
}

func NewRepositories(db *gorm.DB) *Repository {
    return &Repository{
        Users: repository.NewUserRepository(db),
    }
}