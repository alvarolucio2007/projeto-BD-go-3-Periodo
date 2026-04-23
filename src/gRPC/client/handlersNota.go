package grpcclient

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
	"github.com/gin-gonic/gin"
)

func (h *HubConexoes) HandlerAddNota(c *gin.Context) {
	var novaNota models.Notas
	notaNumeroString := c.PostForm("nota_prova")
	notaProva, err := strconv.ParseFloat(notaNumeroString, 64)
	if err != nil {
		SendError(c, err)
		return
	}
	novaNota.NotaProva = float32(notaProva)
	pID, errP := strconv.ParseUint(c.PostForm("id_prova"), 10, 32)
	uID, errU := strconv.ParseUint(c.PostForm("id_usuario"), 10, 32)
	if errP != nil || errU != nil {
		SendError(c, fmt.Errorf("id de prova ou usuário inválido"))
		return
	}
	novaNota.ProvaID = uint32(pID)
	novaNota.UsuarioID = uint32(uID)

	id, err := h.DoCreateNota(&novaNota)
	if err != nil {
		SendError(c, err)
		return
	}
	c.HTML(http.StatusCreated, "add_nota.html", gin.H{
		"message": "Nota criada com sucesso",
		"id":      id,
	})
}

func (h *HubConexoes) HandlerReadNota(c *gin.Context) {
	username := c.PostForm("username")
	res, err := h.DoReadNota(username)
	if err != nil {
		SendError(c, err)
		return
	}
	if len(res) == 0 {
		c.HTML(http.StatusNotFound, "read_notas.html", gin.H{
			"message": "não foram encontradas notas",
			"notas":   res,
		})
		return
	}
	c.HTML(http.StatusOK, "read_notas.html", gin.H{
		"message": "Busca de nota feita com sucesso",
		"notas":   res,
	})
}

func (h *HubConexoes) HandlerUpdateNota(c *gin.Context) {
	var novaNota models.Notas
	if err := c.ShouldBind(&novaNota); err != nil {
		SendError(c, err)
		return
	}
	err := h.DoUpdateNota(&novaNota)
	if err != nil {
		SendError(c, err)
		return
	}
	c.HTML(http.StatusOK, "nota_linha", gin.H{
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
	c.HTML(http.StatusOK, "delete_nota", gin.H{
		"message": "Nota deletada com sucesso",
	})
}
