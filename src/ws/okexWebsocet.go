package ws

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"supermancell/src/model"

	"github.com/gorilla/websocket"
)

// var addr = flag.String("addr", "ws.okx.com:8443", "wss service address")

func start(path string, ch <-chan string, interrupt chan os.Signal) {
	// flag.Parse()
	// log.SetFlags(0)

	// interrupt := make(chan os.Signal, 1)
	// signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "wss", Host: "ws.okx.com:8443", Path: path}
	log.Printf("connection to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)

		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
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
				log.Println("write:", err)
				return
			}
		case msg := <-ch:
			err := c.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
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
	op := model.OpVo{Op: "subscribe"}

	for _, instId := range instSet {
		args := model.OpArgs{Channel: "candle1H", InstId: instId}
		op.Args = append(op.Args, args)
	}

	opJson, err := json.Marshal(op)
	if err != nil {
		fmt.Println(err)
	}

	ch := make(chan string)
	go start("/ws/v5/business", ch, interrupt)

	ch <- string(opJson)
}
