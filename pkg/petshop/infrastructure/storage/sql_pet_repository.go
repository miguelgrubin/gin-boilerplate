package storage

import (
	"errors"
	"log"

	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/domain"
	"github.com/miguelgrubin/gin-boilerplate/pkg/shared"
	"gorm.io/gorm"
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
	var err error
	var prev PetEntity
	err = r.db.First(&prev, "id = ?", pet.ID.String()).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = r.db.Debug().Create(PetEntityFromDomain(pet)).Error
		if err != nil {
			log.Println(err.Error())
			return err
		}
		return nil
	}

	err = r.db.Debug().Save(PetEntityFromDomain(pet)).Error
	return err
}

func (r SQLPetRepository) FindOne(id shared.EntityID) (*domain.Pet, error) {
	var pet PetEntity
	err := r.db.Debug().Where("id = ?", id).Take(&pet).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, &domain.PetNotFound{ID: id.String()}
	}
	if err != nil {
		return nil, err
	}
	petDomain := PetEntityToDomain(pet)
	return &petDomain, nil
}

func (r SQLPetRepository) FindAll() ([]domain.Pet, error) {
	var pets []PetEntity
	err := r.db.Debug().Find(&pets).Error
	if err != nil {
		return nil, err
	}
	domainPets := make([]domain.Pet, len(pets))
	for i, v := range pets {
		domainPets[i] = PetEntityToDomain(v)
	}
	return domainPets, nil
}

func (r SQLPetRepository) Delete(id shared.EntityID) error {
	var pet PetEntity
	err := r.db.Debug().Where("id = ?", id).Delete(&pet).Error
	if err != nil {
		return errors.New("database error, please try again")
	}
	return nil
}

func Automigrate(db *gorm.DB) error {
	return db.AutoMigrate(&PetEntity{})
}
