package pkg

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/repositories"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/server"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/usecases"
	"github.com/miguelgrubin/gin-boilerplate/pkg/shared/storage"
	"github.com/spf13/viper"
)

func RunServer() {
	_, err := ReadConfig()
	if err != nil {
		log.Print("Config file not found: using default config")
	}

	r := SetupRouter()
	err = r.Run(viper.GetString("server.address"))
	if err != nil {
		log.Print(err)
	}
}

/* SetupRouter creates gin router instance with all app routes */
func SetupRouter() *gin.Engine {
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

/* NewServices inyects services on modules (petshop) */
func NewServices(r *gin.RouterGroup) {
	db := storage.NewDbConnection(
		viper.GetString("database.driver"),
		viper.GetString("database.address"),
	)
	pr := repositories.NewPetShopRepositories(db)
	pu := usecases.NewPetShopUseCases(pr)
	pc := server.NewPetShopController(pr, pu)
	pc.SetupRoutes(r)
}
