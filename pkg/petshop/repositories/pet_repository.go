package repositories

import (
	"errors"

	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/domain"
	"gorm.io/gorm"
)

type PetRepository interface {
	FindOne(string) (*domain.Pet, error)
	FindAll() ([]domain.Pet, error)
	Save(domain.Pet) error
	Delete(string) error
	Automigrate() error
	Seed() ([]domain.Pet, error)
}

type SQLPetRepository struct {
	db *gorm.DB
}

func NewPetRepository(db *gorm.DB) PetRepository {
	var petRepository PetRepository = SQLPetRepository{db}
	return petRepository
}

var _ PetRepository = &SQLPetRepository{}

func (r SQLPetRepository) Save(pet domain.Pet) error {
	var err error
	var prev PetEntity
	err = r.db.First(&prev, "id = ?", pet.ID).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = r.db.Create(PetEntityFromDomain(pet)).Error
		return err
	}

	err = r.db.Save(PetEntityFromDomain(pet)).Error
	return err
}

func (r SQLPetRepository) FindOne(id string) (*domain.Pet, error) {
	var pet PetEntity
	err := r.db.Where("id = ?", id).Take(&pet).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = &domain.PetNotFound{ID: id}
	}
	if err != nil {
		return nil, err
	}
	petDomain := PetEntityToDomain(pet)
	return &petDomain, nil
}

func (r SQLPetRepository) FindAll() ([]domain.Pet, error) {
	var pets []PetEntity
	err := r.db.Find(&pets).Error
	domainPets := make([]domain.Pet, len(pets))
	for i, v := range pets {
		domainPets[i] = PetEntityToDomain(v)
	}
	return domainPets, err
}

func (r SQLPetRepository) Delete(id string) error {
	var pet PetEntity
	err := r.db.Where("id = ?", id).Delete(&pet).Error
	return err
}

func (r SQLPetRepository) Automigrate() error {
	return r.db.AutoMigrate(&PetEntity{})
}

func (r SQLPetRepository) Seed() ([]domain.Pet, error) {
	pets := []domain.Pet{
		domain.NewPet(domain.CreatePetParams{
			Name:   "Tommy",
			Status: "bored",
		}),
		domain.NewPet(domain.CreatePetParams{
			Name:   "Katty",
			Status: "sleeping",
		}),
	}
	for _, v := range pets {
		err := r.Save(v)
		if err != nil {
			return nil, err
		}
	}
	return pets, nil
}
