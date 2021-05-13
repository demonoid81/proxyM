package api

import (
	"bytes"
	"github.com/demonoid81/proxyM"
	"github.com/demonoid81/proxyM/config"
	"github.com/gin-gonic/gin"
	"io"
)

func DownloadCert(ctx *gin.Context) {
	var response Response
	caBytes := proxyM.GetCACert()
	if len(caBytes) == 0 {
		response.Code = 1
		response.Message = "please generate CA"
		ctx.JSON(500, response)
		return
	}
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("content-disposition", "attachment; filename=\"Proxy.cer\"")
	_, err := io.Copy(ctx.Writer, bytes.NewBuffer(caBytes))
	if err != nil {
		response.Code = 1
		response.Message = err.Error()
		ctx.JSON(500, response)
		return
	}
}
func GenerateCert(ctx *gin.Context) {
	var response Response
	mitm, err := proxyM.GenerateCA()
	if err != nil {
		response.Code = 1
		response.Message = err.Error()
		ctx.JSON(500, response)
		return
	}
	conf := config.CurrentConfig()
	if conf.Mitm != nil {
		mitm.Rules = conf.Mitm.Rules
	}
	conf.Mitm = mitm
	err = config.SaveConfig(config.CurrentConfigFile(), conf)
	if err != nil {
		response.Code = 1
		response.Message = err.Error()
		ctx.JSON(500, response)
		return
	}
	ctx.JSON(200, response)
}


