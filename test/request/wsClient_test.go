package request_test

import (
	"os"
	"os/signal"
	"testing"
	"time"
	"hertz_demo/request/ws"
)

func TestWsCient(t *testing.T) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	request.Connec(interrupt)

	time.Sleep(1000 * time.Second)
}
