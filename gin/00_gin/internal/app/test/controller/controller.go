package controller

import (
	v1 "gin_test/internal/app/test/controller/v1"
	"gin_test/internal/container"
)

type Controller struct {
	V1 *v1.Controller
}

func New(ctn *container.Container) *Controller {
	return &Controller{v1.New(ctn)}
}
