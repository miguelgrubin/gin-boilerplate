package pkg

import (
	"log"
)

/* MigrateAll runs all DB migrations */
func MigrateAll() error {
	app, err := NewApp()
	if err != nil {
		log.Fatal(err)
	}

	app.SharedServices.DBService.Connect()
	defer app.SharedServices.DBService.Close()

	app.PetShopModule.Automigrate()
	app.UsersModule.Automigrate()

	return err
}

func SeedAll() error {
	app, err := NewApp()
	if err != nil {
		log.Fatal(err)
	}

	app.SharedServices.DBService.Connect()
	defer app.SharedServices.DBService.Close()

	app.PetShopModule.Seed()
	return err
}
