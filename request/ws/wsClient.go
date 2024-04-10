package request

import (
	"flag"
	"net/url"
	"os"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/gorilla/websocket"
)

// wss://ws.okx.com:8443/ws/v5/business
var addr = flag.String("Host", "127.0.0.1:9001", "wss service address")
var scheme = flag.String("Scheme", "ws", "websocket scheme (ws or wss)")
var defaultPath = flag.String("Path", "/echo", "Path")

func listener(path string, ch <-chan string, interrupt chan os.Signal) {

	u := url.URL{Scheme: *scheme, Host: *addr, Path: path}
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

	timePing := time.NewTicker(5 * time.Second)
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

func Connec(interrupt chan os.Signal) {
	flag.Parse()
	ch := make(chan string)
	go listener(*defaultPath, ch, interrupt)
}
