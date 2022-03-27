package server

import (
	"time"

	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/domain"
)

type PetResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func PetReponseFromDomain(p domain.Pet) PetResponse {
	return PetResponse{p.ID.AsString(), p.Name, p.Status, p.CreatedAt.AsTime(), p.UpdatedAt.AsTime()}
}
