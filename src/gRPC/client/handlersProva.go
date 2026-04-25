package grpcclient

import (
	"fmt"
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
	}
	fmt.Printf("Tentando salvar no banco: %+v\n", novaProva)
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
	fmt.Printf("DEBUG: Total de provas recuperadas: %d\n", len(res))
	if len(res) > 0 {
		fmt.Printf("DEBUG: Primeira prova: %+v\n", res[0])
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Provas lidas com sucesso",
		"provas":  res, // O Gin usará as tags `json:"..."` da sua struct models.Provas
	})
}

func (h *HubConexoes) HandlerReadAllProvas(c *gin.Context) {
	// Sem Query, sem JSON, sem barreiras. Chamada direta.
	res, err := h.DoReadAllProva() // Garanta que essa função chame o gRPC ReadAll
	if err != nil {
		SendError(c, err)
		return
	}

	// Proteção para o Python não receber 'null'
	if res == nil {
		c.JSON(http.StatusOK, gin.H{"provas": []any{}})
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
		return
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
