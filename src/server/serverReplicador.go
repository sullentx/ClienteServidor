package server

import (
	"encoding/json"
	"fmt"
	"io"
	entities "main/src/Domain-negocio/Entities"
	"net/http"
	"time"
)

type Replicator struct {
	primaryServerURL string
	replicaServerURL string
}

func NewReplicator(primaryServerURL, replicaServerURL string) *Replicator {
	return &Replicator{
		primaryServerURL: primaryServerURL,
		replicaServerURL: replicaServerURL,
	}
}

func (r *Replicator) Start() {
	ticker := time.NewTicker(10 * time.Second) // Short polling every 10 seconds
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			r.replicate()
		}
	}
}
func (r *Replicator) replicate() {
	resp, err := http.Get(r.primaryServerURL + "/products")
	if err != nil {
		fmt.Println("❌ Error al hacer GET a /products:", err)
		return
	}
	defer resp.Body.Close()

	// 🔍 Imprime la respuesta antes de intentar decodificarla
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("📥 Respuesta del servidor principal:", string(body))

	var products []entities.Product
	if err := json.Unmarshal(body, &products); err != nil {
		fmt.Println("❌ Error al decodificar JSON:", err)
		return
	}

	fmt.Printf("🔄 Productos recibidos en el replicador: %v\n", products)
}

// Continuar replicación...
