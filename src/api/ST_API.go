// Package api cuida da API RESTful externa, para o HTMX.
package api

import (
	grpcclient "github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/gRPC/client"
	"github.com/gin-gonic/gin"
)

func SetupExtRoutes(hub *grpcclient.HubConexoes) error {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	// APIs de usuário
	r.POST("/usuario", hub.HandlerAddUsuario)
	r.GET("/usuario/buscar", hub.HandlerLerUsuario)
	r.PUT("/usuario", hub.HandlerUpdateUsuario)
	r.DELETE("/usuario/:id", hub.HandlerDeleteUsuario)
	r.POST("/usuario/auth", hub.HandlerAuth)
	// APIs de prova
	r.POST("/provas", hub.HandlerCreateProva)
	r.GET("/provas", hub.HandlerReadAllProva)
	r.PUT("/provas", hub.HandlerUpdateProva)
	r.DELETE("/provas/:id", hub.HandlerDeleteProva)
	// APIs de nota
	r.POST("/notas", hub.HandlerAddNota)
	r.GET("/notas/buscar", hub.HandlerReadNota)
	r.PUT("/notas", hub.HandlerUpdateNota)
	r.DELETE("/notas/:id", hub.HandlerDeleteNota)
	// APIs de JOIN
	r.GET("/left_join", hub.HandlerLeftJoin)
	r.GET("/inner_join", hub.HandlerInnerJoin)
	err := r.Run(":8080")
	if err != nil {
		return err
	}
	return nil
}
