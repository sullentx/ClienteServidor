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
	router := gin.Default()
	// Inicializar dependencias
	core.InitPostgres()
	infraestructure.Init()

	routes.SetRoutes(router, infraestructure.PostProductsHandler, infraestructure.GetProductsHandler,
		infraestructure.GetOneProductHadler, infraestructure.DeleteProductHadler, infraestructure.PutProductHadler)
	// Iniciar el servidor

	log.Println("Server started at :8080")
	postReplicatedProductsHandler := server.NewPostReplicatedProductsHandler(nil)
	replica := gin.Default()
	replica.POST("/replicated-products", postReplicatedProductsHandler.Handle)
	replica.GET("/updates", server.NewGetUpdatesHandler(postReplicatedProductsHandler).Handle)

	// Iniciar el servidor replicador
	go func() {
		log.Println("Replica server started at :8081")
		log.Fatal(replica.Run(":8081"))
	}()

	select {}
}
