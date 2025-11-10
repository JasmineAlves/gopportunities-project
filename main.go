package main

import (
	"github.com/JasmineAlves/gopportunities-project.git/config"
	"github.com/JasmineAlves/gopportunities-project.git/router"
)

// OBS: No go, imports e exports são por convenção de nomenclatura (R = exportada, r = não exportada/var local)
var (
	logger *config.Logger
)

func main() {

	// Inicializar Configs
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Inicializar router
	router.Initialize()
}
