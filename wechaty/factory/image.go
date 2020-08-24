package factory

import (
	_interface "go-wechaty/wechaty/interface"
	"go-wechaty/wechaty/user"
)

type ImageFactory struct {
	_interface.IAccessory
}

func (i *ImageFactory) Create(id string) _interface.IImage {
	return user.NewImages(id, i.IAccessory)
}
