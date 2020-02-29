package pkg

import (
	"os"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	address string
	https   bool
}

type DatabaseConfig struct {
	driver  string
	address string
}

type AppConfig struct {
	server   ServerConfig
	database DatabaseConfig
	debug    bool
	testing  bool
}

// @TODO add redis config
// @TODO DDD https://dev.to/stevensunflash/using-domain-driven-design-ddd-in-golang-3ee5

func isValidEnviroment(env string) bool {
	switch env {
	case
		"test",
		"local",
		"dev":
		return true
	}
	return false
}

func defaultConfig() {
	srcName := "config"
	if isValidEnviroment(os.Getenv("APP_ENV")) {
		srcName = srcName + "_" + os.Getenv("APP_ENV")
	}
	viper.SetConfigName(srcName)
	viper.AddConfigPath("../config")
	viper.SetConfigType("yaml")

	viper.SetDefault("server.address", "0.0.0.0:8080")
	viper.SetDefault("server.https", false)
	viper.SetDefault("database.driver", "sqlite3")
	viper.SetDefault("database.address", "database.db")
	viper.SetDefault("debug", false)
	viper.SetDefault("testing", false)
}

func ReadConfig() (AppConfig, error) {
	defaultConfig()
	err := viper.ReadInConfig()
	if err != nil {
		return AppConfig{}, err
	}
	config := AppConfig{
		server: ServerConfig{
			address: viper.GetString("server.address"),
			https:   viper.GetBool("server.https"),
		},
		database: DatabaseConfig{
			driver:  viper.GetString("database.driver"),
			address: viper.GetString("database.address"),
		},
		testing: viper.GetBool("testing"),
		debug:   viper.GetBool("debug"),
	}
	return config, nil
}
