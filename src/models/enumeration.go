package models

import (
	"github.com/jinzhu/gorm"
)

type Enumeration struct {
	gorm.Model
	EnumTypeID uint
	Code       string
	Name       string
	CreateUser string
	UpdateUser string
	Version    uint
}
