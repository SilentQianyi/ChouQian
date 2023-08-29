package util

import (
	"github.com/zeromicro/go-zero/core/collection"
	"time"
)

func initMemoryCache() (memoryCache *collection.Cache, err error) {
	memoryCache, err = collection.NewCache(time.Hour*24, collection.WithName("ChouQian"))
	if err != nil {
		return nil, err
	}

	return memoryCache, nil
}
