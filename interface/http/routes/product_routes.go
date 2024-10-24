package routes

import (
	"commercium/config"
	"commercium/interface/http/handler"
	"commercium/internal/repository"
	"commercium/internal/service"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(version *gin.RouterGroup) {
	db, err := config.SetupDatabase()
	if err != nil {
		panic(err)
	}

	product_repo := repository.NewProductsRepository(db)
	product_serv := service.NewProductsService(product_repo)
	product_handler := handler.NewProductsHandler(product_serv)

	version.GET("product", product_handler.GetAllProducts)
	version.GET("product/:id", product_handler.GetProductByID)
	version.POST("product", product_handler.CreateProduct)
	version.PUT("product/:id", product_handler.UpdateProduct)
	version.DELETE("product/:id", product_handler.DeleteProduct)
}
