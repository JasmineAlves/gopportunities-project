package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Inicializa Router
	router := gin.Default()

	// Inicializar Routes
	initializeRoutes(router)
	// Rodar o servidor
	router.Run(":8080")

}
