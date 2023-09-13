package router

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"supermancell/src/api"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gopkg.in/tylerb/graceful.v1"
)

/**
*  HttpRouter for gin
**/
func HttpRouter() {
	router := gin.Default()
	router.GET("/health", api.Health)
	router.GET("/scribe", api.CandleScribe)

	srv := &graceful.Server{
		Timeout: 5 * time.Second,

		Server: &http.Server{
			Addr:    ":8080",
			Handler: router,
		},
	}

	srv.ListenAndServe()
}

/**
*  WebsocketClientRouter for okex
**/
func WebsocketClientRouter(path string, ch chan string, interrupt chan os.Signal) {
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

			switch {
			case strings.Contains(string(message), "event"):
				log.Printf("recv: %s", message) //TODO 事件处理
			case strings.Contains(string(message), "data") && strings.Contains(string(message), "arg"):
				log.Printf("recv: %s", message) //TODO 推送处理
			default:
				log.Printf("recv: %s", message)
			}
		}
	}()

	for {
		select {
		case <-done:
			return
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
