package grpcclient

import (
	"net/http"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/gRPC/proto"
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type LeftJoinHandler struct {
	Rdb            *redis.Client
	LeftJoinClient *proto.LeftJoinServiceClient
}
type InnerJoinHandler struct {
	Rdb             *redis.Client
	InnerJoinClient *proto.InnerJoinServiceClient
}

func (l *LeftJoinHandler) HandlerLeftJoin(c *gin.Context, leftConn *LeftJoinConexao) {
	res, err := leftConn.DoLeftJoin()
	if err != nil {
		SendError(c, err)
		return
	}
	if res == nil {
		res = []*models.LeftJoinType{}
	}
	c.JSON(http.StatusOK, gin.H{
		"message":   "Left Join realizado com sucesso",
		"resultado": res,
	})
}

func (i *InnerJoinHandler) HandlerInnerJoin(c *gin.Context, innerConn *InnerJoinConexao) {
	res, err := innerConn.DoInnerJoin()
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
