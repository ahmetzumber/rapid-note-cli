package repository

import (
	"github.com/ahmetzumber/rapid-note-cli/internal/config"
	"github.com/ahmetzumber/rapid-note-cli/internal/modal"
	"gorm.io/gorm"
)

type IRepository interface {
	GetUsers() ([]modal.User, error)
	AddUser(createUser config.CreateUserRequest) (modal.User, error)
	AddNote(createNote config.CreateNoteRequest) (modal.Note, error)
}

type Repository struct {
	db 				*gorm.DB
}

func NewRepository(db *gorm.DB) IRepository {
	return &Repository{db: db}
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

func (r *Repository) AddUser(userRequest config.CreateUserRequest) (modal.User, error) {
	newUser := modal.User{
		Username: userRequest.Username,
		Email:    userRequest.Email,
	}
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(newUser).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return modal.User{}, nil
	}
	return newUser, nil
}

func (r *Repository) AddNote(noteRequest config.CreateNoteRequest) (modal.Note, error) {
	newNote := modal.Note{
		Data: noteRequest.Data,
	}
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(newNote).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return modal.Note{}, nil
	}
	return newNote, nil
}

