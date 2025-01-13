package router

import "github.com/gin-gonic/gin"

func Inicializador() {
	println()
	//inicializa o roteador utilizando as configurações default do GIN
	rota := gin.Default()
	//define uma rota
	inicilizandoRotas(rota)
	//roda a API
	rota.Run(":8080")
}
