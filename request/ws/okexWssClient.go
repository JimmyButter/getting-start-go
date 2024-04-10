package request

import (
	"encoding/json"
	"net/url"
	"os"
	"time"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/gorilla/websocket"
)

// wss://ws.okx.com:8443/ws/v5/business
// var addr = flag.String("addr", "ws.okx.com:8443", "wss service address")

func start(path string, ch <-chan string, interrupt chan os.Signal) {

	u := url.URL{Scheme: "wss", Host: "ws.okx.com:8443", Path: path}
	hlog.Infof("connection to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		hlog.Fatalf("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)

		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				hlog.Errorf("read:", err)
				return
			}
			hlog.Infof("recv: %s", message)
		}
	}()

	timePing := time.NewTicker(25 * time.Second)
	defer timePing.Stop()

	for {
		select {
		case <-done:
			return
		case <-timePing.C:
			err := c.WriteMessage(websocket.TextMessage, []byte("ping"))
			if err != nil {
				hlog.Infof("write:", err)
				return
			}
		case msg := <-ch:
			err := c.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				hlog.Infof("write:", err)
				return
			}
		case <-interrupt:
			hlog.Infof("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				hlog.Infof("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}

func PubConnec(interrupt chan os.Signal) {
	instSet := []string{"BTC-USDT-SWAP", "ETH-USDT-SWAP"}
	op := OpVo{Op: "subscribe"}

	for _, instId := range instSet {
		args := OpArgs{Channel: "candle1H", InstId: instId}
		op.Args = append(op.Args, args)
	}

	opJson, err := json.Marshal(op)
	if err != nil {
		hlog.Error(err)
	}

	ch := make(chan string)
	go start("/ws/v5/business", ch, interrupt)

	ch <- string(opJson)
}

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

