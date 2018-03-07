package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

func dbConnect() *gorm.DB {
	// db, err := gorm.Open("postgres", "user=bstojchevski password=Password1 dbname=postgres sslmode=disable")
	db, err := gorm.Open("sqlite3", "user.db")
	if err != nil {
		panic(fmt.Sprintf("Cannot connect to the database!\n ERROR: '%v'", err))
	}

	return db
}

func dbDisconnect(db *gorm.DB) {
	db.Close()
}

func dbAutoMigrate() {

	db = dbConnect()
	defer dbDisconnect(db)

	db.AutoMigrate(&User{})
}
