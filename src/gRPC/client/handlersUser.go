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

	msg := "Erro desconhecido"
	if val, existe := resposta["error"]; existe {
		msg = val.(string)
	}

	if c.GetHeader("HX-Request") != "" {
		c.String(status, msg)
		return
	}

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
	novoUsuario.Username = c.PostForm("username")
	novoUsuario.Password = c.PostForm("password")
	novoUsuario.Role = c.PostForm("role")
	id, err := h.DoCreateUser(&novoUsuario)
	if err != nil {
		SendError(c, err)
		return
	}
	c.HTML(http.StatusCreated, "add_usuario.html", gin.H{
		"message": "Usuário criado com sucesso",
		"id":      id,
	})
}

func (h *HubConexoes) HandlerLerUsuario(c *gin.Context) {
	username := c.PostForm("username")
	result, err := h.DoReadUser(username)
	if err != nil {
		SendError(c, err)
		return
	}
	if len(result) == 0 {
		c.HTML(http.StatusCreated, "add_usuario.html", gin.H{
			"message":  "Usuários buscados com sucesso",
			"usuarios": result,
		})
	}
}

func (h *HubConexoes) HandlerUpdateUsuario(c *gin.Context) {
	var novoUsuario models.Usuario
	if err := c.ShouldBindJSON(&novoUsuario); err != nil {
		SendError(c, err)
		return
	}
	err := h.DoUpdateUser(&novoUsuario)
	if err != nil {
		SendError(c, err)
		return
	}
	c.HTML(http.StatusOK, "usuario-linha", gin.H{
		"message": "Usuário atualizado com sucesso",
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
	c.Status(http.StatusOK)
}

func (h *HubConexoes) HandlerAuth(c *gin.Context) {
	username := c.Param("username")
	var body struct {
		Password string `form:"password" json:"password"`
	}
	if err := c.ShouldBind(&body); err != nil {
		SendError(c, err)
		return
	}
	result, err := h.DoAuth(username, body.Password)
	if err != nil {
		SendError(c, err)
		return
	}
	if !result.Status {
		c.HTML(http.StatusBadRequest, "usuario-auth", gin.H{
			"message": result.Mensagem,
		})
		return
	}
	c.HTML(http.StatusOK, "usuario-auth", gin.H{
		"message": "Usuário autenticado com sucesso",
	})
}
