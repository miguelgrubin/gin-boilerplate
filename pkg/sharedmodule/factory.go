// Package sharedmodule provides shared services, domains and errors.
package sharedmodule

import (
	"github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule/services"
)

type SharedModuleServices struct {
	ConfigService services.ConfigService
	LoggerService services.LoggerService
	JWTService    services.JWTService
	HashService   services.HashService
	RedisService  services.RedisService
	DBService     services.DBService
	RSAService    services.RSAService
}

func NewSharedModuleServices() SharedModuleServices {
	cs := services.NewConfigService()
	c, _ := cs.ReadConfig()
	logger := services.NewLoggerService(c.Debug)

	ds := services.NewDBServiceGorm(c.Database)
	err := ds.Connect()
	if err != nil {
		logger.Fatal("failed to connect to database", services.Err(err))
	}

	rs := services.NewRedisService(c.Redis)
	rsa := services.NewRSAService(c.Jwt.Keys.Private, c.Jwt.Keys.Public)
	err = rsa.Read()
	if err != nil {
		logger.Fatal("failed to read RSA keys", services.Err(err))
	}

	return SharedModuleServices{
		ConfigService: cs,
		LoggerService: logger,
		JWTService:    services.NewJWTServiceRSA(rs, rsa, c.Jwt),
		HashService:   services.NewHashServiceArgon2(),
		RedisService:  rs,
		DBService:     ds,
		RSAService:    rsa,
	}
}
