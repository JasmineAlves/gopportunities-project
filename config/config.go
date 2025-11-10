package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

// Funções que retornam somente um erro
func Init() error {
	var err error

	// Inicia SQLite
	db, err = InitializeSQLite()

	if err != nil {
		return fmt.Errorf("error initializing sqlite: %v", err)
	}

	return nil
}

// Função para manipular uso do SQLite em qualquer parte do código
func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	// Inicializa Logger
	logger = NewLogger(p)
	return logger
}
