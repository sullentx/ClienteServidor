package applicationnegocio

import (
	domainnegocio "tienda/src/Products/Domain-negocio"
	"tienda/src/Products/Domain-negocio/entities"
)

type CreateProductUseCase struct {
	repo domainnegocio.IproductrRepositoy
}

func NewCreateProduct(repo domainnegocio.IproductrRepositoy) *CreateProductUseCase {
	return &CreateProductUseCase{repo: repo}
}

func (uc *CreateProductUseCase) Execute(product entities.Product) error {
	return uc.repo.Save(product)
}
