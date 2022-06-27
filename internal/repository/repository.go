package repository

import (
	"github.com/ahmetzumber/rapid-note-cli/internal/modal"
	"gorm.io/gorm"
	"log"
)

type IRepository interface {
	GetUsers() ([]modal.User, error)
	AddUser(createUser modal.User) (modal.User, error)
	AddNote(createNote modal.Note) (modal.Note, error)
	GetUserIDByUsername(username string) int
	GetCurrentUserNotesByUserID(userID int) []modal.Note
}

type Repository struct {
	db 				*gorm.DB
}

func NewRepository(db *gorm.DB) IRepository {
	return &Repository{db: db}
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(modal.Note{})
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

func (r *Repository) GetUserIDByUsername(username string) int {
	var user modal.User
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where(&modal.User{ Username: username }).First(&user).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}
	return user.ID
}

func (r *Repository) GetCurrentUserNotesByUserID(userID int) []modal.Note {
	var notes []modal.Note
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where(&modal.Note{ UserID: userID }).Find(&notes).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}
	return notes
}