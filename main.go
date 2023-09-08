package main

import (
	"os"
	"os/signal"
	"supermancell/src/router"
	"supermancell/src/service"
	"time"
)

func main() {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go router.HttpRouter()
	go router.WebsocketClientRouter("/ws/v5/business", service.Instance().ChBusiness, interrupt)

	beat := time.NewTicker(20 * time.Second)
	defer beat.Stop()
	for {
		select {
		case <-beat.C:
			service.ChBusinessSend("ping")
		case <-interrupt:
			return
		}

	}

}
