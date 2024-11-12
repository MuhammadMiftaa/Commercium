package routes

import (
	"commercium/interface/http/handler"
	"commercium/internal/repository"
	"commercium/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(version *gin.RouterGroup, db *gorm.DB) {

	User_repo := repository.NewUsersRepository(db)
	User_serv := service.NewUsersService(User_repo)
	User_handler := handler.NewUsersHandler(User_serv)

	version.POST("register", User_handler.Register)
	version.POST("login", User_handler.Login)
	version.GET("users", User_handler.GetAllUsers)
	version.GET("users/:id", User_handler.GetUserByID)
	version.POST("users", User_handler.CreateUser)
	version.PUT("users/:id", User_handler.UpdateUser)
	version.DELETE("users/:id", User_handler.DeleteUser)
}
