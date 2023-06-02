package utils

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func InitDBTest() {
	var err error
	DB, err = gorm.Open(sqlite.Open("databaseTest.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func InitDBNilTest() {
	var err error
	DB, err = gorm.Open(sqlite.Open("databaseNilTest.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}
