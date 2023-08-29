package config

import (
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
)

type BaseConfig struct {
	Name        string `json:"Name"`
	HttpHost    string `json:"HttpHost"`
	HttpPort    int    `json:"HttpPort"`
	LogMode     string `json:"LogMode"`
	LogKeepDays int    `json:"LogKeepDays"`
	LogLevel    string `json:"LogLevel"`
	LogPath     string `json:"LogPath"`
}

func initBaseConfig() *BaseConfig {
	fileName := flag.String("baseConfig", "config/yaml/config.yaml", "this is yaml config file")

	baseConfig := &BaseConfig{}
	conf.MustLoad(*fileName, baseConfig)

	return baseConfig
}
