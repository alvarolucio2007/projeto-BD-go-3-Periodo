// Package api cuida da API RESTful externa, para o HTMX.
package api

import (
	grpcclient "github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/gRPC/client"
	"github.com/gin-gonic/gin"
)

const addr string = "localhost:50051"

func SetupExtRoutes() error {
	hub, err := grpcclient.ConnectAll(addr)
	if err != nil {
		return err
	}
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.POST("/usuario", hub.HandlerAddUsuario)
	r.POST("/usuario/ler", hub.HandlerLerUsuario)
	r.PUT("/usuario", hub.HandlerUpdateUsuario)
	r.DELETE("/usuario/:id", hub.HandlerDeleteUsuario)
	r.POST("/usuario/auth", hub.HandlerAuth)
}
