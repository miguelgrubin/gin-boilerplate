package pkg

import (
	"time"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.String(200, "Health check")
	})
	return r
}

type EntityId struct {
	ID uint64
}

type BaseEntity struct {
	ID        uint64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type IAppModule interface {
	GetRepository
}

type AppModule struct {
	routes       gin.RouterGroup
	service      interface{}
	repositories map[string]interface{}
}

type IRepository interface {
	Create(payload interface{}) (BaseEntity, map[string]string)
	Get(id EntityId) (BaseEntity, error)
	FindOne(filters interface{}) (BaseEntity, error)
	FindAll(filters interface{}) ([]BaseEntity, error)
	Update(id EntityId, payload interface{}) (BaseEntity, map[string]string)
	Delete(id EntityId) error
}

type AppService struct {
	repository IRepository
}

func NewService() AppService {

}
