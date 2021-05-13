package api

import (
	conf "github.com/demonoid81/proxyM/config"
	. "github.com/demonoid81/proxyM/constant"
	"github.com/demonoid81/proxyM/extension/config"
	"github.com/demonoid81/proxyM/upgrade"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
)

var latest string
var url string
var status string

func CheckUpdate(ctx *gin.Context) {
	var err error
	latest, url, status, err = upgrade.CheckUpgrade(conf.ProxyVersion)
	if err != nil {
		ctx.JSON(500, Response{
			Code: 1, Message: err.Error(),
		})
		return
	}
	ctx.JSON(200, Response{
		Code: 0,
		Data: map[string]string{
			"Current": conf.ProxyVersion,
			"Latest":  latest,
			"URL":     url,
			"Status":  status,
		},
	})
}

func NewUpgrade(eventChan chan *EventObj) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		if status == upgrade.VersionEqual || status == upgrade.VersionGreater {
			ctx.JSON(500, Response{
				Code: 1, Message: "You're up-to-date!",
			})
			return
		}
		path := filepath.Join(config.HomeDir, "Downloads", "proxyM.zip")
		os.Remove(path)
		err := upgrade.DownloadFile(path, url)
		if err != nil {
			ctx.JSON(500, Response{
				Code: 1, Message: err.Error(),
			})
			return
		}
		ctx.JSON(200, Response{
			Code: 0, Message: "success",
		})
		eventChan <- EventUpgrade.SetData("proxyM.zip")
	}
}
