package logUtil

import (
	"chouQian-GoZero/config"
	"github.com/SilentQianyi/zeroLog"
	"log"
)

func InitLog(cfg *config.Config) error {
	logConfig := &zeroLog.Config{
		Filename:     cfg.LogConfig.Filename,
		MaxSize:      cfg.LogConfig.MaxSize,
		MaxAge:       cfg.LogConfig.MaxAge,
		MaxBackups:   cfg.LogConfig.MaxBackups,
		Compress:     cfg.LogConfig.Compress,
		RotationTime: cfg.LogConfig.RotationTime,
	}
	err := zeroLog.Init(logConfig)
	if err != nil {
		log.Fatalf("logUtil.InitLog zeroLog.Init error! err[ %s ]", err.Error())
		return err
	}

	return nil
}
