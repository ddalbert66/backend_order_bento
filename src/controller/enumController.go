package controller

import (
	"fmt"
	"orderbento/src/service/enumService"

	"github.com/gin-gonic/gin"
)

type enumReq struct {
	EnumTypeCode string
}

func QueryStoreByEnumTypeCode(ctx *gin.Context) {
	var data enumReq
	err := ctx.Bind(&data)
	if err != nil {
		fmt.Println(err)
		return
	}

	enumType := enumService.QueryByEnumTypeCode("regionEnumType")

	for _, enum := range enumType.Enums {
		fmt.Println(enum.ID) // 1,2,3
	}

}
