package postgre

import (
	"fmt"
	"github.com/ahmetzumber/rapid-note-cli/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Config = config.DB{
	Username: "rapid",
	Password: "rapid1",
	Driver:   "postgres",
	DBName:   "postgres-rapid",
	Host:     "localhost",
	Port:     "5432",
}

// DSN: data source name --> creating connection url parametrically
func createDSN(config config.DB) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		 				config.Host, config.Port, config.Username, config.Password, config.DBName)
}

func NewPostgreDB(config config.DB) (*gorm.DB, error) {
	dsn := createDSN(config)
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}
	return db, nil
}
