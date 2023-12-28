package postgres

import (
	"fmt"

	"github.com/akmalulginan/carjod-be/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Credential struct {
	Host     string
	Port     string
	DBName   string
	Username string
	Password string
}

func Init(cred Credential) (*gorm.DB, error) {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		cred.Host, cred.Username, cred.Password, cred.DBName, cred.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	migrate(db)

	return db, nil
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(
		domain.User{},
		domain.Match{},
	)
}
