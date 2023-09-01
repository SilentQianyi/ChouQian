package config

import (
	"context"
	"flag"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	HttpConfig rest.RestConf
	LogConfig  logx.LogConf

	BaseConfig     *BaseConfig
	GuanYinQianWen GuanYinQianWenList
}

func Init() (*Config, error) {
	flag.Parse()

	baseConfig := initBaseConfig()
	guanYinQianWenList, err := initGuanYinQianWen()
	if err != nil {
		logc.Errorf(context.Background(), "config init initGuanYinQianWen error! err[ %s ]", err.Error())
		return nil, err
	}

	httpConfig := rest.RestConf{
		Host: baseConfig.HttpHost,
		Port: baseConfig.HttpPort,
	}
	logConfig := logx.LogConf{
		ServiceName: baseConfig.Name,
		Mode:        baseConfig.LogMode,
		KeepDays:    baseConfig.LogKeepDays,
		Level:       baseConfig.LogLevel,
		Path:        baseConfig.LogPath,
	}

	return &Config{
		HttpConfig:     httpConfig,
		LogConfig:      logConfig,
		BaseConfig:     baseConfig,
		GuanYinQianWen: guanYinQianWenList,
	}, nil
}
