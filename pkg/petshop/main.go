package petshop

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/application"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/infrastructure/server"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/infrastructure/storage"
)

func NewPetShopServer(db *gorm.DB, r *gin.RouterGroup) {
	pr := storage.NewPetRepository(db)
	useCases := application.NewPetUseCases(pr)
	server.NewRouterGroup(r, &useCases)
}

func NewPetShopMigrator(db *gorm.DB) {

}

func NewPetShopSeeder(db *gorm.DB) {

}
