package api

import (
	"github.com/demonoid81/proxyM/dns"
	"github.com/gin-gonic/gin"
)

func DNSCacheList(ctx *gin.Context) {
	ctx.JSON(200, &Response{
		Data: dns.DNSCacheList(),
	})
}
func ClearDNSCache(ctx *gin.Context) {
	dns.ClearDNSCache()
	ctx.JSON(200, &Response{})
}

