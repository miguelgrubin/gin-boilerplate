// Package usecases provides use cases for petshop module.
package usecases

import "github.com/miguelgrubin/gin-boilerplate/pkg/petshop/repositories"

type PetShopUseCases struct {
	Pet PetUseCasesInterface
}

func NewPetShopUseCases(r repositories.PetShopRepositories) PetShopUseCases {
	petUseCases := NewPetUseCases(r.Pet)
	return PetShopUseCases{
		Pet: &petUseCases,
	}
}
