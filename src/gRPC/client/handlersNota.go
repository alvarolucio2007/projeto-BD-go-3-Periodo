package grpcclient

import (
	"net/http"
	"strconv"

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
	username := c.Query("username")
	res, err := h.DoReadNota(username)
	if err != nil {
		SendError(c, err)
		return
	}
	if res == nil {
		res = []*models.InnerJoinType{}
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
	c.JSON(http.StatusOK, gin.H{
		"message": "Nota atualizada com sucesso",
	})
}

func (h *HubConexoes) HandlerDeleteNota(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		SendError(c, err)
		return
	}
	err = h.DoDeleteNota(uint32(idUint))
	if err != nil {
		SendError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Nota deletada com sucesso",
	})
}
