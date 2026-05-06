package grpcclient

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/cache"
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/gRPC/proto"
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func SendError(c *gin.Context, err error) {
	status, resposta := ErrorHandler(err)
	c.JSON(status, resposta)
}

func ErrorHandler(err error) (int, gin.H) {
	if st, ok := status.FromError(err); ok {
		switch st.Code() {
		case codes.NotFound:
			return http.StatusNotFound, gin.H{"error": "Recurso não encontrado."}
		case codes.Unavailable:
			return http.StatusServiceUnavailable, gin.H{"error": "Servidor gRPC fora do ar."}
		case codes.InvalidArgument:
			return http.StatusBadRequest, gin.H{"error": st.Message()}
		}
	}
	return http.StatusInternalServerError, gin.H{"error": err.Error()}
}

type UsuarioHandler struct {
	Rdb        *redis.Client
	UserClient proto.UsuariosServiceClient
}

func (u *UsuarioHandler) HandlerAddUsuario(c *gin.Context, hub *HubGeral) {
	var novoUsuario models.Usuario
	if err := c.ShouldBindJSON(&novoUsuario); err != nil {
		SendError(c, err)
		return
	}
	id, err := hub.DoCreateUser(&novoUsuario)
	if err != nil {
		SendError(c, err)
		return
	}
	if err := cache.AdicionarUsuarioRedis(c, u.Rdb, id, &novoUsuario); err != nil {
		SendError(c, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuário criado com sucesso",
		"id":      id,
	})
}

func (u *UsuarioHandler) HandlerLerUsuario(c *gin.Context, hub *HubGeral) {
	usernameQuery := c.Query("username")
	ctx := c.Request.Context()
	usuarios, err := cache.LerTodosUsuariosRedis(ctx, u.Rdb)
	if errors.Is(err, redis.Nil) {
		usuarios, err = hub.DoReadAllUser(ctx)
		if err != nil {
			SendError(c, err)
			return
		}
		_ = cache.AdicionarTodosUsuariosRedis(ctx, u.Rdb, usuarios)
	} else if err != nil {
		SendError(c, err)
		return
	}
	if usernameQuery == "" {
		c.JSON(http.StatusOK, usuarios)
		return
	}
	var filtrados []*models.Usuario
	for _, user := range usuarios {
		if strings.Contains(strings.ToLower(user.Username), strings.ToLower(usernameQuery)) {
			filtrados = append(filtrados, user)
		}
	}
	c.JSON(http.StatusOK, filtrados)
}

func (u *UsuarioHandler) HandlerUpdateUsuario(c *gin.Context, hub *HubGeral) {
	var usuarioEdit models.Usuario
	if err := c.ShouldBindJSON(&usuarioEdit); err != nil {
		SendError(c, err)
		return
	}
	if usuarioEdit.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do usuário não enviada"})
		return
	}
	if err := hub.DoUpdateUser(&usuarioEdit); err != nil {
		SendError(c, err)
		return
	}
	if err := cache.DeletarUsuarioRedis(c, u.Rdb, usuarioEdit.ID); err != nil {
		SendError(c, err)
		return
	}
	if err := cache.DeletarTodosUsuariosRedis(c, u.Rdb); err != nil {
		SendError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"usuario": usuarioEdit,
	})
}

func (u *UsuarioHandler) HandlerDeleteUsuario(c *gin.Context, hub *HubGeral) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		SendError(c, err)
		return
	}
	err = hub.DoDeleteUser(uint32(idUint))
	if err != nil {
		SendError(c, err)
		return
	}
	err = cache.DeletarUsuarioRedis(c, u.Rdb, uint32(idUint))
	if err != nil {
		SendError(c, err)
		return
	}
	err = cache.DeletarTodosUsuariosRedis(c, u.Rdb)
	if err != nil {
		SendError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Usuário excluido com sucesso"})
}

func (u *UsuarioHandler) HandlerAuth(c *gin.Context, hub *HubGeral) {
	var credenciais struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&credenciais); err != nil {
		SendError(c, err)
		return
	}
	result, err := hub.DoAuth(credenciais.Username, credenciais.Password)
	if err != nil {
		SendError(c, err)
		return
	}
	if !result.Status {
		c.JSON(http.StatusUnauthorized, gin.H{"error": result.Mensagem})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Login realizado com sucesso",
		"auth":    true,
	})
}
