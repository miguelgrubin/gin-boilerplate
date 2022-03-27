package server

import (
	"github.com/miguelgrubin/gin-boilerplate/pkg/petshop/domain"
)

type PetResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func PetReponseFromDomain(p domain.Pet) PetResponse {
	return PetResponse{p.ID.AsString(), p.Name, p.Status, p.CreatedAt.AsTime().String(), p.UpdatedAt.AsTime().String()}
}

func PetResponseListFromDomain(p []domain.Pet) []PetResponse {
	petList := make([]PetResponse, len(p))
	for i, v := range p {
		petList[i] = PetReponseFromDomain(v)
	}
	return petList
}

type PetCreateRequest struct {
	Name   string `json:"name" binding:"required"`
	Status string `json:"status" binding:"required"`
}

type PetUpdateRequest struct {
	Name   *string `json:"name"`
	Status *string `json:"status"`
}
