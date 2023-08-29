package handler

import (
	"chouQian-GoZero/config"
	model "chouQian-GoZero/model/http"
	"chouQian-GoZero/service"
	"chouQian-GoZero/util"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type QianWenHandler struct {
	utils   *util.Utils
	config  *config.Config
	service *service.Service
}

func initQianWenHandler(utils *util.Utils, cfg *config.Config, svc *service.Service) *QianWenHandler {
	return &QianWenHandler{
		utils:   utils,
		config:  cfg,
		service: svc,
	}
}

func (h QianWenHandler) Chou() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &model.QianWenReq{}
		if err := httpx.Parse(r, req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		accountId := req.AccountId
		qianWen := h.service.QianWen.Chou(r.Context(), accountId)
		resp := &model.QianWenRsp{
			QianWen: &model.QianWen{
				ShunXu:  qianWen.ShunXu,
				QianWen: qianWen.QianWen,
				JiXiong: qianWen.JiXiong,
				GongWei: qianWen.GongWei,
				ShiYi:   qianWen.ShiYi,
				JieYue:  qianWen.JieYue,
				JingSui: qianWen.JingSui,
				DianGu:  qianWen.DianGu,
				JieShi1: qianWen.JieShi1,
				JieShi2: qianWen.JieShi2,
			},
		}
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}
