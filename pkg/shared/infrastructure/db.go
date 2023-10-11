// Package infrastructure provides common services such as database connections, auth, etc.
package infrastructure

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
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

	if Dbdriver == "mysql" {
		var err error
		db, err = gorm.Open(mysql.Open(DBURL), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
	}

	if Dbdriver == "mysql" {
		var err error
		db, err = gorm.Open(mysql.Open(DBURL), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
	}

	if Dbdriver == "postgres" {
		var err error
		db, err = gorm.Open(postgres.Open(DBURL), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
	}

	return db
}
