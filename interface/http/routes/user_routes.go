package routes

import (
	"commercium/config"
	"commercium/interface/http/handler"
	"commercium/internal/repository"
	"commercium/internal/service"

	"github.com/gin-gonic/gin"
)

func UserRoutes(version *gin.RouterGroup) {
	
	db, err := config.SetupDatabase()
	if err != nil {
		panic(err)
	}

	user_repo := repository.NewUsersRepository(db)
	user_serv := service.NewUsersService(user_repo)
	user_handler := handler.NewUsersHandler(user_serv)

	version.GET("user", user_handler.GetAllUsers)
	version.GET("user/:id", user_handler.GetUserByID)
	version.POST("user", user_handler.CreateUser)
	version.PUT("user/:id", user_handler.UpdateUser)
	version.DELETE("user/:id", user_handler.DeleteUser)
}
