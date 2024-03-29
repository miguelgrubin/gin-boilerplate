// Package pkg is the root and provides config service and entry points for running the application.
package pkg

import (
	"os"

	"github.com/spf13/viper"
)

/* ServerConfig contains info about how the app will be served */
type ServerConfig struct {
	Address string
	HTTPS   bool
}

/* DatabaseConfig contains info about SQL Database */
type DatabaseConfig struct {
	Driver  string // (postgres, mysql, sqlite3)
	Address string
}

/* AppConfig contains all configuration to run app instance */
type AppConfig struct {
	Server   ServerConfig
	Database DatabaseConfig
	Debug    bool
	Testing  bool
}

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
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	viper.SetDefault("server.address", "0.0.0.0:8080")
	viper.SetDefault("server.https", false)
	viper.SetDefault("database.driver", "sqlite3")
	viper.SetDefault("database.address", "database.db")
	viper.SetDefault("debug", false)
	viper.SetDefault("testing", false)
}

func configFactory() AppConfig {
	return AppConfig{
		Server: ServerConfig{
			Address: viper.GetString("server.address"),
			HTTPS:   viper.GetBool("server.https"),
		},
		Database: DatabaseConfig{
			Driver:  viper.GetString("database.driver"),
			Address: viper.GetString("database.address"),
		},
		Testing: viper.GetBool("testing"),
		Debug:   viper.GetBool("debug"),
	}
}

// ReadConfig from ./config.yaml or ./config_{{enviroment}}.yaml
func ReadConfig() (AppConfig, error) {
	defaultConfig()
	err := viper.ReadInConfig()
	if err != nil {
		return configFactory(), err
	}
	return configFactory(), nil
}

// WriteConfig to ./config.yaml or ./config_{{environment}}.yaml
func WriteConfig() error {
	defaultConfig()
	return viper.SafeWriteConfig()
}
