package service

import (
	"chouQian-GoZero/config"
	"chouQian-GoZero/util"
)

type Service struct {
	QianWen *QianWenService
}

func InitService(utils *util.Utils, cfg *config.Config) *Service {
	qianWenService := newQianWenService(utils, cfg)

	return &Service{
		QianWen: qianWenService,
	}
}
