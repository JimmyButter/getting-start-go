package test

import (
	"os"
	"os/signal"
	"supermancell/src/ws"
	"testing"
	"time"
)

func TestWs(t *testing.T) {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	ws.PubConnec(interrupt)

	time.Sleep(300 * time.Second)
}
