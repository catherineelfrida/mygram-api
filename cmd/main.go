package main

import (
	"github.com/gin-gonic/gin"
	"github.com/catherineelfrida/mygram-api/internal/infrastructure"
	"github.com/catherineelfrida/mygram-api/internal/model"
	"github.com/catherineelfrida/mygram-api/internal/repository"
	"github.com/catherineelfrida/mygram-api/internal/router"
	"github.com/catherineelfrida/mygram-api/internal/service"
)

func main() {
	r := gin.Default()

	db := infrastructure.NewGormPostgres().GetConnection()

	db.AutoMigrate(&model.User{}, &model.Photo{}, &model.Comment{}, &model.SocialMedia{})

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	router.SetUserRoutes(r, userService)

	r.Run(":8080")
}
