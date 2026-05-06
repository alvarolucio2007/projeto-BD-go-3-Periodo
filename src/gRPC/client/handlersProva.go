package grpcclient

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/cache"
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/gRPC/proto"
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type ProvaHandler struct {
	Rdb         *redis.Client
	ProvaClient proto.ProvaServiceClient
}

func (p *ProvaHandler) HandlerCreateProva(c *gin.Context, hub *HubGeral) {
	var novaProva models.Provas
	if err := c.ShouldBindJSON(&novaProva); err != nil {
		SendError(c, err)
		return
	}
	id, err := hub.DoCreateProva(&novaProva)
	if err != nil {
		SendError(c, err)
		return
	}
	if err := cache.AdicionarTestRedis(c.Request.Context(), p.Rdb, id, &novaProva); err != nil {
		SendError(c, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Prova criada com sucesso",
		"id":      id,
	})
}

func (p *ProvaHandler) HandlerReadAllProva(c *gin.Context, hub *HubGeral) {
	res, err := cache.LerAllTestRedis(c.Request.Context(), p.Rdb)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Provas lidas com sucesso",
			"provas":  res,
		})
		return
	}
	if errors.Is(err, redis.Nil) {
		res, err = hub.DoReadAllProva()
		if err != nil {
			SendError(c, err)
			return
		}
		if res == nil {
			res = []*models.Provas{}
		}
		err = cache.AdicionarTodosTestRedis(c, p.Rdb, res)
		if err != nil {
			SendError(c, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Provas lidas com sucesso",
			"provas":  res,
		})
	} else {
		SendError(c, err)
		return
	}
}

func (p *ProvaHandler) HandlerUpdateProva(c *gin.Context, hub *HubGeral) {
	var novaProva models.Provas
	if err := c.ShouldBindJSON(&novaProva); err != nil {
		SendError(c, err)
		return
	}
	if novaProva.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID da prova é obrigratório"})
		return
	}
	if err := hub.DoUpdateProva(&novaProva); err != nil {
		SendError(c, err)
		return
	}
	if err := cache.AdicionarTestRedis(c, p.Rdb, novaProva.ID, &novaProva); err != nil {
		SendError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Prova atualizada com sucesso",
	})
}

func (p *ProvaHandler) HandlerDeleteProva(c *gin.Context, hub *HubGeral) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		SendError(c, err)
		return
	}
	err = hub.DoDeleteProva(uint32(idUint))
	if err != nil {
		SendError(c, err)
		return
	}
	if err := cache.DeletarTestRedis(c, p.Rdb, uint32(idUint)); err != nil {
		SendError(c, err)
		return
	}
	provas, err := hub.DoReadAllProva()
	if err != nil {
		SendError(c, err)
		return
	}
	if err := cache.AdicionarTodosTestRedis(c, p.Rdb, provas); err != nil {
		SendError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Prova deletada com sucesso",
	})
}
