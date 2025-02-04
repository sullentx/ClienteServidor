package server

import entities "main/src/Domain-negocio/Entities"

type IproductrRepositoy interface {
	Save(product entities.Product) error
	GetAll() ([]entities.Product, error)
}
