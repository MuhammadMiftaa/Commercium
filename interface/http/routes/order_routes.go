package routes

import (
	"commercium/interface/http/handler"
	"commercium/interface/http/middlewares"
	"commercium/internal/repository"
	"commercium/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func OrderRoutes(version *gin.RouterGroup, db *gorm.DB) {

	User_repo := repository.NewUsersRepository(db)
	Product_repo := repository.NewProductsRepository(db)
	Order_repo := repository.NewOrdersRepository(db)
	Order_serv := service.NewOrdersService(Order_repo, User_repo, Product_repo)
	Order_handler := handler.NewOrdersHandler(Order_serv)

	version.Use(middlewares.AuthMiddleware())
	version.GET("orders", Order_handler.GetAllOrders)
	version.GET("orders/:id", Order_handler.GetOrderByID)
	version.POST("orders", Order_handler.CreateOrder)
	version.PUT("orders/:id", Order_handler.UpdateOrder)
	version.DELETE("orders/:id", Order_handler.DeleteOrder)
}
