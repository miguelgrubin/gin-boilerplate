package domain

import (
	"time"

	"github.com/google/uuid"
	sd "github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule/domain"
)

const (
	PetUpdated = "pet.updated"
	PetCreated = "pet.created"
)

// Pet status constants
const (
	PetStatusAvailable = "available"
	PetStatusPending   = "pending"
	PetStatusSold      = "sold"
)

type Pet struct {
	ID            string
	Name          string
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
	eventRegistry *sd.EventRegistry
}

type CreatePetParams struct {
	Name   string
	Status string
}

type UpdatePetParams struct {
	Name   *string
	Status *string
}

type Pets []Pet

func CreatePet(payload CreatePetParams) Pet {
	id := uuid.New().String()
	pet := Pet{
		ID:            id,
		Name:          payload.Name,
		Status:        payload.Status,
		UpdatedAt:     time.Now(),
		CreatedAt:     time.Now(),
		DeletedAt:     nil,
		eventRegistry: sd.NewEventRegistry(),
	}
	pet.eventRegistry.AddEvent(PetCreated)
	return pet
}

func (p *Pet) Update(payload UpdatePetParams) {
	if payload.Name != nil {
		p.Name = *payload.Name
	}
	if payload.Status != nil {
		p.Status = *payload.Status
	}
	p.UpdatedAt = time.Now()
	p.eventRegistry.AddEvent(PetUpdated)
}

// newPet creates a Pet with an initialized event registry.
// This is intended for internal use when hydrating entities from persistence.
// For creating new pets, use CreatePet instead.
func newPet() Pet {
	return Pet{
		eventRegistry: sd.NewEventRegistry(),
	}
}

// HydratePet creates a Pet from persistence data.
// This should only be used by repository mappers to reconstruct entities from storage.
// For creating new pets in the application, use CreatePet instead.
func HydratePet(
	id, name, status string,
	createdAt, updatedAt time.Time,
	deletedAt *time.Time,
) Pet {
	pet := newPet()
	pet.ID = id
	pet.Name = name
	pet.Status = status
	pet.CreatedAt = createdAt
	pet.UpdatedAt = updatedAt
	pet.DeletedAt = deletedAt
	return pet
}
