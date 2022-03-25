package storage_test

import (
	"log"
	"os"

	"github.com/miguelgrubin/gin-boilerplate/pkg"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/domain"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/infrastructure/storage"

	"github.com/jinzhu/gorm"
)

func DBConn() (*gorm.DB, error) {
	return LocalDatabase()
}

func LocalDatabase() (*gorm.DB, error) {
	err := os.Chdir("../../../../test")
	if err != nil {
		log.Println("Can not load test config file")
	}
	appConfig, err := pkg.ReadConfig()
	if err != nil {
		return nil, err
	} else {
		log.Println("APP CONFIG READED")
	}

	conn, err := gorm.Open(appConfig.Database.Driver, appConfig.Database.Driver)
	if err != nil {
		return nil, err
	} else {
		log.Println("CONNECTED TO: ", appConfig.Database.Driver)
	}

	err = conn.DropTableIfExists(&domain.Pet{}).Error
	if err != nil {
		return nil, err
	}
	err = storage.Automigrate(conn)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
