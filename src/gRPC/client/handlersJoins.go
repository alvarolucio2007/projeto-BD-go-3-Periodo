package grpcclient

import (
	"net/http"

	"github.com/alvarolucio2007/projeto-DB-go-3-Periodo/src/models"
	"github.com/gin-gonic/gin"
)

func (h *HubConexoes) HandlerLeftJoin(c *gin.Context) {
	res, err := h.DoLeftJoin()
	if err != nil {
		SendError(c, err)
		return
	}
	if res == nil {
		res = []*models.LeftJoinType{}
	}
	c.JSON(http.StatusOK, gin.H{
		"message":   "Left Join realizado com sucesso",
		"resultado": res,
	})
}

func (h *HubConexoes) HandlerInnerJoin(c *gin.Context) {
	res, err := h.DoInnerJoin()
	if err != nil {
		SendError(c, err)
		return
	}
	if res == nil {
		res = []*models.InnerJoinType{}
	}
	c.JSON(http.StatusOK, gin.H{
		"message":   "Inner Join realizado com sucesso",
		"resultado": res,
	})
}
