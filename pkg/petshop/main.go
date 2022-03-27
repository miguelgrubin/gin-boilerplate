package petshop

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/application"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/infrastructure/server"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/infrastructure/storage"
	"gorm.io/gorm"
)

func NewPetShopServer(db *gorm.DB, r *gin.RouterGroup) {
	pr := storage.NewPetRepository(db)
	useCases := application.NewPetUseCases(pr)
	server.NewRouterGroup(r, &useCases)
}

func NewPetShopMigrator(db *gorm.DB) {
	err := storage.Automigrate(db)

	if err != nil {
		log.Print(err)
	}
}

func NewPetShopSeeder(db *gorm.DB) {

}
