package server

import entities "main/src/Domain-negocio/Entities"

type InMemoryProductRepository struct {
	products []entities.Product
}

func NewInMemoryProductRepository() *InMemoryProductRepository {
	return &InMemoryProductRepository{
		products: []entities.Product{},
	}
}

func (repo *InMemoryProductRepository) Save(product entities.Product) error {
	repo.products = append(repo.products, product)
	return nil
}

func (repo *InMemoryProductRepository) GetAll() ([]entities.Product, error) {
	return repo.products, nil
}
