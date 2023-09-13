package model

type MessageResp struct {
	MsgArg MsgArg `json:"arg"`  //推送数据 - 订阅成功的频道
	Data   string `json:"data"` //推送数据 - 订阅的数据
}

type MsgArg struct {
	Channel string `json:"channel"`          //频道名
	InstId  string `json:"instId,omitempty"` //产品ID
}
