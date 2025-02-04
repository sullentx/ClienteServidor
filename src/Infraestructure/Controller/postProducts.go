package controller

import (
	applicationnegocio "main/src/Application-negocio"
	entities "main/src/Domain-negocio/Entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostProductsHandler struct {
	createProductUseCase *applicationnegocio.CreateProductUseCase
}

// constructor que devuelve un puntero a la estructura PostProductsHandler
func NewPostProductsHandler(createProductUseCase *applicationnegocio.CreateProductUseCase) *PostProductsHandler {
	return &PostProductsHandler{createProductUseCase: createProductUseCase}
}

func (handle *PostProductsHandler) Handle(g *gin.Context) {
	var product entities.Product //se crea una variable product de tipo Product

	if err := g.ShouldBind(&product); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	//se ejecuta el metodo Execute de la estructura CreateProductUseCase
	if err := handle.createProductUseCase.Execute(product); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	g.JSON(http.StatusCreated, gin.H{"Message": "Product Creado"})
}
