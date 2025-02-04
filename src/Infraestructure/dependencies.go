package infraestructure

import (
	applicationnegocio "main/src/Application-negocio"
	core "main/src/Core"
	controller "main/src/Infraestructure/Controller"
)

var (
	PostProductsHandler *controller.PostProductsHandler
	GetProductsHandler  *controller.GetProductsHandler
	GetOneProductHadler *controller.GetOneProductHandler
	DeleteProductHadler *controller.DeleteProductHandler
	PutProductHadler    *controller.PutProductHandler
)

func Init() {
	// Inicializar la conexión a la base de datos
	core.InitPostgres()
	db := core.GetDB()

	// Crear instancias del repositorio y casos de uso
	productRepo := NewPostgresProductRepository(db)
	createProductUseCase := applicationnegocio.NewCreateProduct(productRepo)
	getAllProductsUseCase := applicationnegocio.GetAllProducts(productRepo)
	getOneProductUseCase := applicationnegocio.GetOneProduct(productRepo)
	deleteProductUseCase := applicationnegocio.DeleteProduct(productRepo)
	putProductUseCase := applicationnegocio.PutProduct(productRepo)
	// Crear instancias de los controladores
	PostProductsHandler = controller.NewPostProductsHandler(createProductUseCase)
	GetProductsHandler = controller.NewGetProductsHandler(getAllProductsUseCase)
	GetOneProductHadler = controller.NewGetOneProductHandler(getOneProductUseCase)
	DeleteProductHadler = controller.NewDeleteProductHandler(deleteProductUseCase)
	PutProductHadler = controller.NewPutProductUseCase(putProductUseCase)
}
