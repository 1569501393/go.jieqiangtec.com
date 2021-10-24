package v1

import (
	"gin_test/internal/container"
)

type Controller struct {
	*container.Container
}

func New(ctn *container.Container) *Controller {
	return &Controller{ctn}
}
