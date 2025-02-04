package domainnegocio

import entities "main/src/Domain-negocio/Entities"

type IproductrRepositoy interface {
	Save(product entities.Product) error
	GetAll() ([]entities.Product, error)
	GetOne(id int) (entities.Product, error)
	Delete(id int) error
	Put(id int, product entities.Product) error
}
