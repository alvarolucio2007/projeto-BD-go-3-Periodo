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

type UsuarioHandler struct {
	Rdb        *redis.Client
	UserClient proto.UsuariosServiceClient
}

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

func (u *UsuarioHandler) HandlerAddUsuario(c *gin.Context, userConn *UserConexao) {
	var novoUsuario models.Usuario
	if err := c.ShouldBindJSON(&novoUsuario); err != nil {
		SendError(c, err)
		return
	}
	id, err := userConn.DoCreateUser(&novoUsuario)
	if err != nil {
		SendError(c, err)
		return
	}
	if err := cache.AdicionarUsuarioRedis(c, u.Rdb, id, &novoUsuario); err != nil {
		SendError(c, err)
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuário criado com sucesso",
		"id":      id,
	})
}

func (u *UsuarioHandler) HandlerLerUsuario(c *gin.Context, userConn *UserConexao) {
	username := c.Query("username")
	userEncontradoRedis, err := cache.LerTodosUsuariosRedis(c, u.Rdb)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			result := make([]*models.Usuario, 0, len(userEncontradoRedis))
			for _, u := range userEncontradoRedis {
				if strings.Contains(username, u.Username) {
					result = append(result, u)
				}
			}
			c.JSON(http.StatusOK, result)
			return
		}
		SendError(c, err)
		return
	}
	result, err := userConn.DoReadUser(username)
	if err != nil {
		SendError(c, err)
		return
	}
	if err := cache.AdicionarTodosUsuariosRedis(c, u.Rdb, result); err != nil {
		SendError(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (u *UsuarioHandler) HandlerUpdateUsuario(c *gin.Context, userConn *UserConexao) {
	var usuarioEdit models.Usuario
	if err := c.ShouldBindJSON(&usuarioEdit); err != nil {
		SendError(c, err)
		return
	}
	if usuarioEdit.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do usuário não enviada"})
	}
	if err := cache.DeletarUsuarioRedis(c, u.Rdb, usuarioEdit.ID); err != nil {
		if errors.Is(err, redis.Nil) {
			err = cache.AdicionarUsuarioRedis(c, u.Rdb, usuarioEdit.ID, &usuarioEdit)
			if err != nil {
				SendError(c, err)
				return
			}
		}
	}
	if err := userConn.DoUpdateUser(&usuarioEdit); err != nil {
		SendError(c, err)
		return
	}
	if err := cache.AdicionarUsuarioRedis(c, u.Rdb, usuarioEdit.ID, &usuarioEdit); err != nil {
		SendError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"usuario": usuarioEdit,
	})
}

func (u *UsuarioHandler) HandlerDeleteUsuario(c *gin.Context, userConn *UserConexao) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		SendError(c, err)
		return
	}
	err = userConn.DoDeleteUser(uint32(idUint))
	if err != nil {
		SendError(c, err)
		return
	}
	err = cache.DeletarUsuarioRedis(c, u.Rdb, uint32(idUint))
	if err != nil {
		SendError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Usuário excluido com sucesso"})
}

func (u *UsuarioHandler) HandlerAuth(c *gin.Context, userConn *UserConexao) {
	var credenciais struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&credenciais); err != nil {
		SendError(c, err)
		return
	}
	result, err := userConn.DoAuth(credenciais.Username, credenciais.Password)
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
