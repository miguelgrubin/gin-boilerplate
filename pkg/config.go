package pkg

import (
	"os"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Address string
	Https   bool
}

type DatabaseConfig struct {
	Driver  string
	Address string
}

type AppConfig struct {
	Server   ServerConfig
	Database DatabaseConfig
	Debug    bool
	Testing  bool
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

// ReadConfig from ./config/config.yaml or ./config/config_{{enviroment}}.yaml
func ReadConfig() (AppConfig, error) {
	defaultConfig()
	err := viper.ReadInConfig()
	if err != nil {
		return AppConfig{}, err
	}
	config := AppConfig{
		Server: ServerConfig{
			Address: viper.GetString("server.address"),
			Https:   viper.GetBool("server.https"),
		},
		Database: DatabaseConfig{
			Driver:  viper.GetString("database.driver"),
			Address: viper.GetString("database.address"),
		},
		Testing: viper.GetBool("testing"),
		Debug:   viper.GetBool("debug"),
	}
	return config, nil
}
