package pkg

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop"
	"github.com/miguelgrubin/gin-boilerplate/pkg/shared/infrastructure"
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

func setupRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/health", func(c *gin.Context) {
		c.String(200, "Health check!")
	})
	v1 := r.Group("/v1")
	NewServices(v1)
	return r
}

func NewServices(r *gin.RouterGroup) {
	db := infrastructure.NewDbConnection(
		viper.GetString("database.driver"),
		viper.GetString("database.address"),
	)
	petshop.NewPetShopServer(db, r)
}
