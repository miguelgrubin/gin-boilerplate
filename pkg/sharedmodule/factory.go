// Package sharedmodule provides shared services, domains and errors.
package sharedmodule

import (
	"log"

	"github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule/services"
)

type SharedModuleServices struct {
	ConfigService services.ConfigService
	JWTService    services.JWTService
	HashService   services.HashService
	RedisService  services.RedisServiceInterface
	DBService     services.DBService
	RSAService    services.RSAService
}

func NewSharedModuleServices() SharedModuleServices {
	cs := services.NewConfigService()
	c, _ := cs.ReadConfig()
	ds := services.NewDBServiceGorm(c.Database)
	err := ds.Connect()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	rs := services.NewRedisService(c.Redis)
	rsa := services.NewRSAService(c.Jwt.Keys.Private, c.Jwt.Keys.Public)
	err = rsa.Read()
	if err != nil {
		log.Fatalf("failed to read RSA keys: %v", err)
	}

	return SharedModuleServices{
		ConfigService: cs,
		JWTService:    services.NewJWTServiceRSA(rs, rsa, c.Jwt),
		HashService:   services.NewHashServiceArgon2(),
		RedisService:  rs,
		DBService:     ds,
		RSAService:    rsa,
	}
}
