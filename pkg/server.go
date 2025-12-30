package pkg

import (
	"log"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	app, err := NewApp()
	if err != nil {
		log.Fatal(err)
	}
	address := app.SharedServices.ConfigService.GetConfig().Server.Address

	app.SharedServices.DBService.Connect()
	defer app.SharedServices.DBService.Close()

	r := SetupRouter(app)
	err = r.Run(address)
	if err != nil {
		log.Print(err)
	}
}

/* SetupRouter creates gin router instance with all app routes */
func SetupRouter(app *App) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/health", func(c *gin.Context) {
		c.String(200, "Health check!")
	})
	v1 := r.Group("/v1")

	app.PetShopModule.SetupRoutes(v1)
	app.UsersModule.SetupRoutes(v1)
	return r
}
