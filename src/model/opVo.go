package model

type OpVo struct {
	Op   string   `json:"op"`
	Args []OpArgs `json:"args"`
}

// type instType string

const (
	InstType_Spot    string = "SPOT"    //币币
	InstType_Margin  string = "MARGIN"  //币币杠杆
	InstType_Swap    string = "SWAP"    //永续合约
	InstType_Futures string = "FUTURES" //交割合约
	InstType_Option  string = "OPTION"  //期权
	InstType_Any     string = "ANY"     //全部

	Op_Subscribe   string = "subscribe"
	Op_Unsubscribe string = "unsubscribe"
)

type OpArgs struct {
	Channel    string `json:"channel"`
	InstType   string `json:"instType,omitempty"`
	InstFamily string `json:"instFamily,omitempty"`
	InstId     string `json:"instId,omitempty"`
}
