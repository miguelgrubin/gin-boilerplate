package pkg

import (
	"log"

	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/repositories"
	"github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule"
)

/* MigrateAll runs all DB migrations */
func MigrateAll() error {
	config, err := ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db := sharedmodule.NewDbConnection(config.Database.Driver, config.Database.Address)

	err = repositories.Automigrate(db)
	return err
}

func SeedAll() error {
	config, err := ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db := sharedmodule.NewDbConnection(config.Database.Driver, config.Database.Address)

	_, err = repositories.SeedPets(db)
	return err
}
