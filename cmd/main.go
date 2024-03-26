package main

import (
	"github.com/catherineelfrida/mygram-api/internal/handler"
	"github.com/catherineelfrida/mygram-api/internal/infrastructure"
	"github.com/catherineelfrida/mygram-api/internal/repository"
	"github.com/catherineelfrida/mygram-api/internal/router"
	"github.com/catherineelfrida/mygram-api/internal/service"
	"github.com/gin-gonic/gin"

	_ "github.com/catherineelfrida/mygram-api/cmd/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title			GO DTS USER API DUCUMENTATION
// @version		2.0
// @description	golong kominfo 006 api documentation
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.url	http://www.swagger.io/support
// @contact.email	support@swagger.io
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			localhost:3000
// @BasePath		/
// @schemes		http
func main() {
	g := gin.Default()

	usersGroup := g.Group("/users")

	// dependency injection
	// dig by uber
	// wire

	// https://s8sg.medium.com/solid-principle-in-go-e1a624290346
	gorm := infrastructure.NewGormPostgres()
	userRepo := repository.NewUserQuery(gorm)
	// userRepoMongo := repository.NewUserQueryMongo()
	userSvc := service.NewUserService(userRepo)
	userHdl := handler.NewUserHandler(userSvc)
	userRouter := router.NewUserRouter(usersGroup, userHdl)

	// mount
	userRouter.Mount()
	// swagger
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	g.Run(":3000")
}
