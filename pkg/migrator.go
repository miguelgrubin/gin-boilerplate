package pkg

import (
	"log"

	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop"
	"github.com/miguelgrubin/gin-boilerplate/pkg/shared/infrastructure"
)

func MigrateAll() {
	config, err := ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db := infrastructure.NewDbConnection(config.Database.Driver, config.Database.Address)
	petshop.NewPetShopMigrator(db)
}
