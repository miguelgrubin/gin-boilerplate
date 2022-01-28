package repositories

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/miguelgrubin/gin-boilerplate/pkg/entities"
)

type IPetRepository interface {
	SavePet(*entities.Pet) (*entities.Pet, map[string]string)
	GetPet(uint64) (*entities.Pet, error)
	GetAllPets() ([]entities.Pet, error)
	UpdatePet(*entities.Pet) (*entities.Pet, map[string]string)
	DeletePet(uint64) error
}

type PetRepository struct {
	db *gorm.DB
}

var _ IPetRepository = &PetRepository{}

func NewPetRepository(db *gorm.DB) *PetRepository {
	return &PetRepository{db}
}

func (r *PetRepository) SavePet(pet *entities.Pet) (*entities.Pet, map[string]string) {
	dbErr := map[string]string{}
	err := r.db.Debug().Create(&pet).Error
	if err != nil {
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return pet, nil
}

func (r *PetRepository) GetPet(id uint64) (*entities.Pet, error) {
	var pet entities.Pet
	err := r.db.Debug().Where("id = ?", id).Take(&pet).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("pet not found")
	}
	return &pet, nil
}

func (r *PetRepository) GetAllPets() ([]entities.Pet, error) {
	var pets []entities.Pet
	err := r.db.Debug().Find(&pets).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("user not found")
	}
	return pets, nil
}

func (r *PetRepository) DeletePet(id uint64) error {
	var pet entities.Pet
	err := r.db.Debug().Where("id = ?", id).Delete(&pet).Error
	if err != nil {
		return errors.New("database error, please try again")
	}
	return nil
}

func (r *PetRepository) UpdatePet(pet *entities.Pet) (*entities.Pet, map[string]string) {
	dbErr := map[string]string{}
	err := r.db.Debug().Save(&pet).Error
	if err != nil {
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return pet, nil
}
