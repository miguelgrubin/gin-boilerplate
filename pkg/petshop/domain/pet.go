package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/miguelgrubin/gin-boilerplate/pkg/shared"
	"github.com/miguelgrubin/gin-boilerplate/pkg/shared/domain"
)

const (
	PetUpdated = "pet.updated"
	PetCreated = "pet.created"
)

type Pet struct {
	ID            shared.EntityID
	Name          string
	Status        string
	CreatedAt     shared.DateTime
	UpdatedAt     shared.DateTime
	DeletedAt     *shared.DateTime
	eventRegistry *domain.EventRegistry
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
	p.UpdatedAt = shared.DateTime(time.Now())
	p.eventRegistry.AddEvent(PetUpdated)
}

func NewPet(payload CreatePetParams) Pet {
	id := uuid.New().String()
	pet := Pet{
		ID:            shared.EntityID(id),
		Name:          payload.Name,
		Status:        payload.Status,
		UpdatedAt:     shared.DateTime(time.Now()),
		CreatedAt:     shared.DateTime(time.Now()),
		DeletedAt:     nil,
		eventRegistry: domain.NewEventRegistry(),
	}
	pet.eventRegistry.AddEvent(PetCreated)
	return pet
}
