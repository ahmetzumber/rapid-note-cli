package repository

import (
	"github.com/ahmetzumber/rapid-note-cli/internal/modal"
	"gorm.io/gorm"
)

type IRepository interface {
	GetUsers() ([]modal.User, error)
	AddUser(createUser modal.User) (modal.User, error)
	AddNote(createNote modal.Note) (modal.Note, error)
}

type Repository struct {
	db 				*gorm.DB
}

func NewRepository(db *gorm.DB) IRepository {
	return &Repository{db: db}
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(modal.User{}, modal.Note{})
}

func (r *Repository) GetUsers() ([]modal.User, error) {
	var users []modal.User
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

func (r *Repository) AddUser(userRequest modal.User) (modal.User, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&userRequest).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return modal.User{}, nil
	}
	return userRequest, nil
}

func (r *Repository) AddNote(noteRequest modal.Note) (modal.Note, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&noteRequest).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return modal.Note{}, nil
	}
	return noteRequest, nil
}

