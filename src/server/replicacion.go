package server

import (
	domainnegocio "main/src/Domain-negocio"
	entities "main/src/Domain-negocio/Entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostReplicatedProductsHandler struct {
	productRepository domainnegocio.IproductrRepositoy
	updates           []entities.Product
}

func NewPostReplicatedProductsHandler(productRepository domainnegocio.IproductrRepositoy) *PostReplicatedProductsHandler {
	return &PostReplicatedProductsHandler{
		productRepository: productRepository,
		updates:           []entities.Product{},
	}
}

func (handle *PostReplicatedProductsHandler) Handle(g *gin.Context) {
	var products []entities.Product

	if err := g.ShouldBind(&products); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	for _, product := range products {
		if err := handle.productRepository.Save(product); err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}
		handle.updates = append(handle.updates, product)
	}

	g.JSON(http.StatusCreated, gin.H{"Message": "Products Replicated"})
}
