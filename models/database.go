package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Database struct {
	db     *gorm.DB
	config *DBConfig
}

type DBConfig struct {
	Host     *string
	User     *string
	DBName   *string
	Password *string
	Port     *uint
}

func NewDatabase(config *DBConfig) *Database {
	var db Database
	db.config = config
	connectionParams := fmt.Sprintf("host=%v port=%d user=%v dbname=%v sslmode=disable password=%v",
		*db.config.Host,
		*db.config.Port,
		*db.config.User,
		*db.config.DBName,
		*db.config.Password,
	)
	var err error
	db.db, err = gorm.Open("postgres", connectionParams)
	if err != nil {
		panic(err)
	}
	log.Println("Connection:", connectionParams)
	db.db.AutoMigrate(&Comment{})
	return &db
}

func (db *Database) Close() {
	db.Close()
}

func DBError(err error) error {
	return fmt.Errorf("database error: %v", err)
}
