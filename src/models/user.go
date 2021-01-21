package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	// Id        int64
	Name      string
	Pwd       string
	SessionId string
	LoginTime *time.Time `gorm:"type:time"`
	//gorm.Model 其包括字段 ID、CreatedAt、UpdatedAt、DeletedAt
	gorm.Model
}
