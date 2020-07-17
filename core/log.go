package core

import (
	"fmt"
	"gin_admin/global"
	_ "gin_admin/utils"

	oplogging "github.com/op/go-logging"
)

const (
	logDir      = "log"
	logSoftLink = "latest_log"
	module      = "gin_admin"
)

var (
	defaultFormatter = `%{time:2006/01/02 - 15:04:05.000} %{longfile} %{color:bold}▶ [%{level:.6s}] %{message}%{color:reset}`
)

func init() {
	c := global.GVA_CONFIG.Log
	if c.Prefix == "" {
		_ = fmt.Errorf("logger prefix not found")
	}
	logger := oplogging.MustGetLogger(module)
	var backends []oplogging.Backend
}
