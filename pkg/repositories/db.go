package repositories

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/miguelgrubin/gin-boilerplate/pkg/domain/entity"
)

type Repositories struct {
	Pet PetRepository
	db  *gorm.DB
}

func NewRepositories(Dbdriver, DBURL string) (*Repositories, error) {
	db, err := gorm.Open(Dbdriver, DBURL)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)

	return &Repositories{
		Pet: NewPetRepository(db),
		db:  db,
	}, nil
}

// closes the  database connection
func (s *Repositories) Close() error {
	return s.db.Close()
}

// This migrate all tables
func (s *Repositories) Automigrate() error {
	return s.db.AutoMigrate(&entity.Pet{}).Error
}
