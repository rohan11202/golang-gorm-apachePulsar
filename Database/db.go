package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func ConnectDB() {

	url := "host=postgres user=p4nda password=p4nda_pswd dbname=bookDB port=5432 sslmode=disable"
	d, error := gorm.Open(postgres.Open(url), &gorm.Config{})
	if error != nil {
		log.Fatal("Failed to connect to db ", error)
	}
	log.Println("DB connection established")
	db=d
	
}

func GetDB() *gorm.DB {
	return db
}
