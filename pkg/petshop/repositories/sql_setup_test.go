package repositories_test

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/miguelgrubin/gin-boilerplate/pkg"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/repositories"
	"github.com/miguelgrubin/gin-boilerplate/pkg/shared/storage"
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

	db := storage.NewDbConnection(appConfig.Database.Driver, appConfig.Database.Address)

	err = db.Migrator().DropTable(&repositories.PetEntity{})
	if err != nil {
		return nil, err
	}

	err = repositories.Automigrate(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}
