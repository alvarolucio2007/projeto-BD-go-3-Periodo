package grpcclient

import (
	"net/http"
	"strconv"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
	"github.com/gin-gonic/gin"
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

func (h *HubConexoes) HandlerAddUsuario(c *gin.Context) {
	var novoUsuario models.Usuario
	if err := c.ShouldBindJSON(&novoUsuario); err != nil {
		SendError(c, err)
		return
	}
	id, err := h.DoCreateUser(&novoUsuario)
	if err != nil {
		SendError(c, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuário criado com sucesso",
		"id":      id,
	})
}

func (h *HubConexoes) HandlerLerUsuario(c *gin.Context) {
	username := c.Query("username")
	result, err := h.DoReadUser(username)
	if err != nil {
		SendError(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *HubConexoes) HandlerUpdateUsuario(c *gin.Context) {
	var usuarioEdit models.Usuario
	if err := c.ShouldBindJSON(&usuarioEdit); err != nil {
		SendError(c, err)
		return
	}
	if usuarioEdit.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do usuário não enviada"})
	}
	if err := h.DoUpdateUser(&usuarioEdit); err != nil {
		SendError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"usuario": usuarioEdit,
	})
}

func (h *HubConexoes) HandlerDeleteUsuario(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		SendError(c, err)
		return
	}
	err = h.DoDeleteUser(uint32(idUint))
	if err != nil {
		SendError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Usuário excluido com sucesso"})
}

func (h *HubConexoes) HandlerAuth(c *gin.Context) {
	var credenciais struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&credenciais); err != nil {
		SendError(c, err)
		return
	}
	result, err := h.DoAuth(credenciais.Username, credenciais.Password)
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
