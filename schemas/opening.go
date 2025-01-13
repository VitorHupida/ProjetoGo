package schemas

import "gorm.io/gorm"

type Abertura struct {
	gorm.Model
	Funcao      string
	Compania    string
	Localizacao string
	Remoto      bool
	Link        string
	Salario     int64
}
