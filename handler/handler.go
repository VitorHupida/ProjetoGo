package handler

import (
	"github.com/VitorHupida/ProjetoGo/config"
	"gorm.io/gorm"
)

var (
	logger *config.Registrador
	db     *gorm.DB
)

func IniciaHandler() {
	logger = config.PegaRegistrador("handler")
	db = config.PegaSQLite()
}
