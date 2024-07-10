package repository

import (
	"github.com/munaiplan/munaiplan-backend/internal/infrastructure/database/postgres/repository"
	"gorm.io/gorm"
)

type Repository struct {
    Users UsersRepository
}

func NewRepositories(db *gorm.DB) *Repository {
    return &Repository{
        Users: repository.NewUsersRepository(db),
    }
}