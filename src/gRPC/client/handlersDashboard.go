package grpcclient

import (
	"context"
	"net/http"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/cache"
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/gRPC/proto"
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type DashboardHandler struct {
	Rdb             *redis.Client
	DashboardClient proto.DashboardServiceClient
}

func FromProto(p *proto.EstatisticaAluno) models.EstatisticaAluno {
	return models.EstatisticaAluno{
		QuantidadeProva: int(p.QuantidadeProva),
		MediaProvas:     float64(p.QuantidadeProva),
	}
}

func (d *DashboardHandler) HandlerQuantidadeProvaAluno(c *gin.Context, hub *HubGeral) {
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
			go func(client proto.DashboardServiceClient) {
				bgCtx := context.Background()
				res, err := client.QuantidadeProvaAluno(bgCtx, &proto.QuantidadeProvaAlunoRequest{NomeBusca: ""})
				if err == nil {
					_ = cache.AdicionarQuantidadeProvaAlunos(bgCtx, d.Rdb, res.Response)
				}
			}(hub.Dashboard)
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

func (d *DashboardHandler) HandlerQuantidadeNotaProvaAluno(c *gin.Context, hub *HubGeral) {
	username := c.Query("username")
	ctx := c.Request.Context()
	if username == "" {
		resRedis, err := cache.LerQuantidadeNotaProvaAlunos(ctx, d.Rdb)
		if err != nil {
			SendError(c, err)
			return
		}
		if resRedis != nil {
			c.JSON(http.StatusOK, resRedis)
			go func(client proto.DashboardServiceClient) {
				bgCtx := context.Background()
				res, err := client.QuantidadeNotaProvaAluno(bgCtx, &proto.QuantidadeNotaProvaAlunoRequest{NomeBusca: ""})
				if err == nil {
					mapCerto := make(map[string]models.EstatisticaAluno)
					for n, e := range res.Response {
						mapCerto[n] = FromProto(e)
					}
					_ = cache.AdicionarQuantidadeNotaProvaAlunos(bgCtx, d.Rdb, mapCerto)
				}
			}(hub.Dashboard)
			return
		}
	}
	res, err := hub.Dashboard.QuantidadeNotaProvaAluno(ctx, &proto.QuantidadeNotaProvaAlunoRequest{NomeBusca: username})
	if err != nil {
		SendError(c, err)
		return
	}

	mapCerto := make(map[string]models.EstatisticaAluno)
	for n, e := range res.Response {
		mapCerto[n] = FromProto(e)
	}
	c.JSON(http.StatusOK, mapCerto)
}
