package pkg

import (
	"log"

	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/infrastructure/storage"
	"github.com/miguelgrubin/gin-boilerplate/pkg/shared/infrastructure"
)

/* MigrateAll runs all DB migrations */
func MigrateAll() {
	config, err := ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db := infrastructure.NewDbConnection(config.Database.Driver, config.Database.Address)
	petshop.NewPetShopMigrator(db)
}

func SeedAll() {
	config, err := ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db := infrastructure.NewDbConnection(config.Database.Driver, config.Database.Address)

	_, err = storage.SeedPets(db)
	if err != nil {
		log.Fatal(err)
	}
}
