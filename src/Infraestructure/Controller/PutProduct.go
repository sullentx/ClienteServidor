package controller

import (
	applicationnegocio "main/src/Application-negocio"
	entities "main/src/Domain-negocio/Entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PutProductHandler struct {
	PutProductUseCase *applicationnegocio.PutProductUseCase
}

//constructor

func NewPutProductUseCase(putProductUseCase *applicationnegocio.PutProductUseCase) *PutProductHandler {
	return &PutProductHandler{PutProductUseCase: putProductUseCase}
}

func (h *PutProductHandler) HandlePut(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	var product entities.Product

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "debe de ser numero entero"})
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "datos invalidos"})
		return
	}

	err = h.PutProductUseCase.Execute(id, product)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"message": id})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Producto actualizado"})
}
