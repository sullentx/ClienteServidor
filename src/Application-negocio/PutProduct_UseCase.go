package applicationnegocio

import (
	domainnegocio "main/src/Domain-negocio"
	entities "main/src/Domain-negocio/Entities"
)

type PutProductUseCase struct {
	repo domainnegocio.IproductrRepositoy
}

func PutProduct(repo domainnegocio.IproductrRepositoy) *PutProductUseCase {
	return &PutProductUseCase{repo: repo}
}

func (uc *PutProductUseCase) Execute(id int, product entities.Product) error {
	return uc.repo.Put(id, product)
}
