package enumDao

import (
	"fmt"
	"orderbento/src/dao"
	"orderbento/src/models"

	"github.com/jinzhu/gorm"
)

type Enumeration models.Enumeration

type EnumType models.EnumType

func db() *gorm.DB {
	return dao.GetDB()
}

func QueryByEnumTypeCode(enumTypeCode string) (et EnumType) {

	db().Preload("Enums").Table("enum_type").Find(&et)
	fmt.Println(et)

	return et
}
