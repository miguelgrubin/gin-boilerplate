package pkg

import (
	"log"

	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop"
	"github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule"
)

/* MigrateAll runs all DB migrations */
func MigrateAll() error {
	config, err := ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db := sharedmodule.NewDbConnection(config.Database.Driver, config.Database.Address)

	ps := petshop.NewPetShopModule(db)
	err = ps.Automigrate(db)
	return err
}

func SeedAll() error {
	config, err := ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db := sharedmodule.NewDbConnection(config.Database.Driver, config.Database.Address)

	ps := petshop.NewPetShopModule(db)
	err = ps.Seed(db)
	return err
}
