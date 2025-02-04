package controller

import (
	applicationnegocio "main/src/Application-negocio"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteProductHandler struct {
	DeleteProductUseCase *applicationnegocio.DeleteProductUseCase
}

func NewDeleteProductHandler(deleteProductUseCase *applicationnegocio.DeleteProductUseCase) *DeleteProductHandler {
	return &DeleteProductHandler{DeleteProductUseCase: deleteProductUseCase}
}

func (h *DeleteProductHandler) HandleDelete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Debe de ser un numero entero"})
		return
	}

	err = h.DeleteProductUseCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Producto eliminado"})
}
