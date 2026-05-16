package grpcclient

import (
	"context"
	"net/http"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/cache"
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/gRPC/proto"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type DashboardHandler struct {
	Rdb             *redis.Client
	DashboardClient proto.DashboardServiceClient
}

func (d *DashboardHandler) HandlerLerQuantidadeProvaAluno(c *gin.Context, hub *HubGeral) {
	username := c.Query("username")
	ctx := c.Request.Context()
	if username == "" {
		resRedis, err := cache.LerQuantidadeProvaAlunos(ctx, d.Rdb)
		if err != nil {
			SendError(c, err)
			return
		}
		if resRedis != nil {
			c.JSON(http.StatusOK, resRedis)
			go func() {
				bgCtx := context.Background()
				res, err := hub.Dashboard.QuantidadeProvaAluno(bgCtx, &proto.QuantidadeProvaAlunoRequest{NomeBusca: ""})
				if err == nil {
					_ = cache.AdicionarQuantidadeProvaAlunos(bgCtx, d.Rdb, res.Response)
				}
			}()
			return
		}
	}
	res, err := hub.Dashboard.QuantidadeProvaAluno(ctx, &proto.QuantidadeProvaAlunoRequest{NomeBusca: username})
	if err != nil {
		SendError(c, err)
		return
	}
	c.JSON(http.StatusOK, res.Response)
}
