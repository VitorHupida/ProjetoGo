package router

import (
	"github.com/VitorHupida/ProjetoGo/handler"
	"github.com/gin-gonic/gin"
)

func inicilizandoRotas(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.VeHandler)

		v1.POST("/opening", handler.CriaHandler)

		v1.DELETE("/opening", handler.DeletaHandler)

		v1.PUT("/opening", handler.AtualizaHandler)

		v1.GET("/openings", handler.ListaHandler)
	}
}
