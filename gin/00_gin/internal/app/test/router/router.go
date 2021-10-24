package router

import (
	"gin_test/internal/app/test/controller"
	"gin_test/internal/container"
	"github.com/gin-gonic/gin"
)

type Router struct {
	controller *controller.Controller
}

func New(ctn *container.Container) *Router {
	return &Router{controller.New(ctn)}
}

func (r *Router) Register(app *gin.Engine) {
	v1 := app.Group("v1")
	{
		v1.GET("hello", r.controller.V1.Hello)
		v1.GET("test", r.controller.V1.Test)
		v1.GET("jwt", r.controller.V1.JwtEncode)
	}

}
