package config

import (
<<<<<<< HEAD
	"fmt"

=======
>>>>>>> 74534439bff3cd39b7943bee15590673f13e7b27
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

// Funções que retornam somente um erro
func Init() error {
<<<<<<< HEAD
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

=======
	return nil
}

>>>>>>> 74534439bff3cd39b7943bee15590673f13e7b27
func GetLogger(p string) *Logger {
	// Inicializa Logger
	logger = NewLogger(p)
	return logger
}
