package api

import (
	"github.com/demonoid81/proxyM"
	"github.com/gin-gonic/gin"
)

func GetRecords(ctx *gin.Context) {
	ctx.JSON(200, &Response{
		Data: proxyM.GetRecords(),
	})
}
func ClearRecords(ctx *gin.Context) {
	proxyM.ClearRecords()
	dump := proxyM.GetDump()
	if dump != nil {
		dump.Clear()
	}
	ctx.JSON(200, &Response{})
}
