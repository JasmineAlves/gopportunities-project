package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	// Agrupamento de rotas
	v1 := router.Group("/api/v1")
	{
		// Mostra opening
		v1.GET("/opening", func(ctx *gin.Context) {
			/*
			 * gin.H = tranforma JSON
			 * Handler é uma função anônima
			 */
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "GET opening",
			})
		})

		// Cria opening
		v1.POST("/opening", func(ctx *gin.Context) {
			/*
			 * gin.H = tranforma JSON
			 * Handler é uma função anônima
			 */
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "POST opening",
			})
		})

		// Deleta openning
		v1.DELETE("/opening", func(ctx *gin.Context) {
			/*
			 * gin.H = tranforma JSON
			 * Handler é uma função anônima
			 */
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "DELETE opening",
			})
		})

		v1.PUT("/opening", func(ctx *gin.Context) {
			/*
			 * gin.H = tranforma JSON
			 * Handler é uma função anônima
			 */
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "PUT opening",
			})
		})

		v1.GET("/openings", func(ctx *gin.Context) {
			/*
			 * gin.H = tranforma JSON
			 * Handler é uma função anônima
			 */
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "GET openings",
			})
		})
	}
}
