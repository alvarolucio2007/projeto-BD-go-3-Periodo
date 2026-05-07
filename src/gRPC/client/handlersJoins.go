package grpcclient

import (
	"errors"
	"net/http"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/cache"
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/gRPC/proto"
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type LeftJoinHandler struct {
	Rdb            *redis.Client
	LeftJoinClient proto.LeftJoinServiceClient
}
type InnerJoinHandler struct {
	Rdb             *redis.Client
	InnerJoinClient proto.InnerJoinServiceClient
}

func (l *LeftJoinHandler) HandlerLeftJoin(c *gin.Context, hub *HubGeral) {
	resRedis, err := cache.LerLeftJoinRedis(c, l.Rdb)
	if err == nil && resRedis != nil {
		c.JSON(http.StatusOK, gin.H{
			"message":   "Left Join realizado com sucesso",
			"resultado": resRedis,
		})
		return
	}
	if err != nil && !errors.Is(err, redis.Nil) {
		SendError(c, err)
		return
	}
	res, err := hub.DoLeftJoin()
	if err != nil {
		SendError(c, err)
		return
	}
	if res == nil {
		res = []*models.LeftJoinType{}
	}
	err = cache.AdicionarLeftJoinRedis(c, l.Rdb, res)
	if err != nil {
		SendError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":   "Left Join realizado com sucesso",
		"resultado": res,
	})
}

func (i *InnerJoinHandler) HandlerInnerJoin(c *gin.Context, hub *HubGeral) {
	resRedis, err := cache.LerInnerJoinRedis(c, i.Rdb)
	if err == nil && resRedis != nil {
		c.JSON(http.StatusOK, gin.H{
			"message":   "Inner Join realizado com sucesso",
			"resultado": resRedis,
		})
	}

	if err != nil && !errors.Is(err, redis.Nil) {
		SendError(c, err)
		return
	}
	res, err := hub.DoInnerJoin()
	if err != nil {
		SendError(c, err)
		return
	}
	if res == nil {
		res = []*models.InnerJoinType{}
	}
	c.JSON(http.StatusOK, gin.H{
		"message":   "Inner Join realizado com sucesso",
		"resultado": res,
	})
}
