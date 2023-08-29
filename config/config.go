package config

import (
	"encoding/json"
	"flag"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"os"
)

type Config struct {
	HttpConfig rest.RestConf
	LogConfig  logx.LogConf

	BaseConfig     *BaseConfig
	GuanYinQianWen GuanYinQianWenList
}

func Init() *Config {
	flag.Parse()

	baseConfig := initBaseConfig()
	guanYinQianWenList := initGuanYinQianWen()

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
	}
}

func readFile(fileName string) (b []byte, err error) {
	b, err = os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func parseTable(fileName string, i interface{}) error {
	b, e := readFile(fileName)
	if e != nil {
		return e
	}
	e = json.Unmarshal(b, i)
	if e != nil {
		return e
	}

	//logs.Info("initialize", name, " successful")
	return nil
}
