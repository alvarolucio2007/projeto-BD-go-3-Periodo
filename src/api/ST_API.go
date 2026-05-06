// Package api cuida da API RESTful externa, para o HTMX.
package api

import (
	grpcclient "github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/gRPC/client"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func SetupExtRoutes(hub *grpcclient.HubGeral, rdb *redis.Client) error {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	// APIs de usuário
	userHandler := &grpcclient.UsuarioHandler{
		Rdb:        rdb,
		UserClient: hub.User,
	}
	r.POST("/usuario", func(c *gin.Context) { userHandler.HandlerAddUsuario(c, hub) })
	r.GET("/usuario/buscar", func(c *gin.Context) { userHandler.HandlerLerUsuario(c, hub) })
	r.PUT("/usuario", func(c *gin.Context) { userHandler.HandlerUpdateUsuario(c, hub) })
	r.DELETE("/usuario/:id", func(c *gin.Context) { userHandler.HandlerDeleteUsuario(c, hub) })
	r.POST("/usuario/auth", func(c *gin.Context) { userHandler.HandlerAuth(c, hub) })
	// APIs de prova
	provaHandler := &grpcclient.ProvaHandler{
		Rdb:         rdb,
		ProvaClient: hub.Prova,
	}
	r.POST("/provas", func(c *gin.Context) { provaHandler.HandlerCreateProva(c, hub) })
	r.GET("/provas", func(c *gin.Context) { provaHandler.HandlerReadAllProva(c, hub) })
	r.PUT("/provas", func(c *gin.Context) { provaHandler.HandlerUpdateProva(c, hub) })
	r.DELETE("/provas/:id", func(c *gin.Context) { provaHandler.HandlerDeleteProva(c, hub) })
	// APIs de nota
	notaHandler := &grpcclient.NotaHandler{
		Rdb:        rdb,
		NotaClient: hub.Nota,
	}
	r.POST("/notas", func(c *gin.Context) { notaHandler.HandlerAddNota(c, hub) })
	r.GET("/notas/buscar", func(c *gin.Context) { notaHandler.HandlerReadNota(c, hub) })
	r.PUT("/notas", func(c *gin.Context) { notaHandler.HandlerUpdateNota(c, hub) })
	r.DELETE("/notas/:id", func(c *gin.Context) { notaHandler.HandlerDeleteNota(c, hub) })
	// APIs de JOIN
	leftJoinHandler := &grpcclient.LeftJoinHandler{
		Rdb:            rdb,
		LeftJoinClient: hub.LeftJoin,
	}
	innerJoinHandler := &grpcclient.InnerJoinHandler{
		Rdb:             rdb,
		InnerJoinClient: hub.InnerJoin,
	}
	r.GET("/left_join", func(c *gin.Context) { leftJoinHandler.HandlerLeftJoin(c, hub) })
	r.GET("/inner_join", func(c *gin.Context) { innerJoinHandler.HandlerInnerJoin(c, hub) })
	err := r.Run(":8080")
	if err != nil {
		return err
	}
	return nil
}
