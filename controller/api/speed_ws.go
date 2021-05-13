package api

import (
	"fmt"
	"github.com/demonoid81/proxyM"
	"github.com/demonoid81/proxyM/log"
	"net/http"
	"time"
)

func WsSpeedHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Logger.Errorf("[Proxy-Controller] [Speed] Failed to set websocket upgrade: %v", err)
		return
	}
	ticker := time.NewTicker(time.Second)
	for {
		up, down := proxyM.CurrentSpeed()
		conn.WriteJSON(struct {
			UpSpeed   string `json:"up_speed"`
			DownSpeed string `json:"down_speed"`
		}{
			UpSpeed:   fmt.Sprintf("%s/s", capacityConversion(up)),
			DownSpeed: fmt.Sprintf("%s/s", capacityConversion(down)),
		})
		<-ticker.C
	}
}

func capacityConversion(v int) string {
	unit := "B"
	t := v
	if n := t / 1024; n >= 1 {
		unit = "KB"
		t = n
		if n := t / 1024; n >= 1 {
			unit = "MB"
			t = n
			if n := t / 1024; n >= 1 {
				unit = "GB"
				t = n
			}
		}
	}
	return fmt.Sprintf("%d%s", t, unit)
}
