package container

import "gorm.io/gorm"

type Container struct {
	DB *gorm.DB
}

func New() *Container {
	ctn := &Container{}

	return ctn
}
