package main

import (
	"log"
	core "main/src/Core"
	infraestructure "main/src/Infraestructure"
	routes "main/src/Infraestructure/Routes"
	"main/src/server"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar dependencias
	core.InitPostgres()
	infraestructure.Init()

	// Configuración del servidor principal
	router := gin.Default()
	routes.SetRoutes(router, infraestructure.PostProductsHandler, infraestructure.GetProductsHandler,
		infraestructure.GetOneProductHadler, infraestructure.DeleteProductHadler, infraestructure.PutProductHadler)

	// Iniciar el servidor principal
	go func() {
		log.Println("Server started at :8080")
		log.Fatal(router.Run(":8080"))
	}()

	// Configuración del servidor replicador
	postReplicatedProductsHandler := server.NewPostReplicatedProductsHandler(nil)
	replica := gin.Default()
	replica.POST("/replicated-products", postReplicatedProductsHandler.Handle)
	replica.GET("/updates", server.NewGetUpdatesHandler(postReplicatedProductsHandler).Handle)

	// Iniciar el servidor replicador
	go func() {
		log.Println("Replica server started at :8081")
		log.Fatal(replica.Run(":8081"))
	}()

	// Iniciar el replicador
	rep := server.NewReplicator("http://localhost:8080", "http://localhost:8081")
	go rep.Start()

	select {} // Mantener la función principal en ejecución
}
