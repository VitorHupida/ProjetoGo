package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Registrador
)

func Init() error {
	var err error

	//Inicializa SQLite
	db, err = IniciadorSQLite()

	if err != nil {
		return fmt.Errorf("erro ao iniciar sqlite: %v", err)
	}
	return nil
}

func PegaSQLite() *gorm.DB {
	return db
}

func PegaRegistrador(p string) *Registrador {
	logger = NovoRegistrador(p)
	return logger
}
