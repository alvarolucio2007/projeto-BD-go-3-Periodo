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

func (h *HubConexoes) HandlerReadAllProva(c *gin.Context) {
	res, err := h.DoReadAllProva()
	if err != nil {
		SendError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Provas lidas com sucesso",
		"provas":  res,
	})
}

func (h *HubConexoes) HandlerReadProva(c *gin.Context) {
	nomeProva := c.Param("nome_prova")
	res, err := h.DoReadProva(nomeProva)
	if err != nil {
		SendError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Provas lidas com sucesso",
		"provas":  res,
	})
}
