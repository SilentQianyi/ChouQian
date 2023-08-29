package model

type QianWenReq struct {
	AccountId int64 `json:"account_id"` // 账号id
}

type QianWen struct {
	ShunXu  string `json:"shunXu"`
	QianWen string `json:"qianWen"`
	JiXiong string `json:"jiXiong"`
	GongWei string `json:"gongWei"`
	ShiYi   string `json:"shiYi"`
	JieYue  string `json:"jieYue"`
	JingSui string `json:"jingSui"`
	DianGu  string `json:"dianGu"`
	JieShi1 string `json:"jieShi1"`
	JieShi2 string `json:"jieShi2"`
}

type QianWenRsp struct {
	AccountId int64    `json:"account_id"` // 账号id
	QianWen   *QianWen `json:"qian_wen"`   // 签文
}
