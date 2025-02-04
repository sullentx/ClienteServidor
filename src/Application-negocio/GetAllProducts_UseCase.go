package applicationnegocio

import (
	domainnegocio "tienda/src/Products/Domain-negocio"
	"tienda/src/Products/Domain-negocio/entities"
)

type GetAllProductsUseCase struct {
	repo domainnegocio.IproductrRepositoy
}

// constructor para usar los metodos de la estructura GetAllProductsUseCase
func GetAllProducts(repo domainnegocio.IproductrRepositoy) *GetAllProductsUseCase {
	return &GetAllProductsUseCase{repo: repo}
}

// funcion para ejecutar el metodo GetAll de la estructura GetAllProductsUseCase
func (uc *GetAllProductsUseCase) Execute() ([]entities.Product, error) {
	return uc.repo.GetAll()
}
