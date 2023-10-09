package server

import "github.com/miguelgrubin/gin-boilerplate/pkg/petshop/domain"

func PetReponseFromDomain(p domain.Pet) PetResponse {
	return PetResponse{p.ID.String(), p.Name, p.Status, p.CreatedAt.Time().String(), p.UpdatedAt.Time().String()}
}

func PetResponseListFromDomain(p []domain.Pet) []PetResponse {
	petList := make([]PetResponse, len(p))
	for i, v := range p {
		petList[i] = PetReponseFromDomain(v)
	}
	return petList
}
