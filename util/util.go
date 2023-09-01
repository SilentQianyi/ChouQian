package util

import (
	"github.com/zeromicro/go-zero/core/collection"
	"math/rand"
)

type Utils struct {
	Rand        *rand.Rand
	MemoryCache *collection.Cache
}

func InitUtils() (*Utils, error) {
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
