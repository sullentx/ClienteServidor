package applicationnegocio

import (
	domainnegocio "main/src/Domain-negocio"
	entities "main/src/Domain-negocio/Entities"
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
