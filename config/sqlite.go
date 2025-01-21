package config

import (
	"os"

	"github.com/VitorHupida/ProjetoGo/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func IniciadorSQLite() (*gorm.DB, error) {
	logger := PegaRegistrador("sqlite")
	dbPath := "./db/main.db"

	//Verifica se o arquivo db já existe
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("Arquivo do banco de dados não encontrado, criando...")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	//Cria o DB e conecta
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Erro ao abrir sqlite: %v", err)
		return nil, err
	}

	//Migra o schemas
	err = db.AutoMigrate(&schemas.Abertura{})
	if err != nil {
		logger.Errorf("Erro ao migrar sqlite: %v", err)
		return nil, err
	}

	//Retorna o DB
	return db, nil
}
