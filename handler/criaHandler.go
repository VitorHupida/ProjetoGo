package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CriaHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "POST Opening",
	})
}
