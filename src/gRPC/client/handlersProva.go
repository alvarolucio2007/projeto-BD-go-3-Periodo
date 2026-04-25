package grpcclient

import (
	"net/http"
	"strconv"

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
	if res == nil {
		res = []*models.Provas{}
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Provas lidas com sucesso",
		"provas":  res,
	})
}

func (h *HubConexoes) HandlerReadProva(c *gin.Context) {
	var filtroProva struct {
		NomeProva string `json:"nome_prova"`
	}
	if err := c.ShouldBindJSON(&filtroProva); err != nil {
		filtroProva.NomeProva = c.Query("nome_prova")
	}
	res, err := h.DoReadProva(filtroProva.NomeProva)
	if err != nil {
		SendError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *HubConexoes) HandlerUpdateProva(c *gin.Context) {
	var novaProva models.Provas
	if err := c.ShouldBindJSON(&novaProva); err != nil {
		SendError(c, err)
		return
	}
	if novaProva.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID da prova é obrigratório"})
	}
	if err := h.DoUpdateProva(&novaProva); err != nil {
		SendError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Prova atualizada com sucesso",
	})
}

func (h *HubConexoes) HandlerDeleteProva(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		SendError(c, err)
		return
	}
	err = h.DoDeleteProva(uint32(idUint))
	if err != nil {
		SendError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Prova deletada com sucesso",
	})
}
