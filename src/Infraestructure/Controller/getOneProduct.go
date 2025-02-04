package controller

import (
	"context"
	applicationnegocio "main/src/Application-negocio"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type GetOneProductHandler struct {
	GetOneProductUseCase *applicationnegocio.GetOneProductUseCase
}

func NewGetOneProductHandler(getOneProductUseCase *applicationnegocio.GetOneProductUseCase) *GetOneProductHandler {
	return &GetOneProductHandler{GetOneProductUseCase: getOneProductUseCase}
}

func (h *GetOneProductHandler) HandleGetOne(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Debe de ser un número entero"})
		return
	}

	c.Header("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			// Si el tiempo expira, devuelve un timeout.
			c.JSON(http.StatusRequestTimeout, gin.H{"error": "Request timed out"})
			return
		case <-ticker.C:
			// Verifica si hay nuevos productos.
			product, err := h.GetOneProductUseCase.Execute(id)
			if err == nil {
				c.JSON(http.StatusOK, gin.H{"Producto": product})
				c.Writer.Flush()
				return
			}
		}
	}
}
