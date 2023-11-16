package main

import (
	"chouQian-GoZero/config"
	"chouQian-GoZero/handler"
	"chouQian-GoZero/logUtil"
	"chouQian-GoZero/service"
	"chouQian-GoZero/util"
	"context"
	"github.com/SilentQianyi/zeroLog"
	"github.com/zeromicro/go-zero/rest"
	"log"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatalf("config.Init error! err[ %s ]", err.Error())
		return
	}

	err = logUtil.InitLog(cfg)
	if err != nil {
		log.Fatalf("logUtil.InitLog error! err[ %s ]", err.Error())
		return
	}

	logCtx := context.Background()
	logCtx = context.WithValue(logCtx, "path", "main")

	u, err := util.InitUtils()
	if err != nil {
		log.Fatalf("util.InitUtils error! err[ %s ]", err.Error())
		return
	}

	s := service.InitService(u, cfg)
	h := handler.InitHandler(u, cfg, s)

	hServer := rest.MustNewServer(cfg.HttpConfig)
	defer hServer.Stop()

	handler.RegisterHandlers(hServer, h)

	zeroLog.Info().Ctx(logCtx).Msg("ChouQian Start!!!")
	hServer.Start()
}
