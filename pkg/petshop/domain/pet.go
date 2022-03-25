package domain

import (
	"time"

	"github.com/miguelgrubin/gin-boilerplate/pkg/shared"
	uuid "github.com/satori/go.uuid"
)

type Pet struct {
	ID        shared.EntityId  `gorm:"primary_key;" json:"id"`
	Name      string           `gorm:"size:100;not null;" json:"name"`
	Status    string           `gorm:"size:100;not null;" json:"status"`
	CreatedAt shared.DateTime  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt shared.DateTime  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *shared.DateTime `json:"deleted_at,omitempty"`
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
