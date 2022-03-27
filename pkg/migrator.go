package pkg

import (
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop"
	"github.com/miguelgrubin/gin-boilerplate/pkg/shared/infrastructure"
	"github.com/spf13/viper"
)

func MigrateAll() {
	ReadConfig()
	db := infrastructure.NewDbConnection(
		viper.GetString("database.driver"),
		viper.GetString("database.address"),
	)
	petshop.NewPetShopMigrator(db)
}
