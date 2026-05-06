package grpcclient

import (
	"net/http"
	"strconv"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/cache"
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/gRPC/proto"
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type NotaHandler struct {
	Rdb        *redis.Client
	NotaClient proto.NotaServiceClient
}

func (n *NotaHandler) HandlerAddNota(c *gin.Context, hub *HubGeral) {
	var novaNota models.Notas
	if err := c.ShouldBindJSON(&novaNota); err != nil {
		SendError(c, err)
		return
	}
	id, err := hub.DoCreateNota(&novaNota)
	if err != nil {
		SendError(c, err)
		return
	}
	if err := cache.AdicionarNotaRedis(c, n.Rdb, id, &novaNota); err != nil {
		SendError(c,err)
		return
	}
	listaTodasNotas,err:=hub.DoReadNota(username string)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Nota criada com sucesso",
		"id":      id,
	})
}

func (n *NotaHandler) HandlerReadNota(c *gin.Context, hub *HubGeral) {
	username := c.Query("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username é obrigatório para a pesquisa."})
		return
	}
	res, err := hub.DoReadNota(username)
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

func (n *NotaHandler) HandlerUpdateNota(c *gin.Context, hub *HubGeral) {
	var novaNota models.Notas
	if err := c.ShouldBindJSON(&novaNota); err != nil {
		SendError(c, err)
		return
	}
	err := hub.DoUpdateNota(&novaNota)
	if err != nil {
		SendError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Nota atualizada com sucesso",
	})
}

func (n *NotaHandler) HandlerDeleteNota(c *gin.Context, hub *HubGeral) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		SendError(c, err)
		return
	}
	err = hub.DoDeleteNota(uint32(idUint))
	if err != nil {
		SendError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Nota deletada com sucesso",
	})
}
