package applicationnegocio

import (
	domainnegocio "main/src/Domain-negocio"
	entities "main/src/Domain-negocio/Entities"
)

type GetOneProductUseCase struct {
	repo domainnegocio.IproductrRepositoy
}

// constructor
func GetOneProduct(repo domainnegocio.IproductrRepositoy) *GetOneProductUseCase {
	return &GetOneProductUseCase{repo: repo}
}

func (uc *GetOneProductUseCase) Execute(id int) (entities.Product, error) {
	return uc.repo.GetOne(id)
}
