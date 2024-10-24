package router

import (
	"commercium/interface/http/routes"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	routes.UserRoutes(v1)
	routes.ProductRoutes(v1)

	return router
}
