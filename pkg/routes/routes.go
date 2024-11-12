package routes

import (
	"my-go-api/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/receipts/{id}/points", handlers.GetHandler)
	router.POST("/your-post-endpoint", handlers.PostHandler)
	return router
}
