package pkg

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.String(200, "Health check")
	})
	return r
}

func RunServer() {
	r := setupRouter()
	r.Run(viper.GetString("server.address"))
}
