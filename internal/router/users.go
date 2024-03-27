package router

import (
	"github.com/gin-gonic/gin"
	"github.com/catherineelfrida/mygram-api/internal/handler"
	"github.com/catherineelfrida/mygram-api/internal/service"
	"github.com/catherineelfrida/mygram-api/pkg"
)

func SetUserRoutes(r *gin.Engine, userService *service.UserService) {
	userHandler := handler.NewUserHandler(userService)

	r.POST("/users/register", userHandler.RegisterUser)
	r.POST("/users/login", userHandler.LoginUser)

	userGroup := r.Group("/users")
	userGroup.Use(pkg.AuthMiddleware())
	{
		userGroup.PUT("/:userId", userHandler.UpdateUser)
		userGroup.DELETE("/:userId", userHandler.DeleteUser)
	}
}
