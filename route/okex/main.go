package main

import (
	"context"
	"log"

	"github.com/amir-the-h/okex"
	"github.com/amir-the-h/okex/api"
	"github.com/amir-the-h/okex/events"
	"github.com/amir-the-h/okex/events/public"
	ws_public_requests "github.com/amir-the-h/okex/requests/ws/public"
	"hertz_demo/middleware/nacos"

)

func main() {

	nacos.GetNacosConfigInstance()
	dest := okex.NormalServer // The main API server
	ctx := context.Background()
	client, err := api.NewClient(ctx, nacos.OkexConfigYAML.ApiKey, nacos.OkexConfigYAML.SecretKey, nacos.OkexConfigYAML.Passphrase, dest)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Starting")
	errChan := make(chan *events.Error)
	subChan := make(chan *events.Subscribe)
	uSubChan := make(chan *events.Unsubscribe)
	logChan := make(chan *events.Login)
	sucChan := make(chan *events.Success)
	client.Ws.SetChannels(errChan, subChan, uSubChan, logChan, sucChan)

	obCh := make(chan *public.OrderBook)
	err = client.Ws.Public.OrderBook(ws_public_requests.OrderBook{
		InstID:  "BTC-USD-SWAP",
		Channel: "books",
	}, obCh)
	if err != nil {
		log.Fatalln(err)
	}

	csCh := make(chan *public.Tickers) // Correct channel type for Candlesticks
	err = client.Ws.Public.Tickers(ws_public_requests.Tickers{
		InstID: "BTC-USD-SWAP",
	}, csCh) // Use the correct channel here
	if err != nil {
		log.Fatalln(err)
	}

	for {
		select {

		case <-logChan:
			log.Print("[Authorized]")
		case success := <-sucChan:
			log.Printf("[SUCCESS]\t%+v", success)
		case sub := <-subChan:
			channel, _ := sub.Arg.Get("channel")
			log.Printf("[Subscribed]\t%s", channel)
		case uSub := <-uSubChan:
			channel, _ := uSub.Arg.Get("channel")
			log.Printf("[Unsubscribed]\t%s", channel)
		case err := <-client.Ws.ErrChan:
			log.Printf("[Error]\t%+v", err)
			for _, datum := range err.Data {
				log.Printf("[Error]\t\t%+v", datum)
			}
		case i := <-obCh:
			ch, _ := i.Arg.Get("channel")
			log.Printf("[Event]\t%s", ch)
			for _, p := range i.Books {
				for i := len(p.Asks) - 1; i >= 0; i-- {
					log.Printf("\t\tAsk\t%+v\t%+v\t%+v", p.Asks[i].DepthPrice, p.Asks[i].Size, p.Asks[i].OrderNumbers)
				}
				for _, bid := range p.Bids {
					log.Printf("\t\tBid\t%+v\t%+v\t%+v", bid.DepthPrice, bid.Size, bid.OrderNumbers)
				}
			}
		case cs := <-csCh:
			for _, v := range cs.Tickers {
				log.Printf("[Ticker]\t\t%+v\t%+v", v.InstID, v.Last)
			}

		case b := <-client.Ws.DoneChan:
			log.Printf("[End]:\t%v", b)
			return
		}
	}
}
