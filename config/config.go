package config

import (
	"errors"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Registrador
)

func Init() error {
	return errors.New("fake error")
}

func PegaRegistrador(p string) *Registrador {
	logger = NovoRegistrador(p)
	return logger
}
