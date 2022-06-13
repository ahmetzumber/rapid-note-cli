package repository

import (
	"github.com/ahmetzumber/rapid-note-cli/internal/config"
	"github.com/ahmetzumber/rapid-note-cli/internal/user"
	"gorm.io/gorm"
)

type IRepository interface {
	GetUsers() ([]user.User, error)
	AddUser(createUser config.CreateUserRequest) (user.User, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) IRepository {
	return &Repository{db: db}
}

func (r *Repository) GetUsers() ([]user.User, error) {
	var users []user.User
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Find(&users).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *Repository) AddUser(user config.CreateUserRequest) (user.User, error) {
	panic("implement me")
}


