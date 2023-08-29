package util

import (
	"chouQian-GoZero/config"
	"github.com/zeromicro/go-zero/core/collection"
	"math/rand"
)

type Utils struct {
	Rand        *rand.Rand
	MemoryCache *collection.Cache
}

func InitUtils(cfg *config.Config) (*Utils, error) {

	initLog(cfg)

	memoryCache, err := initMemoryCache()
	if err != nil {
		return nil, err
	}

	u := &Utils{
		MemoryCache: memoryCache,
	}
	newRand := initRand(u)
	u.Rand = newRand

	return u, nil
}
