package util

import (
	"chouQian-GoZero/config"
	"github.com/zeromicro/go-zero/core/logc"
)

func initLog(cfg *config.Config) {
	logc.MustSetup(cfg.LogConfig)
}
