package postgre

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateNewPostgreDBConnectionSuccesfully(t *testing.T){
	db, err := NewPostgreDB(Config)
	assert.Nil(t, err)
	assert.NotNil(t, db)
}