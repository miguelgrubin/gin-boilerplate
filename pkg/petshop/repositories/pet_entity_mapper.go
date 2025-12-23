package repositories

import (
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/domain"
)

func PetEntityToDomain(pe PetEntity) domain.Pet {
	return domain.Pet{
		ID:        pe.ID,
		Name:      pe.Name,
		Status:    pe.Status,
		CreatedAt: pe.CreatedAt,
		UpdatedAt: pe.UpdatedAt,
		DeletedAt: pe.DeletedAt,
	}
}

func PetEntityFromDomain(p domain.Pet) PetEntity {
	return PetEntity{
		ID:        p.ID,
		Name:      p.Name,
		Status:    p.Status,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		DeletedAt: p.DeletedAt,
	}
}
