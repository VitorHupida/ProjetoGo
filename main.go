package main

import (
	"github.com/VitorHupida/ProjetoGo/config"
	"github.com/VitorHupida/ProjetoGo/router"
)

var (
	logger config.Registrador
)

func main() {
	logger = *config.PegaRegistrador("main")
	//inicia as configurações
	err := config.Init()
	if err != nil {
		logger.Errorf("Erro ao iniciar configuração: %v", err)
		return
	}
	//inicia o roteador
	router.Inicializador()
}
