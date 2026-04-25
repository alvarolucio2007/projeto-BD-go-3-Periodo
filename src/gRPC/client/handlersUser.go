package grpcclient

import (
	"fmt"
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
	c.HTML(http.StatusCreated, "index.html", gin.H{
		"message": "Usuário criado com sucesso",
		"id":      id,
	})
}

func (h *HubConexoes) HandlerLerUsuario(c *gin.Context) {
	username := c.PostForm("username")
	result, err := h.DoReadUser(username)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.HTML(http.StatusOK, "usuario_edicao_fragmento.html", gin.H{
		"usuarios": result,
	})
}

func (h *HubConexoes) HandlerUpdateUsuario(c *gin.Context) {
	idStr := c.PostForm("id")
	idUint, _ := strconv.ParseUint(idStr, 10, 32)

	usuario := models.Usuario{
		ID:       uint32(idUint),
		Username: c.PostForm("username"),
		Password: c.PostForm("password"), // AGORA RECEBE O VALOR
		Role:     c.PostForm("role"),
	}

	if err := h.DoUpdateUser(&usuario); err != nil {
		fmt.Println("ERRO NO GRPC:", err) // Veja o erro aqui se ainda der 500
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.HTML(http.StatusOK, "usuario_lista_fragmento_unica.html", gin.H{
		"usuario": usuario,
	})
}

func (h *HubConexoes) HandlerDeleteUsuario(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	err = h.DoDeleteUser(uint32(idUint))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}

func (h *HubConexoes) HandlerAuth(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	result, err := h.DoAuth(username, password)
	if err != nil {
		SendError(c, err)
		return
	}
	if !result.Status {
		c.String(http.StatusOK, result.Mensagem)
		return
	}
	c.Header("HX-Redirect", "/hub")
	c.Status(http.StatusOK)
}
