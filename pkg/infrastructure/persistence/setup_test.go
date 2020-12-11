package persistence

import (
	"log"

	"github.com/miguelgrubin/gin-boilerplate/pkg"
	"github.com/miguelgrubin/gin-boilerplate/pkg/domain/entity"

	"github.com/jinzhu/gorm"
)

func DBConn() (*gorm.DB, error) {
	return LocalDatabase()
}

func LocalDatabase() (*gorm.DB, error) {
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

	err = conn.DropTableIfExists(&entity.Pet{}).Error
	if err != nil {
		return nil, err
	}
	err = conn.Debug().AutoMigrate(
		entity.Pet{},
	).Error
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func seedPet(db *gorm.DB) (*entity.Pet, error) {
	pet := &entity.Pet{
		ID:     1,
		Name:   "Fluffy",
		Status: "happy",
	}
	err := db.Create(&pet).Error
	if err != nil {
		return nil, err
	}
	return pet, nil
}

func seedPets(db *gorm.DB) ([]entity.Pet, error) {
	pets := []entity.Pet{
		{
			ID:     1,
			Name:   "Tommy",
			Status: "bored",
		},
		{
			ID:     2,
			Name:   "Katty",
			Status: "sleeping",
		},
	}
	for _, v := range pets {
		err := db.Create(&v).Error
		if err != nil {
			return nil, err
		}
	}
	return pets, nil
}
