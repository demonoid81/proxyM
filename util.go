package proxyM

import (
	"github.com/demonoid81/proxyM/log"
)

func Recover(fs ...func()) {
	if err := recover(); err != nil {
		log.Logger.Errorf("[PANIC] %v", err)
		for _, f := range fs {
			f()
		}
	}
}
