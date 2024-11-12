package routes

import (
	"commercium/interface/http/handler"
	"commercium/interface/http/middlewares"
	"commercium/internal/repository"
	"commercium/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProductRoutes(version *gin.RouterGroup, db *gorm.DB) {
	
	Product_repo := repository.NewProductsRepository(db)
	Product_serv := service.NewProductsService(Product_repo)
	Product_handler := handler.NewProductsHandler(Product_serv)

	version.Use(middlewares.AuthMiddleware())
	version.GET("products", Product_handler.GetAllProducts)
	version.GET("products/:id", Product_handler.GetProductByID)
	version.POST("products", Product_handler.CreateProduct)
	version.PUT("products/:id", Product_handler.UpdateProduct)
	version.DELETE("products/:id", Product_handler.DeleteProduct)
}
