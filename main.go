package main

import (
	"chouQian-GoZero/config"
	"chouQian-GoZero/handler"
	"chouQian-GoZero/logUtil"
	"chouQian-GoZero/service"
	"chouQian-GoZero/util"
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/rest"
	"log"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatalf("config.Init error! err[ %s ]", err.Error())
		return
	}

	logUtil.InitLog(cfg)
	defer logc.Close()

	u, err := util.InitUtils()
	if err != nil {
		log.Fatalf("util.InitUtils error! err[ %s ]", err.Error())
		logc.Errorf(context.Background(), "util.InitUtils error! err[ %s ]", err.Error())
		return
	}

	s := service.InitService(u, cfg)
	h := handler.InitHandler(u, cfg, s)

	hServer := rest.MustNewServer(cfg.HttpConfig)
	defer hServer.Stop()

	handler.RegisterHandlers(hServer, h)

	log.Println("ChouQian Start!!!")
	logc.Infof(context.Background(), "ChouQian Start!!!")
	hServer.Start()
}
