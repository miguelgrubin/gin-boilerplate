package storage

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/domain"
	"github.com/miguelgrubin/gin-boilerplate/pkg/shared"
)

type SQLPetRepository struct {
	db *gorm.DB
}

func NewPetRepository(db *gorm.DB) domain.PetRepository {
	var petRepository domain.PetRepository = SQLPetRepository{db}
	return petRepository
}

var _ domain.PetRepository = &SQLPetRepository{}

func (r SQLPetRepository) Save(pet domain.Pet) error {
	err := r.db.Debug().Save(pet).Error
	return err
}

func (r SQLPetRepository) FindOne(id shared.EntityId) (*domain.Pet, error) {
	var pet domain.Pet
	err := r.db.Debug().Where("id = ?", id).Take(&pet).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("pet not found")
	}
	return &pet, nil
}

func (r SQLPetRepository) FindAll() ([]domain.Pet, error) {
	var pets []domain.Pet
	err := r.db.Debug().Find(&pets).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("user not found")
	}
	return pets, nil
}

func (r SQLPetRepository) Delete(id shared.EntityId) error {
	var pet domain.Pet
	err := r.db.Debug().Where("id = ?", id).Delete(&pet).Error
	if err != nil {
		return errors.New("database error, please try again")
	}
	return nil
}

func Automigrate(db *gorm.DB) error {
	return db.Debug().AutoMigrate(&domain.Pet{}).Error
}
