package handler

import (
	"chouQian-GoZero/config"
	"chouQian-GoZero/service"
	"chouQian-GoZero/util"
)

type Handler struct {
	QianWen *QianWenHandler
}

func InitHandler(utils *util.Utils, cfg *config.Config, svc *service.Service) *Handler {
	qianWenHandler := initQianWenHandler(utils, cfg, svc)

	return &Handler{
		QianWen: qianWenHandler,
	}
}
