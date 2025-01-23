package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func enviaErro(ctx *gin.Context, codigo int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(codigo, gin.H{
		"mensagem":   msg,
		"codigoErro": codigo,
	})
}

func enviaSucesso(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"mensagem": fmt.Sprintf("operação do handler: %s bem sucedida", op),
		"data":     data,
	})
}
