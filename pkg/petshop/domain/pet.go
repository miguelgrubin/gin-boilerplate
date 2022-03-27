package domain

import (
	"time"

	"github.com/miguelgrubin/gin-boilerplate/pkg/shared"
	uuid "github.com/satori/go.uuid"
)

type Pet struct {
	ID        shared.EntityId
	Name      string
	Status    string
	CreatedAt shared.DateTime
	UpdatedAt shared.DateTime
	DeletedAt *shared.DateTime
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
}

func NewPet(payload CreatePetParams) Pet {
	id := uuid.NewV4().String()
	pet := Pet{
		ID:        shared.EntityId(id),
		Name:      payload.Name,
		Status:    payload.Status,
		UpdatedAt: shared.DateTime(time.Now()),
		CreatedAt: shared.DateTime(time.Now()),
		DeletedAt: nil,
	}
	return pet
}
