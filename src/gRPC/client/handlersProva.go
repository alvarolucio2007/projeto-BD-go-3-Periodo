package grpcclient

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
	"github.com/gin-gonic/gin"
)

func (h *HubConexoes) HandlerCreateProva(c *gin.Context) {
	var novaProva models.Provas
	novaProva.NomeProva = c.PostForm("nome_prova")
	novaProva.MateriaProva = c.PostForm("materia_prova")
	novaProva.TurmaProva = c.PostForm("turma_prova")
	dataString := c.PostForm("data_prova")
	dataProva, err := time.Parse("2006-01-02", dataString)
	if err != nil {
		SendError(c, fmt.Errorf("formato de data inválido: %v", err))
		return
	}
	novaProva.DataProva = dataProva
	id, err := h.DoCreateProva(&novaProva)
	if err != nil {

		SendError(c, err)
		return
	}
	c.HTML(http.StatusCreated, "add_prova.html", gin.H{
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
	c.HTML(http.StatusOK, "read_all_prova.html", gin.H{
		"message": "Provas lidas com sucesso",
		"provas":  res,
	})
}

func (h *HubConexoes) HandlerReadProva(c *gin.Context) {
	nomeProva := c.Query("nome_prova")
	res, err := h.DoReadProva(nomeProva)
	if err != nil {
		SendError(c, err)
		return
	}
	c.HTML(http.StatusOK, "read_prova.html", gin.H{
		"provas": res,
	})
}

func (h *HubConexoes) HandlerUpdateProva(c *gin.Context) {
	var novaProva models.Provas
	if err := c.ShouldBind(&novaProva); err != nil {
		SendError(c, err)
		return
	}
	if err := h.DoUpdateProva(&novaProva); err != nil {
		SendError(c, err)
		return
	}
	c.HTML(http.StatusOK, "prova_linha", gin.H{
		"message": "Prova atualizada com sucesso",
	})
}

func (h *HubConexoes) HandlerDeleteProva(c *gin.Context) {
	id := c.PostForm("id")
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
	c.HTML(http.StatusOK, "delete_prova", gin.H{
		"message": "Prova deletada com sucesso",
	})
}
