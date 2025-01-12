package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeletaHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "DELETE Opening",
	})
}
