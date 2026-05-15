package grpcclient

import (
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/gRPC/proto"
	"github.com/redis/go-redis/v9"
)

type DashboardHandler struct {
	Rdb             *redis.Client
	DashboardClient proto.DashboardServiceClient
}
