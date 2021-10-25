package router

import (
	testRouter "gin_test/internal/app/test/router"
	"gin_test/internal/container"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Router struct {
	container  *container.Container
	testRouter *testRouter.Router
}

func New(c *container.Container) *Router {
	return &Router{
		container:  c,
		testRouter: testRouter.New(c),
	}
}

func (r *Router) Registry() *gin.Engine {
	app := gin.Default()

	r.testRouter.Register(app)
	app.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "my 404")
	})

	return app
}
