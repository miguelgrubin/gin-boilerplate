package storage

import (
	"errors"
	"log"
	"time"

	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/domain"
	"github.com/miguelgrubin/gin-boilerplate/pkg/shared"
	"gorm.io/gorm"
)

type PetEntity struct {
	ID        string     `gorm:"primary_key;"`
	Name      string     `gorm:"size:100;not null;"`
	Status    string     `gorm:"size:100;not null;"`
	CreatedAt time.Time  `gorm:"not null;"`
	UpdatedAt time.Time  `gorm:"not null;"`
	DeletedAt *time.Time ``
}

func (pe *PetEntity) TableName() string {
	return "pets"
}

func (pe *PetEntity) ToDomain() domain.Pet {
	return domain.Pet{
		ID:        shared.EntityId(pe.ID),
		Name:      pe.Name,
		Status:    pe.Status,
		CreatedAt: shared.DateTime(pe.CreatedAt),
		UpdatedAt: shared.DateTime(pe.UpdatedAt),
		DeletedAt: (*shared.DateTime)(pe.DeletedAt),
	}
}

func PetEntityFromDomain(p domain.Pet) PetEntity {
	return PetEntity{p.ID.AsString(), p.Name, p.Status, p.CreatedAt.AsTime(), p.UpdatedAt.AsTime(), (*time.Time)(p.DeletedAt)}
}

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
	err = r.db.First(&prev, "id = ?", pet.ID.AsString()).Error

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

func (r SQLPetRepository) FindOne(id shared.EntityId) (*domain.Pet, error) {
	var pet PetEntity
	err := r.db.Debug().Where("id = ?", id).Take(&pet).Error
	if err != nil {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("pet not found")
	}
	petDomain := pet.ToDomain()
	return &petDomain, nil
}

func (r SQLPetRepository) FindAll() ([]domain.Pet, error) {
	var pets []PetEntity
	err := r.db.Debug().Find(&pets).Error
	if err != nil {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}
	domainPets := make([]domain.Pet, len(pets))
	for i, v := range pets {
		domainPets[i] = v.ToDomain()
	}
	return domainPets, nil
}

func (r SQLPetRepository) Delete(id shared.EntityId) error {
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
