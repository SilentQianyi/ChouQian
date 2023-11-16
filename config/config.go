package config

import (
	"flag"
	"github.com/zeromicro/go-zero/rest"
	"log"
)

type Config struct {
	HttpConfig rest.RestConf

	LogConfig      *LogConfig
	BaseConfig     *BaseConfig
	GuanYinQianWen GuanYinQianWenList
}

type LogConfig struct {
	Filename     string `json:"Filename"`
	MaxSize      int    `json:"MaxSize"`
	MaxAge       int    `json:"MaxAge"`
	MaxBackups   int    `json:"MaxBackups"`
	Compress     bool   `json:"Compress"`
	RotationTime int    `json:"RotationTime"`
}

func Init() (*Config, error) {
	flag.Parse()

	baseConfig := initBaseConfig()
	guanYinQianWenList, err := initGuanYinQianWen()
	if err != nil {
		log.Fatalf("config init initGuanYinQianWen error! err[ %s ]", err.Error())
		return nil, err
	}

	httpConfig := rest.RestConf{
		Host: baseConfig.HttpHost,
		Port: baseConfig.HttpPort,
	}
	logConfig := &LogConfig{
		Filename:     baseConfig.LogPath,
		MaxSize:      baseConfig.LogMaxSize * 1024 * 1024,
		MaxAge:       baseConfig.LogMaxAge,
		MaxBackups:   baseConfig.LogMaxBackups,
		Compress:     baseConfig.LogCompress,
		RotationTime: baseConfig.LogRotationTime,
	}

	return &Config{
		HttpConfig:     httpConfig,
		LogConfig:      logConfig,
		BaseConfig:     baseConfig,
		GuanYinQianWen: guanYinQianWenList,
	}, nil
}
