package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

var HomeDir string
var ProxyHomeDir string

func init() {
	var err error
	HomeDir, err = HomePath()
	if err != nil {
		ioutil.WriteFile("error.log", []byte(err.Error()), 0664)
		panic(err)
	} else {
		HomeDir += string(os.PathSeparator)
	}
	ProxyHomeDir = filepath.Join(HomeDir, "Documents", "proxyM")
}
