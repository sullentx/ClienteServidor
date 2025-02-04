package server

import (
	"bytes"
	"encoding/json"
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
	ticker := time.NewTicker(1 * time.Minute)
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
		// handle error
		return
	}
	defer resp.Body.Close()

	var products []entities.Product
	if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
		// handle error
		return
	}

	// Send products to replica server
	jsonData, err := json.Marshal(products)
	if err != nil {
		// handle error
		return
	}

	_, err = http.Post(r.replicaServerURL+"/replicated-products", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		// handle error
		return
	}
}
