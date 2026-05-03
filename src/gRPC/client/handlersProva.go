package grpcclient

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/gRPC/proto"
	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis"
)

type ProvaHandler struct {
	Rdb         *redis.Client
	ProvaClient *proto.ProvaServiceClient
}

func (p *ProvaHandler) HandlerCreateProva(c *gin.Context, provaConn *ProvaConexao) {
	var novaProva models.Provas
	if err := c.ShouldBindJSON(&novaProva); err != nil {
		SendError(c, err)
		return
	}

	id, err := provaConn.DoCreateProva(&novaProva)
	if err != nil {
		SendError(c, err)
	}
	fmt.Printf("Tentando salvar no banco: %+v\n", novaProva)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Prova criada com sucesso",
		"id":      id,
	})
}

func (p *ProvaHandler) HandlerReadAllProva(c *gin.Context, provaConn *ProvaConexao) {
	res, err := provaConn.DoReadAllProva()
	if err != nil {
		SendError(c, err)
		return
	}
	if res == nil {
		res = []*models.Provas{}
	}
	fmt.Printf("DEBUG: Total de provas recuperadas: %d\n", len(res))
	if len(res) > 0 {
		fmt.Printf("DEBUG: Primeira prova: %+v\n", res[0])
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Provas lidas com sucesso",
		"provas":  res, // O Gin usará as tags `json:"..."` da sua struct models.Provas
	})
}

func (p *ProvaHandler) HandlerReadAllProvas(c *gin.Context, provaConn *ProvaConexao) {
	// Sem Query, sem JSON, sem barreiras. Chamada direta.
	res, err := provaConn.DoReadAllProva() // Garanta que essa função chame o gRPC ReadAll
	if err != nil {
		SendError(c, err)
		return
	}

	// Proteção para o Python não receber 'null'
	if res == nil {
		c.JSON(http.StatusOK, gin.H{"provas": []any{}})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (p *ProvaHandler) HandlerUpdateProva(c *gin.Context, provaConn *ProvaConexao) {
	var novaProva models.Provas
	if err := c.ShouldBindJSON(&novaProva); err != nil {
		SendError(c, err)
		return
	}
	if novaProva.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID da prova é obrigratório"})
		return
	}
	if err := provaConn.DoUpdateProva(&novaProva); err != nil {
		SendError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Prova atualizada com sucesso",
	})
}

func (p *ProvaHandler) HandlerDeleteProva(c *gin.Context, provaConn *ProvaConexao) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		SendError(c, err)
		return
	}
	err = provaConn.DoDeleteProva(uint32(idUint))
	if err != nil {
		SendError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Prova deletada com sucesso",
	})
}
