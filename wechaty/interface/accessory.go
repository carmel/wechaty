package _interface

import (
  wechatyPuppet "go-wechaty/wechaty-puppet"
)

// IAccessory accessory interface
type IAccessory interface {
  GetPuppet() wechatyPuppet.IPuppetAbstract

  GetWechaty() IWechaty
}
