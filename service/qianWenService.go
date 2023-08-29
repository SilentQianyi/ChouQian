package service

import (
	"chouQian-GoZero/config"
	"chouQian-GoZero/util"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type QianWenService struct {
	utils  *util.Utils
	config *config.Config
}

func newQianWenService(utils *util.Utils, cfg *config.Config) *QianWenService {
	return &QianWenService{
		utils:  utils,
		config: cfg,
	}
}

func (s *QianWenService) GuanYinQianWenKey(accountId int64) string {
	day := s.utils.GetCurDay()
	return fmt.Sprintf("GuanYin_%d_%d", day, accountId)
}

func (s *QianWenService) Chou(ctx context.Context, accountId int64) (qianWen *config.GuanYinQianWen) {
	ctx = logx.ContextWithFields(ctx, logx.Field("path", "service.QianWen.Chou"))

	key := s.GuanYinQianWenKey(accountId)
	cache, ok := s.utils.MemoryCache.Get(key)
	if ok {
		return cache.(*config.GuanYinQianWen)
	}

	qianWen = s.config.GuanYinQianWen.Random(s.utils)

	s.utils.MemoryCache.Set(key, qianWen)

	return qianWen
}
