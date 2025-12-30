package sharedmodule

import "github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule/services"

type SharedModuleServices struct {
	ConfigService services.ConfigService
	JWTService    services.JWTService
	HashService   services.HashService
	RedisService  services.RedisServiceInterface
	DBService     services.DBService
}

func NewSharedModuleServices() SharedModuleServices {
	cs := services.NewConfigService()
	c, _ := cs.ReadConfig()

	rs := services.NewRedisService(c.Redis)

	return SharedModuleServices{
		ConfigService: cs,
		JWTService:    services.NewJWTServiceRSA(rs, c.Jwt),
		HashService:   services.NewHashServiceArgon2(),
		RedisService:  rs,
		DBService:     services.NewDBServiceGorm(c.Database),
	}
}
