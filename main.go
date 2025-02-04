package main

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Quantity     int    `json:"quantity"`
	CodigoBarras string `json:"codigo_barras"`
}

var (
	products    []Product
	productsMux sync.Mutex
	updateChan  = make(chan struct{})
)

// Servidor Principal: CRUD de Productos
func mainServer() {
	r := gin.Default()

	r.POST("/products", func(c *gin.Context) {
		var product Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
			return
		}
		productsMux.Lock()
		product.ID = len(products) + 1
		products = append(products, product)
		productsMux.Unlock()
		updateChan <- struct{}{}
		c.JSON(http.StatusCreated, product)
	})

	r.GET("/products", func(c *gin.Context) {
		productsMux.Lock()
		c.JSON(http.StatusOK, products)
		productsMux.Unlock()
	})

	r.PUT("/products/:id", func(c *gin.Context) {
		var updatedProduct Product
		if err := c.ShouldBindJSON(&updatedProduct); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
			return
		}
		productsMux.Lock()
		for i, p := range products {
			if p.ID == updatedProduct.ID {
				products[i] = updatedProduct
				updateChan <- struct{}{}
				break
			}
		}
		productsMux.Unlock()
		c.Status(http.StatusOK)
	})

	r.DELETE("/products/:id", func(c *gin.Context) {
		var id int
		if err := c.ShouldBindJSON(&id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
			return
		}
		productsMux.Lock()
		for i, p := range products {
			if p.ID == id {
				products = append(products[:i], products[i+1:]...)
				updateChan <- struct{}{}
				break
			}
		}
		productsMux.Unlock()
		c.Status(http.StatusOK)
	})

	log.Println("Servidor Principal corriendo en :8081")
	r.Run(":8081")
}

// Servidor de Replicación: Long y Short Polling
func replicationServer() {
	r := gin.Default()

	r.GET("/replication", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		flusher, ok := c.Writer.(http.Flusher)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Long polling no compatible"})
			return
		}

		select {
		case <-updateChan:
			productsMux.Lock()
			c.JSON(http.StatusOK, products)
			productsMux.Unlock()
			flusher.Flush()
		case <-time.After(30 * time.Second):
			c.Status(http.StatusNoContent)
		}
	})

	log.Println("Servidor de Replicación corriendo en :8082")
	r.Run(":8082")
}

func main() {
	go mainServer()
	replicationServer()
}
