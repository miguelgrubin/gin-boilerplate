package repositories

import (
	"time"

	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/domain"
	"github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule"
)

func PetEntityToDomain(pe PetEntity) domain.Pet {
	return domain.Pet{
		ID:        sharedmodule.EntityID(pe.ID),
		Name:      pe.Name,
		Status:    pe.Status,
		CreatedAt: sharedmodule.DateTime(pe.CreatedAt),
		UpdatedAt: sharedmodule.DateTime(pe.UpdatedAt),
		DeletedAt: (*sharedmodule.DateTime)(pe.DeletedAt),
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
