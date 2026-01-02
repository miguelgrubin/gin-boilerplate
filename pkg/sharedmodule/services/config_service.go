// Package services contains services used across modules.
package services

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type ConfigService interface {
	ReadConfig() (AppConfig, error)
	WriteConfig() error
	GetConfig() AppConfig
}

/* AppConfig contains all configuration to run app instance */
type AppConfig struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
	Jwt      JwtConfig
	Debug    bool
	Testing  bool
}

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

type RedisConfig struct {
	Address  string
	Password string
	DB       int
}

/* JwtConfig contains info about JWT security settings */
type JwtConfig struct {
	Keys       JwtKeys
	Exp        int
	ExpRefresh int
}

type JwtKeys struct {
	Private string
	Public  string
}

type ConfigServiceViper struct {
	configPath *string
	config     AppConfig
}

func NewConfigService() *ConfigServiceViper {
	return &ConfigServiceViper{configPath: nil}
}

func NewConfigServiceWithPath(configPath *string) *ConfigServiceViper {
	return &ConfigServiceViper{configPath: configPath}
}

func (c *ConfigServiceViper) ReadConfig() (AppConfig, error) {
	c.defaultConfig()
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
		return AppConfig{}, err
	}
	c.config = configFactory()
	return c.config, nil
}

func (c *ConfigServiceViper) WriteConfig() error {
	c.defaultConfig()
	return viper.SafeWriteConfig()
}

func (c *ConfigServiceViper) GetConfig() AppConfig {
	return c.config
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

func (c *ConfigServiceViper) defaultConfig() {
	srcName := "config"
	if isValidEnviroment(os.Getenv("APP_ENV")) {
		srcName = srcName + "_" + os.Getenv("APP_ENV")
	}
	viper.SetConfigName(srcName)
	if c.configPath != nil {
		viper.AddConfigPath(*c.configPath)
	} else {
		viper.AddConfigPath(".")
	}
	viper.SetConfigType("yaml")

	viper.SetDefault("server.address", "0.0.0.0:8080")
	viper.SetDefault("server.https", false)
	viper.SetDefault("database.driver", "sqlite3")
	viper.SetDefault("database.address", "database.db")
	viper.SetDefault("redis.address", "0.0.0.0:6379")
	viper.SetDefault("redis.password", "")
	viper.SetDefault("redis.db", 0)
	viper.SetDefault("jwt.keys.private", "")
	viper.SetDefault("jwt.keys.public", "")
	viper.SetDefault("jwt.exp", 60)
	viper.SetDefault("jwt.exp_refresh", 10080)
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
		Redis: RedisConfig{
			Address:  viper.GetString("redis.address"),
			Password: viper.GetString("redis.password"),
			DB:       viper.GetInt("redis.db"),
		},
		Jwt: JwtConfig{
			Keys: JwtKeys{
				Private: viper.GetString("jwt.keys.private"),
				Public:  viper.GetString("jwt.keys.public"),
			},
			Exp:        viper.GetInt("jwt.exp"),
			ExpRefresh: viper.GetInt("jwt.exp_refresh"),
		},
		Testing: viper.GetBool("testing"),
		Debug:   viper.GetBool("debug"),
	}
}
