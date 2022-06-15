package repository

import (
	"github.com/ahmetzumber/rapid-note-cli/internal/postgre"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRepositoryCreatedByGivenDB(t *testing.T) {
	db, err := postgre.NewPostgreDB(postgre.Config)
	repo := NewRepository(db)
	assert.Nil(t, err)
	assert.NotNil(t, repo)
}