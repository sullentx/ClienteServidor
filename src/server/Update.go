package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetUpdatesHandler struct {
	postReplicatedProductsHandler *PostReplicatedProductsHandler
}

func NewGetUpdatesHandler(postReplicatedProductsHandler *PostReplicatedProductsHandler) *GetUpdatesHandler {
	return &GetUpdatesHandler{postReplicatedProductsHandler: postReplicatedProductsHandler}
}

func (handle *GetUpdatesHandler) Handle(g *gin.Context) {
	g.JSON(http.StatusOK, handle.postReplicatedProductsHandler.updates)
}
