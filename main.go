package main

import (
	"log"
	infraestructure "main/src/Infraestructure"
	routes "main/src/Infraestructure/Routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// Inicializar dependencias
	infraestructure.Init()

	routes.SetRoutes(router, infraestructure.PostProductsHandler, infraestructure.GetProductsHandler,
		infraestructure.GetOneProductHadler, infraestructure.DeleteProductHadler, infraestructure.PutProductHadler)
	// Iniciar el servidor

	log.Println("Server started at :8080")
	log.Fatal(router.Run())
}
