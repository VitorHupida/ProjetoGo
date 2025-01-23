package handler

import "fmt"

func erroParametroObrigatorio(nome, tipo string) error {
	return fmt.Errorf("parametro: %s (tipo: %s) é necessário", nome, tipo)
}

type CriaAberturaDeChamado struct {
	Funcao      string `json: "funcao"`
	Compania    string `json: "compania"`
	Localizacao string `json: "localizacao"`
	Remoto      *bool  `json: "remoto"`
	Link        string `json: "link"`
	Salario     int64  `json: "salario"`
}

func (c *CriaAberturaDeChamado) Validate() error {
	if c.Funcao == "" && c.Compania == "" && c.Localizacao == "" && c.Remoto == nil && c.Salario <= 0 {
		return fmt.Errorf("o corpo da solicitação está incompleto")
	}
	if c.Funcao == "" {
		return erroParametroObrigatorio("funcao", "string")
	}
	if c.Compania == "" {
		return erroParametroObrigatorio("compania", "string")
	}
	if c.Localizacao == "" {
		return erroParametroObrigatorio("localizacao", "string")
	}
	if c.Link == "" {
		return erroParametroObrigatorio("link", "string")
	}
	if c.Remoto == nil {
		return erroParametroObrigatorio("remoto", "bool")
	}
	if c.Salario <= 0 {
		return erroParametroObrigatorio("salario", "int64")
	}
	return nil
}
