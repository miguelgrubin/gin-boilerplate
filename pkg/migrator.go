package pkg

import (
	"log"

	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/repositories"
	"github.com/miguelgrubin/gin-boilerplate/pkg/shared/storage"
)

/* MigrateAll runs all DB migrations */
func MigrateAll() {
	config, err := ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db := storage.NewDbConnection(config.Database.Driver, config.Database.Address)
	petshop.NewPetShopMigrator(db)
}

func SeedAll() {
	config, err := ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db := storage.NewDbConnection(config.Database.Driver, config.Database.Address)

	_, err = repositories.SeedPets(db)
	if err != nil {
		log.Fatal(err)
	}
}
