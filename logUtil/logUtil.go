package logUtil

import (
	"chouQian-GoZero/config"
	"github.com/zeromicro/go-zero/core/logc"
)

func InitLog(cfg *config.Config) {
	logc.MustSetup(cfg.LogConfig)
}
