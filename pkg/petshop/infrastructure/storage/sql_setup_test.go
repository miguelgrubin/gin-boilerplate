package storage_test

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/miguelgrubin/gin-boilerplate/pkg"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/infrastructure/storage"
	"github.com/miguelgrubin/gin-boilerplate/pkg/shared/infrastructure"
	"gorm.io/gorm"
)

func DBConn() (*gorm.DB, error) {
	return LocalDatabase()
}

func LocalDatabase() (*gorm.DB, error) {
	var (
		_, b, _, _ = runtime.Caller(0)
		basepath   = filepath.Dir(b)
	)
	rootpath := fmt.Sprintf("%s/../../../../test", basepath)
	err := os.Chdir(rootpath)

	if err != nil {
		log.Println("Can not load test config file")
		return nil, err
	}

	appConfig, err := pkg.ReadConfig()
	if err != nil {
		return nil, err
	}

	db := infrastructure.NewDbConnection(appConfig.Database.Driver, appConfig.Database.Address)

	err = db.Migrator().DropTable(&storage.PetEntity{})
	if err != nil {
		return nil, err
	}

	err = storage.Automigrate(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}
