package config

import (
	"chouQian-GoZero/util"
	"fmt"
	"github.com/SilentQianyi/parseJson"
	"log"
	"path/filepath"
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

func initGuanYinQianWen() (GuanYinQianWenList, error) {
	fpath, err := filepath.Abs("config/qianWen")
	if err != nil {
		log.Fatalf("parse conf: %s failed, %s", fpath, err.Error())
	}

	filename := fmt.Sprintf("%v/%v", fpath, "guanYinQianWen.json")

	list := make([]*GuanYinQianWen, 0)
	err = parseJson.ParseTable(filename, &list)
	if err != nil {
		log.Fatalf("initGuanYinQianWen parseJson.ParseTable error! err[ %s ]", err.Error())
		return nil, err
	}

	return list, nil
}

func (list GuanYinQianWenList) Random(utils *util.Utils) *GuanYinQianWen {
	length := len(list)
	index := utils.Rand.Intn(length)
	return list[index]
}
