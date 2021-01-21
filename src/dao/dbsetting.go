package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	var err error
	dbtype := "mysql"
	dsn := "root:abcd1234@tcp(127.0.0.1:3306)/bento?charset=utf8&loc=Local&parseTime=true"
	db, err = gorm.Open(dbtype, dsn)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
}
