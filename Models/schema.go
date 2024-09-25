package schema

import (
	"log"

	"gorm.io/gorm"

	database "backend/Database"
)

var (
	db *gorm.DB
)

type Author struct {
	gorm.Model
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"not null"`
	Books []Book `gorm:"foreignKey:AuthorID"`
}

type Book struct {	
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	AuthorID    uint   `gorm:"not null"`
	Name        string `gorm:"not null"`
	Description string `gorm:"default:''"`
}

func Initiate() {
	database.ConnectDB()
	db = database.GetDB()
	if err := db.AutoMigrate(&Author{}, &Book{}); err != nil {
		log.Fatalf("Error Migrating Tables: %v", err)
		return
	}
	log.Println("Migration Successful")
}
