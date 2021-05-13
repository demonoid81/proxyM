package ciphers

import (
	"fmt"
	"github.com/demonoid81/proxyM/ciphers/ssaead"
	"github.com/demonoid81/proxyM/ciphers/ssstream"
	connect "github.com/demonoid81/proxyM/conn"
)

type ConnDecorate func(password string, conn connect.IConn) (connect.IConn, error)

func CipherDecorate(password, method string, conn connect.IConn) (connect.IConn, error) {
	d := ssstream.GetStreamCiphers(method)
	if d != nil {
		return d(password, conn)
	}
	d = ssaead.GetAEADCiphers(method)
	if d != nil {
		return d(password, conn)
	}
	return nil, fmt.Errorf("[SS Cipher] not support : %s", method)
}
