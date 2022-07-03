package repository

import (
	"github.com/ahmetzumber/rapid-note-cli/internal/modal"
	"github.com/ahmetzumber/rapid-note-cli/internal/postgre"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"log"
	"testing"
)

func TestNewRepositoryCreatedByGivenDB(t *testing.T) {
	db, err := postgre.NewPostgreDB(postgre.Config)
	repo := NewRepository(db)
	assert.Nil(t, err)
	assert.NotNil(t, repo)
}

func TestRepository_GetUsers(t *testing.T) {
	db, err := postgre.NewPostgreDB(postgre.Config)
	repo := NewRepository(db)
	users, err := repo.GetUsers()
	assert.Nil(t, err)
	assert.NotNil(t, users)
}

func TestRepository_AddUser(t *testing.T) {
	mockUser := modal.User{
		Model:    gorm.Model{},
		ID:       0,
		Username: "deneme_name",
		Email:    "deneme_email",
	}
	db, err := postgre.NewPostgreDB(postgre.Config)
	repo := NewRepository(db)
	user, err := repo.AddUser(mockUser)
	assert.NotNil(t, user)
	assert.Nil(t, err)
}

func TestRepository_AddNote(t *testing.T) {
	mockNote := modal.Note{
		Model:  gorm.Model{},
		ID:     0,
		UserID: 0,
		Data:   "deneme_note",
	}
	db, err := postgre.NewPostgreDB(postgre.Config)
	repo := NewRepository(db)
	note, err := repo.AddNote(mockNote)
	assert.Nil(t, err)
	assert.NotNil(t, note)
}

func TestRepository_GetUserIDByUsername(t *testing.T) {
	mockUsername := "Ahmet"
	db, err := postgre.NewPostgreDB(postgre.Config)
	if err != nil {
		log.Fatalln(err)
	}
	repo := NewRepository(db)
	userID := repo.GetUserIDByUsername(mockUsername)
	assert.NotNil(t, userID)
}

func TestRepository_GetCurrentUserNotesByUserID(t *testing.T) {
	db, err := postgre.NewPostgreDB(postgre.Config)
	if err != nil {
		log.Fatalln(err)
	}
	repo := NewRepository(db)
	notes := repo.GetCurrentUserNotesByUserID(1)
	assert.NotNil(t, notes)
}