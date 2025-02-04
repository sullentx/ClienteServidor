package routes

//mapper
import (
	controller "main/src/Infraestructure/Controller"

	gin "github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.Engine, postController *controller.PostProductsHandler, getController *controller.GetProductsHandler, getOneController *controller.GetOneProductHandler, deleteProduct *controller.DeleteProductHandler, putController *controller.PutProductHandler) {
	router.GET("/products", getController.Handle)
	router.POST("/products", postController.Handle)
	router.GET("/products/:id", getOneController.HandleGetOne)
	router.DELETE("/products/:id", deleteProduct.HandleDelete)
	router.PUT("/products/:id", putController.HandlePut)
}
