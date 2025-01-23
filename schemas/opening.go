package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Abertura struct {
	gorm.Model
	Funcao      string
	Compania    string
	Localizacao string
	Remoto      bool
	Link        string
	Salario     int64
}

type RespostaDeAbertura struct {
	ID          uint      `json: "id"`
	Criado      time.Time `json: "criado"`
	Atualizado  time.Time `json: "atualizado"`
	Deletado    time.Time `json: "deletado, omitempty"`
	Funcao      string    `json: "funcao"`
	Compania    string    `json: "compania"`
	Localizacao string    `json: "localizacao"`
	Remoto      bool      `json: "remoto"`
	Link        string    `json: "link"`
	Salario     int64     `json: "salario"`
}
