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

func (h *HubConexoes) HandlerReadNota(c *gin.Context) {
	username := c.Param("username")
	res, err := h.DoReadNota(username)
	if err != nil {
		SendError(c, err)
		return
	}
	if len(res) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "não foram encontradas notas",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Busca de nota feita com sucesso",
		"notas":   res,
	})
}

func (h *HubConexoes) HandlerUpdateNota(c *gin.Context) {
	var novaNota models.Notas
	if err := c.ShouldBindJSON(&novaNota); err != nil {
		SendError(c, err)
		return
	}
	err := h.DoUpdateNota(&novaNota)
	if err != nil {
		SendError(c, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Nota criada com sucesso",
	})
}
