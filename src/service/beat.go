package service

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

func Beat(c *websocket.Conn) {

	timer := time.NewTimer(25000000000)
	<-timer.C

	err := c.WriteMessage(websocket.TextMessage, []byte("ping"))
	if err != nil {
		log.Println("write:", err)
		return
	}
}
