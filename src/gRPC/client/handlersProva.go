package grpcclient

import (
	"net/http"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
	"github.com/gin-gonic/gin"
)

func (h *HubConexoes) HandlerCreateProva(c *gin.Context) {
	var novaProva models.Provas
	if err := c.ShouldBindJSON(&novaProva); err != nil {
		SendError(c, err)
		return
	}
	id, err := h.DoCreateProva(&novaProva)
	if err != nil {

		SendError(c, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Prova criada com sucesso",
		"id":      id,
	})
}
