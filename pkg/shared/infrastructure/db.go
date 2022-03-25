package infrastructure

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewDbConnection(Dbdriver, DBURL string) (*gorm.DB, error) {
	db, err := gorm.Open(Dbdriver, DBURL)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	db.LogMode(true)

	return db, nil
}
