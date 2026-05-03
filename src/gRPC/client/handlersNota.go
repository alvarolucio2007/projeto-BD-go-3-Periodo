package grpcclient

import (
	"net/http"
	"strconv"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/gRPC/proto"
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis"
)

type NotaHandler struct {
	Rdb        *redis.Client
	NotaClient *proto.NotaServiceClient
}

func (n *NotaHandler) HandlerAddNota(c *gin.Context, notaConn *NotaConexao) {
	var novaNota models.Notas
	if err := c.ShouldBindJSON(&novaNota); err != nil {
		SendError(c, err)
		return
	}
	id, err := notaConn.DoCreateNota(&novaNota)
	if err != nil {
		SendError(c, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Nota criada com sucesso",
		"id":      id,
	})
}

func (n *NotaHandler) HandlerReadNota(c *gin.Context, notaConn *NotaConexao) {
	username := c.Query("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username é obrigatório para a pesquisa."})
		return
	}
	res, err := notaConn.DoReadNota(username)
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

func (n *NotaHandler) HandlerUpdateNota(c *gin.Context, notaConn *NotaConexao) {
	var novaNota models.Notas
	if err := c.ShouldBindJSON(&novaNota); err != nil {
		SendError(c, err)
		return
	}
	err := notaConn.DoUpdateNota(&novaNota)
	if err != nil {
		SendError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Nota atualizada com sucesso",
	})
}

func (n *NotaHandler) HandlerDeleteNota(c *gin.Context, notaConn *NotaConexao) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		SendError(c, err)
		return
	}
	err = notaConn.DoDeleteNota(uint32(idUint))
	if err != nil {
		SendError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Nota deletada com sucesso",
	})
}
