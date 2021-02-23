package enumService

import (
	"orderbento/src/dao/enumDao"
)

/* 查詢用戶名稱 */
func QueryByEnumTypeCode(name string) (et enumDao.EnumType) {
	return enumDao.QueryByEnumTypeCode(name)
}
