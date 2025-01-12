package router

import "github.com/gin-gonic/gin"

func Inicializador() {
	println()
	//inicializa o Router utilizando as configurações default do GIN
	router := gin.Default()
	//define uma rota
	inicilizandoRotas(router)
	//roda a API
	router.Run(":8080")
}
