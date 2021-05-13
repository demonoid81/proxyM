package conf

import (
	"github.com/gin-gonic/gin"
	"github.com/demonoid81/proxyM"
)

func GetMitMRules(ctx *gin.Context) {
	var response Response
	response.Data = proxyM.GetMitMRules()
	ctx.JSON(200, response)
}

func AppendMitMRules(ctx *gin.Context) {
	d := ctx.Query("domain")
	if len(d) > 0 {
		proxyM.AppendMitMRules(d)
	}
	var response Response
	response.Data = proxyM.GetMitMRules()
	ctx.JSON(200, response)
}

func DelMitMRules(ctx *gin.Context) {
	d := ctx.Query("domain")
	if len(d) > 0 {
		proxyM.RemoveMitMRules(d)
	}
	var response Response
	response.Data = proxyM.GetMitMRules()
	ctx.JSON(200, response)
}
