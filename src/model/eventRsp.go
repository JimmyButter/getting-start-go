package model

type EventResp struct {
	Event    string   `json:"event"` //事件
	Code     string   `json:"code"`  //错误码
	Msg      string   `json:"msg"`   //错误信息
	EventArg EventArg `json:"arg"`
}

const (
	EventType_ERROR     string = "error"
	EventType_SUBSCRIBE string = "subscribe"
)

type EventArg struct {
	Channel string `json:"channel"`          //频道名
	InstId  string `json:"instId,omitempty"` //产品ID
}
