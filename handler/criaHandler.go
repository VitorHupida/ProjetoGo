package handler

import (
	"net/http"

	"github.com/VitorHupida/ProjetoGo/schemas"
	"github.com/gin-gonic/gin"
)

func CriaHandler(ctx *gin.Context) {
	requisicao := CriaAberturaDeChamado{}

	ctx.BindJSON(&requisicao)

	if err := requisicao.Validate(); err != nil {
		logger.Errorf("erro de validação: %v", err.Error())
		enviaErro(ctx, http.StatusBadRequest, err.Error())
		return
	}

	abertura := schemas.Abertura{
		Funcao:      requisicao.Funcao,
		Compania:    requisicao.Compania,
		Localizacao: requisicao.Localizacao,
		Remoto:      *requisicao.Remoto,
		Link:        requisicao.Link,
		Salario:     requisicao.Salario,
	}

	if err := db.Create(&abertura).Error; err != nil {
		logger.Errorf("erro ao criar abertura: %v", err.Error())
		enviaErro(ctx, http.StatusInternalServerError, "erro ao criar abertura do bando de dados")
		return
	}

	enviaSucesso(ctx, "criaAbertura", abertura)
}
