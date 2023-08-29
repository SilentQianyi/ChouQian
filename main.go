package main

import (
	"chouQian-GoZero/config"
	"chouQian-GoZero/handler"
	"chouQian-GoZero/service"
	"chouQian-GoZero/util"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

func main() {
	cfg := config.Init()
	u, err := util.InitUtils(cfg)
	if err != nil {
		logx.Infof("util.InitUtils error! err[ %s ]", err.Error())
		return
	}
	defer logc.Close()

	s := service.InitService(u, cfg)
	h := handler.InitHandler(u, cfg, s)

	hServer := rest.MustNewServer(cfg.HttpConfig)
	defer hServer.Stop()

	handler.RegisterHandlers(hServer, h)

	logx.Infof("ChouQian Start!!!")
	hServer.Start()
}
