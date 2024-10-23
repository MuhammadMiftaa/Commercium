package router

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// v1 := router.Group("/v1")

	// v1.GET("books", BookHandler.GetHandler)
	// v1.GET("books/:id", BookHandler.GetHandlerById)
	// v1.POST("books", BookHandler.PostHandler)
	// v1.PUT("books/:id", BookHandler.UpdateHandler)
	// v1.DELETE("books/:id", BookHandler.DeleteHandler)

	return router
}