package models

import "time"

type LoginRecord struct {
	Id         uint `gorm:"primary_key"`
	Name       string
	UserId     uint
	LoginTime  time.Time
	UserAgent  string
	Ip         string
	Header     string
	LoginState int
}
