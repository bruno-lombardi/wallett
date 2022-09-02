package sqlite

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Connect(url string) *gorm.DB {
	if db, err = gorm.Open(sqlite.Open(url), &gorm.Config{}); err != nil {
		log.Fatalf("it was not possible to connect with db %v", url)
	} else {
		return db
	}
	return nil
}

func GetDB() *gorm.DB {
	return db
}
