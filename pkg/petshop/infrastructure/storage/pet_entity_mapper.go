package storage

import (
	"time"

	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/domain"
	"github.com/miguelgrubin/gin-boilerplate/pkg/shared"
)

func PetEntityToDomain(pe PetEntity) domain.Pet {
	return domain.Pet{
		ID:        shared.EntityID(pe.ID),
		Name:      pe.Name,
		Status:    pe.Status,
		CreatedAt: shared.DateTime(pe.CreatedAt),
		UpdatedAt: shared.DateTime(pe.UpdatedAt),
		DeletedAt: (*shared.DateTime)(pe.DeletedAt),
	}
}

func PetEntityFromDomain(p domain.Pet) PetEntity {
	return PetEntity{
		ID:        p.ID.String(),
		Name:      p.Name,
		Status:    p.Status,
		CreatedAt: p.CreatedAt.Time(),
		UpdatedAt: p.UpdatedAt.Time(),
		DeletedAt: (*time.Time)(p.DeletedAt),
	}
}
