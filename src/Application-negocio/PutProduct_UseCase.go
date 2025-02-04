package applicationnegocio

import (
	domainnegocio "tienda/src/Products/Domain-negocio"
	"tienda/src/Products/Domain-negocio/entities"
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
