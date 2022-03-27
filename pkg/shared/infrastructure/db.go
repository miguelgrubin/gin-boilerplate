package infrastructure

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDbConnection(Dbdriver, DBURL string) *gorm.DB {
	var db *gorm.DB
	if Dbdriver == "sqlite3" {
		var err error
		db, err = gorm.Open(sqlite.Open(DBURL), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
	}

	return db
}
