package router

import (
	"github.com/catherineelfrida/mygram-api/internal/handler"
	"github.com/gin-gonic/gin"
)

type UserRouter interface {
	Mount()
}

type userRouterImpl struct {
	v       *gin.RouterGroup
	handler handler.UserHandler
}

func NewUserRouter(v *gin.RouterGroup, handler handler.UserHandler) UserRouter {
	return &userRouterImpl{v: v, handler: handler}
}

func (u *userRouterImpl) Mount() {
	u.v.GET("", u.handler.GetUsers)
	u.v.GET("/:id", u.handler.GetUsersById)
}
