package grpcclient

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *HubConexoes) HandlerLeftJoin(c *gin.Context) {
	res, err := h.DoLeftJoin()
	if err != nil {
		SendError(c, err)
		return
	}
	c.HTML(http.StatusOK, "left_join", gin.H{
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
	c.HTML(http.StatusOK, "inner_join", gin.H{
		"message":   "Inner Join realizado com sucesso",
		"resultado": res,
	})
}
