package grpcclient

import (
	"net/http"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
	"github.com/gin-gonic/gin"
)

func (h *HubConexoes) HandlerAddNota(c *gin.Context) {
	var novaNota models.Notas
	if err := c.ShouldBindJSON(&novaNota); err != nil {
		SendError(c, err)
		return
	}
	id, err := h.DoCreateNota(&novaNota)
	if err != nil {
		SendError(c, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Nota criada com sucesso",
		"id":      id,
	})
}
