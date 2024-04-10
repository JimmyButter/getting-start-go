package request_test

import (
	"hertz_demo/request/ws"
	"os"
	"os/signal"
	"testing"
	"time"
)
//wss://ws.okx.com:8443/ws/v5/business
func TestWs(t *testing.T) {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	request.PubConnec(interrupt)

	time.Sleep(300 * time.Second)
}
