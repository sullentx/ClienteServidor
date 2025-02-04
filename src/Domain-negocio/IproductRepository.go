package domainnegocio

import entities "tienda/src/Products/Domain-negocio/entities"

//OCP Open Close Principle (Principio de abierto/cerrado)
//Se pueden agregar nuevos repositorios sin modificar el código de la aplicación
//ISP Interface Segregation Principle (Principio de segregación de interfaz)
// Se crean interfaces específicas para cada repositorio
//DIP Dependency Inversion Principle (Principio de inversión de dependencias)
//Los casos de uso dependen de las interfaces y no de las implementaciones
type IproductrRepositoy interface {
	Save(product entities.Product) error
	GetAll() ([]entities.Product, error)
	GetOne(id int) (entities.Product, error)
	Delete(id int) error
	Put(id int, product entities.Product) error
}
