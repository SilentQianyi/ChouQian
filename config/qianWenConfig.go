package config

import (
	"chouQian-GoZero/util"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

type GuanYinQianWen struct {
	QianWen string `json:"签文"`
	JiXiong string `json:"吉凶"`
	GongWei string `json:"宫位"`
	ShiYi   string `json:"释义"`
	JieYue  string `json:"解签"`
	JingSui string `json:"精髓"`
	DianGu  string `json:"典故"`
	JieShi1 string `json:"解释1"`
	JieShi2 string `json:"解释2"`
}

type GuanYinQianWenList []*GuanYinQianWen

func initGuanYinQianWen() GuanYinQianWenList {
	configFile := flag.String("guanYinQianWen", "config/qianWen/test.json", "the json file")

	logx.Infof("configFile[ %s ]", *configFile)

	//list := make(map[string]GuanYinQianWen)
	//qianWen := &GuanYinQianWen{}
	list := make([]*GuanYinQianWen, 0)
	conf.MustLoad(*configFile, &list)

	return make([]*GuanYinQianWen, 0)
}

func (list GuanYinQianWenList) Random(utils *util.Utils) *GuanYinQianWen {
	length := len(list)
	index := utils.Rand.Intn(length)
	return list[index]
}
