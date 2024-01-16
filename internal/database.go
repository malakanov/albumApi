package internal

import "github.com/jinzhu/gorm"

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "albums.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = database
}
