package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule"
)

const (
	PetUpdated = "pet.updated"
	PetCreated = "pet.created"
)

type Pet struct {
	ID            string
	Name          string
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
	eventRegistry *sharedmodule.EventRegistry
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

func NewPet(payload CreatePetParams) Pet {
	id := uuid.New().String()
	pet := Pet{
		ID:            id,
		Name:          payload.Name,
		Status:        payload.Status,
		UpdatedAt:     time.Now(),
		CreatedAt:     time.Now(),
		DeletedAt:     nil,
		eventRegistry: sharedmodule.NewEventRegistry(),
	}
	pet.eventRegistry.AddEvent(PetCreated)
	return pet
}
