package config

import (
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
)

type BaseConfig struct {
	Name            string `json:"Name"`
	HttpHost        string `json:"HttpHost"`
	HttpPort        int    `json:"HttpPort"`
	LogPath         string `json:"LogPath"`
	LogMaxSize      int    `json:"LogMaxSize"`
	LogMaxAge       int    `json:"LogMaxAge"`
	LogCompress     bool   `json:"LogCompress"`
	LogMaxBackups   int    `json:"LogMaxBackups"`
	LogRotationTime int    `json:"LogRotationTime"`
}

func initBaseConfig() *BaseConfig {
	fileName := flag.String("baseConfig", "config/yaml/config.yaml", "this is yaml config file")

	baseConfig := &BaseConfig{}
	conf.MustLoad(*fileName, baseConfig)

	return baseConfig
}
