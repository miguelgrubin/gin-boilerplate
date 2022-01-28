package pkg

import (
	"log"

	"github.com/spf13/viper"
)

func RunServer() {
	_, err := ReadConfig()
	if err != nil {
		log.Print("Config file not found: using default config")
	}
	r := setupRouter()
	err = r.Run(viper.GetString("server.address"))
	if err != nil {
		log.Print(err)
	}
}
