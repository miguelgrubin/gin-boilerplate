package persistence

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/miguelgrubin/gin-boilerplate/pkg/domain/entity"
	"github.com/miguelgrubin/gin-boilerplate/pkg/domain/repository"
)

type PetRepo struct {
	db *gorm.DB
}

func NewPetRepository(db *gorm.DB) *PetRepo {
	return &PetRepo{db}
}

var _ repository.PetRepository = &PetRepo{}

func (r *PetRepo) SavePet(pet *entity.Pet) (*entity.Pet, map[string]string) {
	dbErr := map[string]string{}
	err := r.db.Debug().Create(&pet).Error
	if err != nil {
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return pet, nil
}

func (r *PetRepo) GetPet(id uint64) (*entity.Pet, error) {
	var pet entity.Pet
	err := r.db.Debug().Where("id = ?", id).Take(&pet).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("pet not found")
	}
	return &pet, nil
}

func (r *PetRepo) GetAllPets() ([]entity.Pet, error) {
	var pets []entity.Pet
	err := r.db.Debug().Find(&pets).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("user not found")
	}
	return pets, nil
}

func (r *PetRepo) DeletePet(id uint64) error {
	var pet entity.Pet
	err := r.db.Debug().Where("id = ?", id).Delete(&pet).Error
	if err != nil {
		return errors.New("database error, please try again")
	}
	return nil
}

func (r *PetRepo) UpdatePet(pet *entity.Pet) (*entity.Pet, map[string]string) {
	dbErr := map[string]string{}
	err := r.db.Debug().Save(&pet).Error
	if err != nil {
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return pet, nil
}
