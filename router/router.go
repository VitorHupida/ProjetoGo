package router

import "github.com/gin-gonic/gin"

func Inicializador() {
	//inicializa o Router utilizando as configurações default do GIN
	router := gin.Default()
	//define uma rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//roda a API
	router.Run(":8080")
}
