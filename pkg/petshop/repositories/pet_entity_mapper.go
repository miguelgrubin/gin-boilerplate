package repositories

import (
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/domain"
)

func PetEntityToDomain(pe PetEntity) domain.Pet {
	return domain.HydratePet(
		pe.ID,
		pe.Name,
		pe.Status,
		pe.CreatedAt,
		pe.UpdatedAt,
		pe.DeletedAt,
	)
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
