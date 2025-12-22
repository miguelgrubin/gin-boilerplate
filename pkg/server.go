package pkg

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/miguelgrubin/gin-boilerplate/docs"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/repositories"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/server"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/usecases"
	"github.com/miguelgrubin/gin-boilerplate/pkg/shared/storage"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
/* RunServer reads config and runs http/https server */
func RunServer() {
	_, err := ReadConfig()
	if err != nil {
		log.Print("Config file not found: using default config")
	}

	docs.SwaggerInfo.Title = "Petshop API"
	docs.SwaggerInfo.Description = "This is a sample server Petshop server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

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
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
