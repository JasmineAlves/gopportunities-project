package config

import (
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

// Funções que retornam somente um erro
func Init() error {
	return nil
}

func GetLogger(p string) *Logger {
	// Inicializa Logger
	logger = NewLogger(p)
	return logger
}
